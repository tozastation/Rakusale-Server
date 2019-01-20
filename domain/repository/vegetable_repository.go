package repository

import (
	"context"
	"github.com/2018-miraikeitai-org/Rakusale-Another-Server/interfaces/server/protocol"
)

// VegetableRepository is Interface of Infrastructure Repository
type VegetableRepository interface {
	FindMyBoughtVegetables(ctx context.Context, token string) ([]*protocol.Vegetable, error)
	FindMySoldVegetables(ctx context.Context, token string) ([]*protocol.Vegetable, error)
	FindAllVegetables(ctx context.Context) ([]*protocol.Vegetable, error)
	AddMyVegetable(ctx context.Context, token string, v *protocol.Vegetable) error
	UpdateMyVegetable(ctx context.Context, token string, vID int64, v *protocol.Vegetable) error
	DeleteMyVegetable(ctx context.Context, token string, vID int64) error
}
