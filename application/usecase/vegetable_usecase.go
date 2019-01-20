package usecase

import (
	"context"
	"github.com/2018-miraikeitai-org/Rakusale-Another-Server/domain/repository"
	pb "github.com/2018-miraikeitai-org/Rakusale-Another-Server/interfaces/server/protocol"
	"net/http"
)

// VegetableUseCase is ...
type VegetableUseCase interface {
	GetMyVegetables(ctx context.Context, p *pb.GetMyVegetablesRequest) *pb.GetMyVegetablesResponse
	GetAllVegetables(ctx context.Context) *pb.GetAllVegetablesResponse
	PostMyVegetable(ctx context.Context, p *pb.PostMeRequest) *pb.PostMeResponse
	PutMyVegetable(ctx context.Context, p *pb.PutMyVegetableRequest) *pb.PutMyVegetableResponse
	DeleteMyVegetable(ctx context.Context, p *pb.DeleteMyVegetableRequest) *pb.DeleteMyVegetableResponse
}

type vegetableUseCase struct {
	repository.VegetableRepository
}

// NewVegetableUseCase is ...
func NewVegetableUseCase(r repository.VegetableRepository) VegetableUseCase {
	return &vegetableUseCase{r}
}

func (u *vegetableUseCase) GetMyVegetables(ctx context.Context, p *pb.GetMyVegetablesRequest) *pb.GetMyVegetablesResponse {
	res := pb.GetMyVegetablesResponse{}
	token := p.GetToken()
	vegetables, err := u.VegetableRepository.FindMyVegetables(ctx, token)
	if err != nil {
		res.Status = http.StatusNoContent
		return &res
	}
	res.Vegetables = vegetables
	res.Status = http.StatusOK
	return &res
}

func (u *vegetableUseCase) GetAllVegetables(ctx context.Context) *pb.GetAllVegetablesResponse {
	res := pb.GetAllVegetablesResponse{}
	vegetables, err := u.VegetableRepository.FindAllVegetables(ctx)
	if err != nil {
		res.Status = http.StatusNoContent
		return &res
	}
	res.Vegetables = vegetables
	res.Status = http.StatusOK
	return &res
}

func (u *vegetableUseCase) PostMyVegetable(ctx context.Context, p *pb.PostMyVegetableRequest) *pb.PostMyVegetableResponse {
	res := pb.PostMyVegetableResponse{}
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

func (u *vegetableUseCase) PutMyVegetable(ctx context.Context, p *pb.PutMyVegetableRequest) *pb.PutMyVegetableResponse {
	res := pb.PutMyVegetableResponse{}
	token := p.GetToken()
	vegetable := p.GetVegetable()
	err := u.VegetableRepository.UpdateMyVegetable(ctx, token, vegetable)
	if err != nil {
		res.Status = http.StatusBadRequest
		return &res
	}
	res.Status = http.StatusCreated
	return &res
}

func (u *vegetableUseCase) DeleteMyVegetable(ctx context.Context, p *pb.DeleteMyVegetableRequest) *pb.DeleteMyVegetableResponse {
	res := pb.DeleteMyVegetableResponse{}
	token := p.GetToken()
	id := p.GetVegetableId()
	err := u.VegetableRepository.DeleteMyVegetable(ctx, token, id)
	if err != nil {
		res.Status = http.StatusBadRequest
		return &res
	}
	res.Status = http.StatusAccepted
	return &res
}
