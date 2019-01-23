package repository

import (
	"context"
	ps "github.com/2018-miraikeitai-org/Rakusale-Another-Server/interfaces/server/rpc/shop"
)

// ShopRepository is Interface of Infrastructure Repository
type ShopRepository interface {
	FindMyShop(ctx context.Context, token string) (*ps.ResponseShop, error)
	FindAllShops(ctx context.Context) ([]*ps.ResponseShop, error)
	AddMyShop(ctx context.Context, token string, s *ps.RequestShop) error
	UpdateMyShop(ctx context.Context, token string, s *ps.RequestShop) error
	DeleteMyShop(ctx context.Context, token string, sID int64) error
}
