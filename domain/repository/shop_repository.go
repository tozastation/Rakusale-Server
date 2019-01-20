package repository

import (
	"context"
	"github.com/2018-miraikeitai-org/Rakusale-Another-Server/interfaces/server/protocol"
)

// ShopRepository is Interface of Infrastructure Repository
type ShopRepository interface {
	FindMyShop(ctx context.Context, token string) (*protocol.ResponseShop, error)
	FindAllShops(ctx context.Context) ([]*protocol.ResponseShop, error)
	AddMyShop(ctx context.Context, token string, s *protocol.RequestShop) error
	UpdateMyShop(ctx context.Context, token string, s *protocol.RequestShop) error
	DeleteMyShop(ctx context.Context, token string, sID int64) error
}
