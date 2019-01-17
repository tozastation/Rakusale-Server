package usecase

import (
	"context"
	"github.com/2018-miraikeitai-org/Rakusale-Another-Server/domain/repository"
	pb "github.com/2018-miraikeitai-org/Rakusale-Another-Server/interfaces/server/protocol"
)

// UserUseCase is RPC
type UserUseCase interface {
	Get(ctx context.Context, message *pb.GetUserRequest) (*pb.GetUserResponse, error)
	Post(ctx context.Context, message *pb.PostUserRequest) (*pb.PostUserResponse, error)
	Put(ctx context.Context, message *pb.PutUserRequest) (*pb.PutUserResponse, error)
	Delete(ctx context.Context, message *pb.DeleteUserRequest) (*pb.DeleteUserResponse, error)
	UploadImage(ctx context.Context, message *pb.UserImage) (*pb.UploadUserImageResponse, error)
}

type userUseCase struct {
	repository.UserRepository
}

func NewUserUseCase(r repository.UserRepository) UserUseCase {
	return &userUseCase{r}
}

func (u *userUseCase) Get(ctx context.Context, message *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	_, err := u.FindMe(ctx, message.GetToken())
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func (u *userUseCase) Post(ctx context.Context, message *pb.PostUserRequest) (*pb.PostUserResponse, error) {
	return nil, nil
}

func (u *userUseCase) Put(ctx context.Context, message *pb.PutUserRequest) (*pb.PutUserResponse, error) {
	return nil, nil
}

func (u *userUseCase) Delete(ctx context.Context, message *pb.DeleteUserRequest) (*pb.DeleteUserResponse, error) {
	return nil, nil
}

func (u *userUseCase) UploadImage(ctx context.Context, message *pb.UserImage) (*pb.UploadUserImageResponse, error) {
	return nil, nil
}
