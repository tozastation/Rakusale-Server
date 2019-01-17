package service

import (
	"context"
)

// UserService is Interface of Service Repository
type UserService interface {
}

type userService struct{}

// NewUserService is returning userService Struct
func NewUserService() UserService {
	return &userService{}
}

func (x *userService) DoSomething(ctx context.Context, foo int) error {
	return nil
}
