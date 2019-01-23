package usecase

import (
	"context"
	"github.com/2018-miraikeitai-org/Rakusale-Another-Server/domain/repository"
	pu "github.com/2018-miraikeitai-org/Rakusale-Another-Server/interfaces/server/rpc/user"
	"net/http"
)

// UserUseCase is RPC
type UserUseCase interface {
	GetMe(ctx context.Context, p *pu.GetMeRequest) *pu.GetMeResponse
	GetSingleUser(ctx context.Context, p *pu.GetSingleUserRequest) *pu.GetSingleUserResponse
	GetAllUsers(ctx context.Context) *pu.GetAllUsersResponse
	PostMe(ctx context.Context, p *pu.PostMeRequest) *pu.PostMeResponse
	PutMe(ctx context.Context, p *pu.PutMeRequest) *pu.PutMeResponse
	DeleteMe(ctx context.Context, p *pu.DeleteMeRequest) *pu.DeleteMeResponse
}

type userUseCase struct {
	repository.UserRepository
}

// NewUserUseCase is
func NewUserUseCase(r repository.UserRepository) UserUseCase {
	return &userUseCase{r}
}

func (u *userUseCase) GetMe(ctx context.Context, message *pu.GetMeRequest) *pu.GetMeResponse {
	res := pu.GetMeResponse{}
	user, err := u.UserRepository.FindMe(ctx, message.GetToken())
	if err != nil {
		if err != nil {
			res.Status = http.StatusBadRequest
			return &res
		}
	}
	res.User = user
	res.Status = http.StatusOK
	return &res
}

func (u *userUseCase) GetSingleUser(ctx context.Context, p *pu.GetSingleUserRequest) *pu.GetSingleUserResponse {
	res := pu.GetSingleUserResponse{}
	user, err := u.UserRepository.FindSingleUser(ctx, p.GetUserId())
	if err != nil {
		if err != nil {
			res.Status = http.StatusBadRequest
			return &res
		}
	}
	res.User = user
	res.Status = http.StatusOK
	return &res
}

func (u *userUseCase) GetAllUsers(ctx context.Context) *pu.GetAllUsersResponse {
	res := pu.GetAllUsersResponse{}
	users, err := u.UserRepository.FindAllUsers(ctx)
	if err != nil {
		if err != nil {
			res.Status = http.StatusBadRequest
			return &res
		}
	}
	res.Users = users
	res.Status = http.StatusOK
	return &res
}

func (u *userUseCase) PostMe(ctx context.Context, p *pu.PostMeRequest) *pu.PostMeResponse {
	res := pu.PostMeResponse{}
	err := u.UserRepository.AddMe(ctx, p.GetUser())
	if err != nil {
		if err != nil {
			res.Status = http.StatusBadRequest
			return &res
		}
	}
	res.Status = http.StatusCreated
	return &res
}

func (u *userUseCase) PutMe(ctx context.Context, p *pu.PutMeRequest) *pu.PutMeResponse {
	res := pu.PutMeResponse{}
	err := u.UserRepository.UpdateMe(ctx, p.GetToken(), p.GetUser())
	if err != nil {
		if err != nil {
			res.Status = http.StatusBadRequest
			return &res
		}
	}
	res.Status = http.StatusCreated
	return &res
}

func (u *userUseCase) DeleteMe(ctx context.Context, p *pu.DeleteMeRequest) *pu.DeleteMeResponse {
	res := pu.DeleteMeResponse{}
	err := u.UserRepository.DeleteMe(ctx, p.GetToken())
	if err != nil {
		if err != nil {
			res.Status = http.StatusBadRequest
			return &res
		}
	}
	res.Status = http.StatusCreated
	return &res
}
