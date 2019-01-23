package usecase

import (
	"context"
	"github.com/2018-miraikeitai-org/Rakusale-Another-Server/domain/repository"
	ps "github.com/2018-miraikeitai-org/Rakusale-Another-Server/interfaces/server/rpc/shop"
	"net/http"
)

// ShopUseCase is ...
type ShopUseCase interface {
	GetMyShop(ctx context.Context, p *ps.GetMyShopRequest) *ps.GetMyShopResponse
	GetAllShops(ctx context.Context) *ps.GetAllShopsResponse
	PostMyShop(ctx context.Context, p *ps.PostMyShopRequest) *ps.PostMyShopResponse
	PutMyShop(ctx context.Context, p *ps.PutMyShopRequest) *ps.PutMyShopResponse
	DeleteMyShop(ctx context.Context, p *ps.DeleteMyShopRequest) *ps.DeleteMyShopResponse
}

type shopUseCase struct {
	repository.ShopRepository
}

// NewShopUseCase is ...
func NewShopUseCase(r repository.ShopRepository) ShopUseCase {
	return &shopUseCase{r}
}

func (u *shopUseCase) GetMyShop(ctx context.Context, p *ps.GetMyShopRequest) *ps.GetMyShopResponse {
	res := ps.GetMyShopResponse{}
	shop, err := u.ShopRepository.FindMyShop(ctx, p.GetToken())
	if err != nil {
		res.Status = http.StatusBadRequest
		return &res
	}
	res.Shop = shop
	res.Status = http.StatusOK
	return &res
}

func (u *shopUseCase) GetAllShops(ctx context.Context) *ps.GetAllShopsResponse {
	res := ps.GetAllShopsResponse{}
	shops, err := u.ShopRepository.FindAllShops(ctx)
	if err != nil {
		res.Status = http.StatusBadRequest
		return &res
	}
	res.Shops = shops
	res.Status = http.StatusOK
	return &res
}

func (u *shopUseCase) PostMyShop(ctx context.Context, p *ps.PostMyShopRequest) *ps.PostMyShopResponse {
	res := ps.PostMyShopResponse{}
	err := u.ShopRepository.AddMyShop(ctx, p.GetToken(), p.GetShop())
	if err != nil {
		res.Status = http.StatusBadRequest
		return &res
	}
	res.Status = http.StatusCreated
	return &res
}

func (u *shopUseCase) PutMyShop(ctx context.Context, p *ps.PutMyShopRequest) *ps.PutMyShopResponse {
	res := ps.PutMyShopResponse{}
	err := u.ShopRepository.UpdateMyShop(ctx, p.GetToken(), p.GetShop())
	if err != nil {
		res.Status = http.StatusBadRequest
		return &res
	}
	res.Status = http.StatusCreated
	return &res
}

func (u *shopUseCase) DeleteMyShop(ctx context.Context, p *ps.DeleteMyShopRequest) *ps.DeleteMyShopResponse {
	res := ps.DeleteMyShopResponse{}
	err := u.ShopRepository.DeleteMyShop(ctx, p.GetToken(), p.GetShopId())
	if err != nil {
		res.Status = http.StatusBadRequest
		return &res
	}
	res.Status = http.StatusCreated
	return &res
}
