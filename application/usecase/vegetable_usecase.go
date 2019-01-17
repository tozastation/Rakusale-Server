package usecase

import (
	"context"
	"github.com/2018-miraikeitai-org/Rakusale-Another-Server/domain/repository"
)

type VegetableUseCase interface {
	GetVegetableList(ctx context.Context) error
}

type vegetableUseCase struct {
	repository.VegetableRepository
}

func NewVegetableUseCase(r repository.VegetableRepository) VegetableUseCase {
	return &vegetableUseCase{r}
}

func (u *vegetableUseCase) GetVegetableList(ctx context.Context) error {
	return nil
}
