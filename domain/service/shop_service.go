package service

import (
	"context"
)

// ShopService is Interface of Service Repository
type ShopService interface {
}

type shopService struct{}

// NewShopService is returning userService Struct
func NewShopService() ShopService {
	return &shopService{}
}

func (x *shopService) DoSomething(ctx context.Context, foo int) error {
	return nil
}
