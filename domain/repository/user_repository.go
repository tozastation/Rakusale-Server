package repository

import (
	"context"
	"github.com/2018-miraikeitai-org/Rakusale-Another-Server/interfaces/server/protocol"
)

// UserRepository is Interface of Infrastructure Repository
type UserRepository interface {
	FindMe(ctx context.Context, userAccessToken string) (*protocol.ResponseUser, error)
	FindAllUsers(ctx context.Context) (*[]protocol.ResponseUser, error)
	AddMe(ctx context.Context, userAccessToken string, user protocol.RequestUser) error
	UpdateMe(ctx context.Context, userAccessToken string, user protocol.RequestUser) error
	DeleteMe(ctx context.Context, userAccessToken string, userID uint) error
}
