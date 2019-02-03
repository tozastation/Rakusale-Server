package repository

import (
	"context"
	pv "github.com/2018-miraikeitai-org/Rakusale-Another-Server/interfaces/server/rpc/vegetable"
)

// VegetableRepository is Interface of Infrastructure Repository
type VegetableRepository interface {
	FindMyBoughtVegetables(ctx context.Context, token string) ([]*pv.ResponseVegetable, error)
	FindMySoldVegetables(ctx context.Context, token string) ([]*pv.ResponseVegetable, error)
	FindAllVegetables(ctx context.Context) ([]*pv.ResponseVegetable, error)
	FindSingleShopAllVegetables(ctx context.Context, sID int64) ([]*pv.ResponseVegetable, error)
	AddMyVegetable(ctx context.Context, token string, p *pv.PostMyVegetableRequest) error
	UpdateMyVegetable(ctx context.Context, token string, vID int64, v *pv.RequestVegetable) error
	DeleteMyVegetable(ctx context.Context, token string, vID int64) error
}
