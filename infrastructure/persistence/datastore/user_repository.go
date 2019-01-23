package datastore

import (
	"context"
	"github.com/2018-miraikeitai-org/Rakusale-Another-Server/domain/model"
	"github.com/2018-miraikeitai-org/Rakusale-Another-Server/domain/repository"
	"github.com/2018-miraikeitai-org/Rakusale-Another-Server/interfaces/server/auth"
	pu "github.com/2018-miraikeitai-org/Rakusale-Another-Server/interfaces/server/rpc/user"
	"github.com/jinzhu/gorm"
)

// UserRepository is
type UserRepository struct {
	Conn *gorm.DB
}

// NewUserRepository is
func NewUserRepository(Conn *gorm.DB) repository.UserRepository {
	return &UserRepository{Conn}
}

// FindMe is ...
func (r *UserRepository) FindMe(ctx context.Context, token string) (*pu.ResponseUser, error) {
	user := model.User{}
	if err := r.Conn.Find(&user, "access_token = ?", token).Error; err != nil {
		return nil, err
	}
	return SingleUserToProtocol(user), nil
}

// FindSingleUser is ...
func (r *UserRepository) FindSingleUser(ctx context.Context, uID int64) (*pu.ResponseUser, error) {
	user := model.User{}
	if err := r.Conn.Find(&user, uID).Error; err != nil {
		return nil, err
	}
	return SingleUserToProtocol(user), nil
}

// FindAllUsers is
func (r *UserRepository) FindAllUsers(ctx context.Context) ([]*pu.ResponseUser, error) {
	user := []model.User{}
	if err := r.Conn.Limit(100).Find(&user).Error; err != nil {
		return nil, err
	}
	return UserToProtocol(user), nil
}

// AddMe is
func (r *UserRepository) AddMe(ctx context.Context, u *pu.RequestUser) error {
	name := u.GetName()
	email := u.GetEmail()
	birthday := u.GetBirthday()
	isSaler := u.GetIsSaler()
	isBuyer := u.GetIsBuyer()
	password := u.GetPassword()
	accessToken := auth.CreateJWT(name, email)
	hash, err := model.Hashed(password)
	if err != nil {
		return err
	}
	user := model.User{
		Name:        name,
		Email:       email,
		Birthday:    birthday,
		IsSaler:     isSaler,
		IsBuyer:     isBuyer,
		Password:    hash,
		AccessToken: accessToken,
	}
	if err := r.Conn.Create(&user).Error; err != nil {
		return err
	}
	return nil
}

// UpdateMe is
func (r *UserRepository) UpdateMe(ctx context.Context, token string, u *pu.RequestUser) error {
	user := model.User{}
	if err := r.Conn.Find(&user, "access_token = ?", token).Error; err != nil {
		return err
	}
	name := u.GetName()
	if name != "" {
		user.Name = name
	}
	email := u.GetEmail()
	if email != "" {
		user.Email = email
	}
	birthday := u.GetBirthday()
	if birthday != "" {
		user.Birthday = birthday
	}
	isSaler := u.GetIsSaler()
	user.IsSaler = isSaler
	isBuyer := u.GetIsBuyer()
	user.IsBuyer = isBuyer
	password := u.GetPassword()
	hash, err := model.Hashed(password)
	if err != nil {
		return err
	}
	user.Password = hash
	if err := r.Conn.Save(&user).Error; err != nil {
		return err
	}
	return nil
}

// DeleteMe is
func (r *UserRepository) DeleteMe(ctx context.Context, token string) error {
	user := model.User{}
	if err := r.Conn.Delete(&user, "access_token = ?", token).Error; err != nil {
		return err
	}
	return nil
}

// SingleUserToProtocol is ...
func SingleUserToProtocol(u model.User) *pu.ResponseUser {
	result := pu.ResponseUser{
		Id:       u.ID,
		Name:     u.Name,
		Email:    u.Email,
		Birthday: u.Birthday,
		IsSaler:  u.IsSaler,
		IsBuyer:  u.IsBuyer,
	}
	return &result
}

// UserToProtocol is ...
func UserToProtocol(u []model.User) []*pu.ResponseUser {
	result := []*pu.ResponseUser{}
	for _, a := range u {
		b := pu.ResponseUser{
			Id:       a.ID,
			Name:     a.Name,
			Email:    a.Email,
			Birthday: a.Birthday,
			IsSaler:  a.IsSaler,
			IsBuyer:  a.IsBuyer,
		}
		result = append(result, &b)
	}
	return result
}

// ProtocolToUser is ...
func ProtocolToUser(u *pu.RequestUser) model.User {
	result := model.User{
		Name:     u.Name,
		Email:    u.Email,
		Birthday: u.Birthday,
		IsSaler:  u.IsSaler,
		IsBuyer:  u.IsBuyer,
	}
	return result
}
