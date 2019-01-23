package usecase

import (
	"context"
	"github.com/2018-miraikeitai-org/Rakusale-Another-Server/domain/repository"
	pv "github.com/2018-miraikeitai-org/Rakusale-Another-Server/interfaces/server/rpc/vegetable"
	"net/http"
)

// VegetableUseCase is ...
type VegetableUseCase interface {
	GetMyBoughtVegetables(ctx context.Context, p *pv.GetMyVegetablesRequest) *pv.GetMyVegetablesResponse
	GetMySoldVegetables(ctx context.Context, p *pv.GetMyVegetablesRequest) *pv.GetMyVegetablesResponse
	GetAllVegetables(ctx context.Context) *pv.GetAllVegetablesResponse
	PostMyVegetable(ctx context.Context, p *pv.PostMyVegetableRequest) *pv.PostMyVegetableResponse
	PutMyVegetable(ctx context.Context, p *pv.PutMyVegetableRequest) *pv.PutMyVegetableResponse
	DeleteMyVegetable(ctx context.Context, p *pv.DeleteMyVegetableRequest) *pv.DeleteMyVegetableResponse
}

type vegetableUseCase struct {
	repository.VegetableRepository
}

// NewVegetableUseCase is ...
func NewVegetableUseCase(r repository.VegetableRepository) VegetableUseCase {
	return &vegetableUseCase{r}
}

func (u *vegetableUseCase) GetMyBoughtVegetables(ctx context.Context, p *pv.GetMyVegetablesRequest) *pv.GetMyVegetablesResponse {
	res := pv.GetMyVegetablesResponse{}
	token := p.GetToken()
	vegetables, err := u.VegetableRepository.FindMyBoughtVegetables(ctx, token)
	if err != nil {
		res.Status = http.StatusNoContent
		return &res
	}
	res.Vegetables = vegetables
	res.Status = http.StatusOK
	return &res
}

func (u *vegetableUseCase) GetMySoldVegetables(ctx context.Context, p *pv.GetMyVegetablesRequest) *pv.GetMyVegetablesResponse {
	res := pv.GetMyVegetablesResponse{}
	token := p.GetToken()
	vegetables, err := u.VegetableRepository.FindMySoldVegetables(ctx, token)
	if err != nil {
		res.Status = http.StatusNoContent
		return &res
	}
	res.Vegetables = vegetables
	res.Status = http.StatusOK
	return &res
}

func (u *vegetableUseCase) GetAllVegetables(ctx context.Context) *pv.GetAllVegetablesResponse {
	res := pv.GetAllVegetablesResponse{}
	vegetables, err := u.VegetableRepository.FindAllVegetables(ctx)
	if err != nil {
		res.Status = http.StatusNoContent
		return &res
	}
	res.Vegetables = vegetables
	res.Status = http.StatusOK
	return &res
}

func (u *vegetableUseCase) PostMyVegetable(ctx context.Context, p *pv.PostMyVegetableRequest) *pv.PostMyVegetableResponse {
	res := pv.PostMyVegetableResponse{}
	token := p.GetToken()
	vegetable := p.GetVegetable()
	err := u.VegetableRepository.AddMyVegetable(ctx, token, vegetable)
	if err != nil {
		res.Status = http.StatusBadRequest
		return &res
	}
	res.Status = http.StatusCreated
	return &res
}

func (u *vegetableUseCase) PutMyVegetable(ctx context.Context, p *pv.PutMyVegetableRequest) *pv.PutMyVegetableResponse {
	res := pv.PutMyVegetableResponse{}
	err := u.VegetableRepository.UpdateMyVegetable(ctx, p.GetToken(), p.GetVId(), p.GetVegetable())
	if err != nil {
		res.Status = http.StatusBadRequest
		return &res
	}
	res.Status = http.StatusCreated
	return &res
}

func (u *vegetableUseCase) DeleteMyVegetable(ctx context.Context, p *pv.DeleteMyVegetableRequest) *pv.DeleteMyVegetableResponse {
	res := pv.DeleteMyVegetableResponse{}
	err := u.VegetableRepository.DeleteMyVegetable(ctx, p.GetToken(), p.GetVId())
	if err != nil {
		res.Status = http.StatusBadRequest
		return &res
	}
	res.Status = http.StatusAccepted
	return &res
}
