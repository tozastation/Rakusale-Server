package service

import (
	"context"
)

// VegetableService is Interface of Service Repository
type VegetableService interface {
}

type vegetableService struct{}

// NewVegetableService is returning userService Struct
func NewVegetableService() VegetableService {
	return &vegetableService{}
}

func (x *vegetableService) DoSomething(ctx context.Context, foo int) error {
	return nil
}
