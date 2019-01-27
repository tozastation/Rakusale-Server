package registry

import (
	"github.com/2018-miraikeitai-org/Rakusale-Another-Server/application/usecase"
	"github.com/2018-miraikeitai-org/Rakusale-Another-Server/domain/repository"
)

// Registry is ...
type Registry interface {
	NewUserRepository() repository.UserRepository
	NewUserUseCase() usecase.UserUseCase
	NewShopRepository() repository.ShopRepository
	NewShopUseCase() usecase.ShopUseCase
	NewVegetableRepository() repository.VegetableRepository
	NewVegetableUseCase() usecase.VegetableUseCase
}
