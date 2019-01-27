package datastore

import (
	"context"
	"github.com/2018-miraikeitai-org/Rakusale-Another-Server/domain/model"
	"github.com/2018-miraikeitai-org/Rakusale-Another-Server/domain/repository"
	"github.com/2018-miraikeitai-org/Rakusale-Another-Server/interfaces/server/auth"
	pe "github.com/2018-miraikeitai-org/Rakusale-Another-Server/interfaces/server/rpc/entry"
	"github.com/jinzhu/gorm"
)

// EntryRepository is
type EntryRepository struct {
	Conn *gorm.DB
}

// NewEntryRepository is
func NewEntryRepository(Conn *gorm.DB) repository.EntryRepository {
	return &EntryRepository{Conn}
}

// CheckCorrectUser is ...
func (r *EntryRepository) CheckCorrectUser(ctx context.Context, p *pe.SignInRequest) (*pe.EntryResponse, error) {
	user := model.User{}
	inputEmail := p.GetEmail()
	inputPassword := p.GetPassword()
	if err := r.Conn.Find(&user, "email = ?", inputEmail).Error; err != nil {
		return nil, err
	}
	// Check Password
	hashed, err := model.Hashed(inputPassword)
	if err != nil {
		return nil, err
	}
	err = model.CheckHash(hashed, inputPassword)
	if err != nil {
		return nil, err
	}
	// Make Response
	res := pe.EntryResponse{}
	res.Token = user.AccessToken
	return &res, nil
}

// CreateNewUser is
func (r *EntryRepository) CreateNewUser(ctx context.Context, p *pe.SignUpRequest) (*pe.EntryResponse, error) {
	// Set Value
	name := p.GetName()
	email := p.GetEmail()
	password := p.GetPassword()
	hash, err := model.Hashed(password)
	if err != nil {
		return nil, err
	}
	user := model.User{
		Name:        p.GetName(),
		Email:       p.GetEmail(),
		Birthday:    p.GetBirthday(),
		IsSaler:     p.GetIsSaler(),
		IsBuyer:     p.GetIsBuyer(),
		Password:    hash,
		AccessToken: auth.CreateJWT(name, email),
	}
	if err := r.Conn.Create(&user).Error; err != nil {
		return nil, err
	}
	res := pe.EntryResponse{}
	res.Token = user.AccessToken
	return &res, nil
}
