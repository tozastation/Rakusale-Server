package repository

import (
	"context"
	"github.com/2018-miraikeitai-org/Rakusale-Another-Server/interfaces/server/protocol"
)

// UserRepository is Interface of Infrastructure Repository
type UserRepository interface {
	FindMe(ctx context.Context, token string) (*protocol.ResponseUser, error)
	FindSingleUser(ctx context.Context, uID int64) (*protocol.ResponseUser, error)
	FindAllUsers(ctx context.Context) ([]*protocol.ResponseUser, error)
	AddMe(ctx context.Context, token string, u *protocol.RequestUser) error
	UpdateMe(ctx context.Context, token string, u *protocol.RequestUser) error
	DeleteMe(ctx context.Context, token string, uID int64) error
}
