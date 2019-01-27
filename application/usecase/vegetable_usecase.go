package usecase

import (
	"context"
	"github.com/2018-miraikeitai-org/Rakusale-Another-Server/domain/repository"
	pv "github.com/2018-miraikeitai-org/Rakusale-Another-Server/interfaces/server/rpc/vegetable"
	"net/http"
)

// VegetableUseCase is ...
type VegetableUseCase interface {
	GetMyBoughtVegetables(ctx context.Context, p *pv.GetMyVegetablesRequest) (*pv.GetMyVegetablesResponse, error)
	GetMySoldVegetables(ctx context.Context, p *pv.GetMyVegetablesRequest) (*pv.GetMyVegetablesResponse, error)
	GetAllVegetables(ctx context.Context, p *pv.VegetablesEmpty) (*pv.GetAllVegetablesResponse, error)
	PostMyVegetable(ctx context.Context, p *pv.PostMyVegetableRequest) (*pv.PostMyVegetableResponse, error)
	PutMyVegetable(ctx context.Context, p *pv.PutMyVegetableRequest) (*pv.PutMyVegetableResponse, error)
	DeleteMyVegetable(ctx context.Context, p *pv.DeleteMyVegetableRequest) (*pv.DeleteMyVegetableResponse, error)
}

type vegetableUseCase struct {
	repository.VegetableRepository
}

// NewVegetableUseCase is ...
func NewVegetableUseCase(r repository.VegetableRepository) VegetableUseCase {
	return &vegetableUseCase{r}
}

func (u *vegetableUseCase) GetMyBoughtVegetables(ctx context.Context, p *pv.GetMyVegetablesRequest) (*pv.GetMyVegetablesResponse, error) {
	res := pv.GetMyVegetablesResponse{}
	token := p.GetToken()
	vegetables, err := u.VegetableRepository.FindMyBoughtVegetables(ctx, token)
	if err != nil {
		res.Status = http.StatusNoContent
		return &res, nil
	}
	res.Vegetables = vegetables
	res.Status = http.StatusOK
	return &res, nil
}

func (u *vegetableUseCase) GetMySoldVegetables(ctx context.Context, p *pv.GetMyVegetablesRequest) (*pv.GetMyVegetablesResponse, error) {
	res := pv.GetMyVegetablesResponse{}
	token := p.GetToken()
	vegetables, err := u.VegetableRepository.FindMySoldVegetables(ctx, token)
	if err != nil {
		res.Status = http.StatusNoContent
		return &res, nil
	}
	res.Vegetables = vegetables
	res.Status = http.StatusOK
	return &res, nil
}

func (u *vegetableUseCase) GetAllVegetables(ctx context.Context, p *pv.VegetablesEmpty) (*pv.GetAllVegetablesResponse, error) {
	res := pv.GetAllVegetablesResponse{}
	vegetables, err := u.VegetableRepository.FindAllVegetables(ctx)
	if err != nil {
		res.Status = http.StatusNoContent
		return &res, nil
	}
	res.Vegetables = vegetables
	res.Status = http.StatusOK
	return &res, nil
}

func (u *vegetableUseCase) PostMyVegetable(ctx context.Context, p *pv.PostMyVegetableRequest) (*pv.PostMyVegetableResponse, error) {
	res := pv.PostMyVegetableResponse{}
	token := p.GetToken()
	err := u.VegetableRepository.AddMyVegetable(ctx, token, p)
	if err != nil {
		res.Status = http.StatusBadRequest
		return &res, nil
	}
	res.Status = http.StatusCreated
	return &res, nil
}

func (u *vegetableUseCase) PutMyVegetable(ctx context.Context, p *pv.PutMyVegetableRequest) (*pv.PutMyVegetableResponse, error) {
	res := pv.PutMyVegetableResponse{}
	err := u.VegetableRepository.UpdateMyVegetable(ctx, p.GetToken(), p.GetVId(), p.GetVegetable())
	if err != nil {
		res.Status = http.StatusBadRequest
		return &res, nil
	}
	res.Status = http.StatusCreated
	return &res, nil
}

func (u *vegetableUseCase) DeleteMyVegetable(ctx context.Context, p *pv.DeleteMyVegetableRequest) (*pv.DeleteMyVegetableResponse, error) {
	res := pv.DeleteMyVegetableResponse{}
	err := u.VegetableRepository.DeleteMyVegetable(ctx, p.GetToken(), p.GetVId())
	if err != nil {
		res.Status = http.StatusBadRequest
		return &res, nil
	}
	res.Status = http.StatusAccepted
	return &res, nil
}
