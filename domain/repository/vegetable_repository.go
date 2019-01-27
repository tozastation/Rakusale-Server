package repository

import (
	"context"
	pv "github.com/2018-miraikeitai-org/Rakusale-Another-Server/interfaces/server/rpc/vegetable"
)

// VegetableRepository is Interface of Infrastructure Repository
type VegetableRepository interface {
	FindMyBoughtVegetables(ctx context.Context, token string) ([]*pv.Vegetable, error)
	FindMySoldVegetables(ctx context.Context, token string) ([]*pv.Vegetable, error)
	FindAllVegetables(ctx context.Context) ([]*pv.Vegetable, error)
	AddMyVegetable(ctx context.Context, token string, p *pv.PostMyVegetableRequest) error
	UpdateMyVegetable(ctx context.Context, token string, vID int64, v *pv.Vegetable) error
	DeleteMyVegetable(ctx context.Context, token string, vID int64) error
}
