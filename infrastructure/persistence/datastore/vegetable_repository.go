package datastore

import (
	"context"
	"github.com/2018-miraikeitai-org/Rakusale-Another-Server/domain/repository"
	"github.com/2018-miraikeitai-org/Rakusale-Another-Server/interfaces/server/protocol"
)

// VegetableRepository is
type VegetableRepository struct{}

// NewVegetableRepository is
func NewVegetableRepository() repository.VegetableRepository {
	return &VegetableRepository{}
}

// FindMyVegetable is
func (r *VegetableRepository) FindMyVegetable(ctx context.Context, userAccessToken string) (*[]protocol.Vegetable, error) {
	return nil, nil
}

// FindAllVegetables is
func (r *VegetableRepository) FindAllVegetables(ctx context.Context) (*[]protocol.Vegetable, error) {
	return nil, nil
}

// AddMyVegetable is
func (r *VegetableRepository) AddMyVegetable(ctx context.Context, userAccessToken string, user protocol.Vegetable) error {
	return nil
}

// UpdateMyVegetable is
func (r *VegetableRepository) UpdateMyVegetable(ctx context.Context, userAccessToken string, user protocol.Vegetable) error {
	return nil
}

// DeleteMyVegetable is
func (r *ShopRepository) DeleteMyVegetable(ctx context.Context, userAccessToken string, vegetableID uint) error {
	return nil
}
