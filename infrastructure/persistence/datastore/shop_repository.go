package datastore

import (
	"context"
	"github.com/2018-miraikeitai-org/Rakusale-Another-Server/domain/model"
	"github.com/2018-miraikeitai-org/Rakusale-Another-Server/domain/repository"
	ps "github.com/2018-miraikeitai-org/Rakusale-Another-Server/interfaces/server/rpc/shop"
	"github.com/jinzhu/gorm"
)

// ShopRepository is
type ShopRepository struct {
	Conn *gorm.DB
}

// NewShopRepository is
func NewShopRepository(Conn *gorm.DB) repository.ShopRepository {
	return &ShopRepository{Conn}
}

// FindMyShop is
func (r *ShopRepository) FindMyShop(ctx context.Context, token string) (*ps.ResponseShop, error) {
	user := model.User{}
	// トークンに紐付く直売所を取得
	if err := r.Conn.Find(&user, "access_token = ?", token).Related(&user.MyShop).Error; err != nil {
		return nil, err
	}
	return SingleShopToProtocol(user.MyShop), nil
}

// FindAllShops is
func (r *ShopRepository) FindAllShops(ctx context.Context) ([]*ps.ResponseShop, error) {
	a := []*model.Shop{}
	if err := r.Conn.Limit(100).Find(&a).Error; err != nil {
		return nil, err
	}
	return ShopToProtocol(a), nil
}

// AddMyShop is
func (r *ShopRepository) AddMyShop(ctx context.Context, token string, s *ps.RequestShop) error {
	user := model.User{}
	shop := ProtocolToShop(s)
	// トークンに紐付く直売所を取得
	if err := r.Conn.Find(&user, "access_token = ?", token).Error; err != nil {
		return err
	}
	user.MyShop = shop
	if err := r.Conn.Save(&user).Error; err != nil {
		return err
	}
	return nil
}

// UpdateMyShop is
func (r *ShopRepository) UpdateMyShop(ctx context.Context, token string, s *ps.RequestShop) error {
	user := model.User{}
	shop := model.Shop{}
	// トークンに紐付く直売所を取得
	if err := r.Conn.Find(&user, "access_token = ?", token).Related(&user.MyShop).Error; err != nil {
		return err
	}
	// トークンに紐付く直売所を取得
	if err := r.Conn.Find(&shop, user.MyShop.ID).Error; err != nil {
		return err
	}
	// 値あるところだけ変更
	if s.ImagePath != "" {
		shop.ImagePath = s.ImagePath
	}
	if s.Introduction != "" {
		shop.Introduction = s.Introduction
	}
	if s.Latitude != 0 {
		shop.Latitude = s.Latitude
	}
	if s.Longitude != 0 {
		shop.Longitude = s.Longitude
	}
	if s.Name != "" {
		shop.Name = s.Name
	}
	if err := r.Conn.Save(&shop).Error; err != nil {
		return err
	}
	return nil
}

// DeleteMyShop is
func (r *ShopRepository) DeleteMyShop(ctx context.Context, token string, sID int64) error {
	user := model.User{}
	shop := model.Shop{}
	if err := r.Conn.Find(&user, "access_token = ?", token).Related(&user.MyShop).Error; err != nil {
		return err
	}
	if err := r.Conn.Delete(&shop, user.MyShop.ID).Error; err != nil {
		return err
	}
	return nil
}

// SingleShopToProtocol is ...
func SingleShopToProtocol(v model.Shop) *ps.ResponseShop {
	result := ps.ResponseShop{
		Id:           v.ID,
		Name:         v.Name,
		Introduction: v.Introduction,
		Latitude:     v.Latitude,
		Longitude:    v.Latitude,
	}
	return &result
}

// ShopToProtocol is ...
func ShopToProtocol(v []*model.Shop) []*ps.ResponseShop {
	result := []*ps.ResponseShop{}
	for _, a := range v {
		b := ps.ResponseShop{
			Id:           a.ID,
			Name:         a.Name,
			Introduction: a.Introduction,
			Latitude:     a.Latitude,
			Longitude:    a.Latitude,
		}
		result = append(result, &b)
	}
	return result
}

// ProtocolToShop is ...
func ProtocolToShop(v *ps.RequestShop) model.Shop {
	result := model.Shop{
		Name:         v.Name,
		Introduction: v.Introduction,
		Latitude:     v.Latitude,
		Longitude:    v.Latitude,
	}
	return result
}
