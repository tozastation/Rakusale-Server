package repository

import (
	"context"
	pu "github.com/2018-miraikeitai-org/Rakusale-Another-Server/interfaces/server/rpc/user"
)

// UserRepository is Interface of Infrastructure Repository
type UserRepository interface {
	FindMe(ctx context.Context, token string) (*pu.ResponseUser, error)
	FindSingleUser(ctx context.Context, uID int64) (*pu.ResponseUser, error)
	FindAllUsers(ctx context.Context) ([]*pu.ResponseUser, error)
	AddMe(ctx context.Context, u *pu.RequestUser) error
	UpdateMe(ctx context.Context, token string, u *pu.RequestUser) error
	DeleteMe(ctx context.Context, token string) error
}
