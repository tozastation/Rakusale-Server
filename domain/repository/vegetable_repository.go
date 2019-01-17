package repository

import (
	"context"
	"github.com/2018-miraikeitai-org/Rakusale-Another-Server/interfaces/server/protocol"
)

// VegetableRepository is Interface of Infrastructure Repository
type VegetableRepository interface {
	FindMyVegetable(ctx context.Context, userAccessToken string) (*[]protocol.Vegetable, error)
	FindAllVegetables(ctx context.Context) (*[]protocol.Vegetable, error)
	AddMyVegetable(ctx context.Context, userAccessToken string, user protocol.Vegetable) error
	UpdateMyVegetable(ctx context.Context, userAccessToken string, user protocol.Vegetable) error
	DeleteMyVegetable(ctx context.Context, userAccessToken string, vegetableID uint) error
}
