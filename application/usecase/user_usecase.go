package usecase

import (
	"context"
	"github.com/2018-miraikeitai-org/Rakusale-Another-Server/domain/repository"
	pu "github.com/2018-miraikeitai-org/Rakusale-Another-Server/interfaces/server/rpc/user"
	"net/http"
)

// UserUseCase is RPC
type UserUseCase interface {
	GetMe(ctx context.Context, p *pu.GetMeRequest) (*pu.GetMeResponse, error)
	GetSingleUser(ctx context.Context, p *pu.GetSingleUserRequest) (*pu.GetSingleUserResponse, error)
	GetAllUsers(ctx context.Context, p *pu.UsersEmpty) (*pu.GetAllUsersResponse, error)
	PostMe(ctx context.Context, p *pu.PostMeRequest) (*pu.PostMeResponse, error)
	PutMe(ctx context.Context, p *pu.PutMeRequest) (*pu.PutMeResponse, error)
	DeleteMe(ctx context.Context, p *pu.DeleteMeRequest) (*pu.DeleteMeResponse, error)
}

type userUseCase struct {
	repository.UserRepository
}

// NewUserUseCase is
func NewUserUseCase(r repository.UserRepository) UserUseCase {
	return &userUseCase{r}
}

func (u *userUseCase) GetMe(ctx context.Context, message *pu.GetMeRequest) (*pu.GetMeResponse, error) {
	res := pu.GetMeResponse{}
	user, err := u.UserRepository.FindMe(ctx, message.GetToken())
	if err != nil {
		if err != nil {
			res.Status = http.StatusBadRequest
			return &res, nil
		}
	}
	res.User = user
	res.Status = http.StatusOK
	return &res, nil
}

func (u *userUseCase) GetSingleUser(ctx context.Context, p *pu.GetSingleUserRequest) (*pu.GetSingleUserResponse, error) {
	res := pu.GetSingleUserResponse{}
	user, err := u.UserRepository.FindSingleUser(ctx, p.GetUserId())
	if err != nil {
		if err != nil {
			res.Status = http.StatusBadRequest
			return &res, nil
		}
	}
	res.User = user
	res.Status = http.StatusOK
	return &res, nil
}

func (u *userUseCase) GetAllUsers(ctx context.Context, p *pu.UsersEmpty) (*pu.GetAllUsersResponse, error) {
	res := pu.GetAllUsersResponse{}
	users, err := u.UserRepository.FindAllUsers(ctx)
	if err != nil {
		if err != nil {
			res.Status = http.StatusBadRequest
			return &res, nil
		}
	}
	res.Users = users
	res.Status = http.StatusOK
	return &res, nil
}

func (u *userUseCase) PostMe(ctx context.Context, p *pu.PostMeRequest) (*pu.PostMeResponse, error) {
	res := pu.PostMeResponse{}
	err := u.UserRepository.AddMe(ctx, p.GetUser())
	if err != nil {
		if err != nil {
			res.Status = http.StatusBadRequest
			return &res, nil
		}
	}
	res.Status = http.StatusCreated
	return &res, nil
}

func (u *userUseCase) PutMe(ctx context.Context, p *pu.PutMeRequest) (*pu.PutMeResponse, error) {
	res := pu.PutMeResponse{}
	err := u.UserRepository.UpdateMe(ctx, p.GetToken(), p.GetUser())
	if err != nil {
		if err != nil {
			res.Status = http.StatusBadRequest
			return &res, nil
		}
	}
	res.Status = http.StatusCreated
	return &res, nil
}

func (u *userUseCase) DeleteMe(ctx context.Context, p *pu.DeleteMeRequest) (*pu.DeleteMeResponse, error) {
	res := pu.DeleteMeResponse{}
	err := u.UserRepository.DeleteMe(ctx, p.GetToken())
	if err != nil {
		if err != nil {
			res.Status = http.StatusBadRequest
			return &res, nil
		}
	}
	res.Status = http.StatusCreated
	return &res, nil
}
