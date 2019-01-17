package datastore

import (
	"context"
	"github.com/2018-miraikeitai-org/Rakusale-Another-Server/domain/repository"
	"github.com/2018-miraikeitai-org/Rakusale-Another-Server/interfaces/server/protocol"
	"github.com/jinzhu/gorm"
)

// ShopRepository is
type ShopRepository struct {
	Conn *gorm.DB
}

// NewShopRepository is
func NewShopRepository() repository.ShopRepository {
	return &ShopRepository{}
}

// FindMyShop is
func (r *ShopRepository) FindMyShop(ctx context.Context, userAccessToken string) (*[]protocol.Shop, error) {
	return nil, nil
}

// FindAllShops is
func (r *ShopRepository) FindAllShops(ctx context.Context) (*[]protocol.Shop, error) {
	return nil, nil
}

// AddMyShop is
func (r *ShopRepository) AddMyShop(ctx context.Context, userAccessToken string, user protocol.Shop) error {
	return nil
}

// UpdateMyShop is
func (r *ShopRepository) UpdateMyShop(ctx context.Context, userAccessToken string, user protocol.Shop) error {
	return nil
}

// DeleteMyShop is
func (r *ShopRepository) DeleteMyShop(ctx context.Context, userAccessToken string, shopID uint) error {
	return nil
}
