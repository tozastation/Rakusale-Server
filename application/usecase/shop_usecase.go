package usecase

import (
	"context"
	"github.com/2018-miraikeitai-org/Rakusale-Another-Server/domain/repository"
)

type ShopUseCase interface {
	GetShopList(ctx context.Context) error
}

type shopUseCase struct {
	repository.ShopRepository
}

func NewShopUseCase(r repository.ShopRepository) ShopUseCase {
	return &shopUseCase{r}
}

func (u *shopUseCase) GetShopList(ctx context.Context) error {
	return nil
}
