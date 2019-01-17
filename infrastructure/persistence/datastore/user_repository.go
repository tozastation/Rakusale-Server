package datastore

import (
	"context"
	"github.com/2018-miraikeitai-org/Rakusale-Another-Server/domain/repository"
	"github.com/2018-miraikeitai-org/Rakusale-Another-Server/interfaces/server/protocol"
	"github.com/jinzhu/gorm"
)

// UserRepository is
type UserRepository struct {
	Conn *gorm.DB
}

// NewUserRepository is
func NewUserRepository() repository.UserRepository {
	return &UserRepository{}
}

// FindMe is
func (r *UserRepository) FindMe(ctx context.Context, userAccessToken string) (*protocol.ResponseUser, error) {
	return nil, nil
}

// FindAllUsers is
func (r *UserRepository) FindAllUsers(ctx context.Context) (*[]protocol.ResponseUser, error) {
	return nil, nil
}

// AddMe is
func (r *UserRepository) AddMe(ctx context.Context, userAccessToken string, user protocol.RequestUser) error {
	return nil
}

// UpdateMe is
func (r *UserRepository) UpdateMe(ctx context.Context, userAccessToken string, user protocol.RequestUser) error {
	return nil
}

// DeleteMe is
func (r *UserRepository) DeleteMe(ctx context.Context, userAccessToken string, userID uint) error {
	return nil
}
