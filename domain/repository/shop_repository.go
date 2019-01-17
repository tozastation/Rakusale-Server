package repository

import (
	"context"
	"github.com/2018-miraikeitai-org/Rakusale-Another-Server/interfaces/server/protocol"
)

// ShopRepository is Interface of Infrastructure Repository
type ShopRepository interface {
	FindMyShop(ctx context.Context, userAccessToken string) (*[]protocol.Shop, error)
	FindAllShops(ctx context.Context) (*[]protocol.Shop, error)
	AddMyShop(ctx context.Context, userAccessToken string, user protocol.Shop) error
	UpdateMyShop(ctx context.Context, userAccessToken string, user protocol.Shop) error
	DeleteMyShop(ctx context.Context, userAccessToken string, shopID uint) error
}
