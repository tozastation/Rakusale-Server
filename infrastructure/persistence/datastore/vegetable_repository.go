package datastore

import (
	"context"
	"github.com/2018-miraikeitai-org/Rakusale-Another-Server/domain/model"
	"github.com/2018-miraikeitai-org/Rakusale-Another-Server/domain/repository"
	"github.com/2018-miraikeitai-org/Rakusale-Another-Server/interfaces/server/handler"
	pv "github.com/2018-miraikeitai-org/Rakusale-Another-Server/interfaces/server/rpc/vegetable"
	"github.com/jinzhu/gorm"
	"os"
)

// VegetableRepository is
type VegetableRepository struct {
	Conn *gorm.DB
}

// NewVegetableRepository is
func NewVegetableRepository(Conn *gorm.DB) repository.VegetableRepository {
	return &VegetableRepository{Conn}
}

// FindMyBoughtVegetables is ...
func (r *VegetableRepository) FindMyBoughtVegetables(ctx context.Context, token string) ([]*pv.Vegetable, error) {
	user := model.User{}
	buyList := model.BuyList{}
	vegetables := []*model.Vegetable{}
	// ユーザの購入リスト取得
	if err := r.Conn.Find(&user, "access_token = ?", token).Related(&user.BuyList).Error; err != nil {
		return nil, err
	}
	// 購入リストから購入を１つずつ取得
	if err := r.Conn.Find(&buyList, user.BuyList.ID).Related(&buyList.List).Error; err != nil {
		return nil, err
	}
	for _, a := range buyList.List { // 購入から野菜を１つずつ購入
		buy := model.Buy{} // 購入
		if err := r.Conn.Find(&buy, a.ID).Related(&buy.Vegetables).Error; err != nil {
			return nil, err
		} // 購入に紐付く野菜を取得
		for _, b := range buy.Vegetables {
			vegetables = append(vegetables, &b)
		}
	}
	return VegetableToProtocol(vegetables), nil
}

// FindMySoldVegetables is ...
func (r *VegetableRepository) FindMySoldVegetables(ctx context.Context, token string) ([]*pv.Vegetable, error) {
	user := model.User{}
	sellList := model.SellList{}
	vegetables := []*model.Vegetable{}
	// ユーザの購入リスト取得
	if err := r.Conn.Find(&user, "access_token = ?", token).Related(&user.SellList).Error; err != nil {
		return nil, err
	}
	// 購入リストから購入を１つずつ取得
	if err := r.Conn.Find(&sellList, user.SellList.ID).Related(&sellList.List).Error; err != nil {
		return nil, err
	}
	for _, a := range sellList.List { // 購入から野菜を１つずつ購入
		sell := model.Sell{} // 購入
		if err := r.Conn.Find(&sell, a.ID).Related(&sell.Vegetables).Error; err != nil {
			return nil, err
		} // 購入に紐付く野菜を取得
		for _, b := range sell.Vegetables {
			vegetables = append(vegetables, &b)
		}
	}
	return VegetableToProtocol(vegetables), nil
}

// FindAllVegetables is
func (r *VegetableRepository) FindAllVegetables(ctx context.Context) ([]*pv.Vegetable, error) {
	a := []*model.Vegetable{}
	if err := r.Conn.Limit(100).Find(&a).Error; err != nil {
		return nil, err
	}
	return VegetableToProtocol(a), nil
}

// AddMyVegetable is
func (r *VegetableRepository) AddMyVegetable(ctx context.Context, token string, p *pv.PostMyVegetableRequest) error {
	user := model.User{}
	shop := model.Shop{}
	// トークンに紐付く直売所を取得
	if err := r.Conn.Find(&user, "access_token = ?", token).Related(&user.MyShop).Error; err != nil {
		return err
	}
	// 直売所に紐付く野菜を取得
	if err := r.Conn.Find(&shop, user.MyShop.ID).Related(&shop.Vegetables).Error; err != nil {
		return err
	}
	vegetable := ProtocolToVegetable(p.GetVegetable())
	shop.Vegetables = append(shop.Vegetables, vegetable)
	// データを更新し、更新
	if err := r.Conn.Save(&shop).Error; err != nil {
		return err
	}
	// Regist Vegetable Image
	vID := shop.Vegetables[len(shop.Vegetables)-1].ID
	vImage := p.GetImage().GetData()
	err := handler.SendImage(vImage, string(vID), "VEGETABLE_PATH")
	if err != nil {
		return err
	}
	if err := r.Conn.Find(&vegetable, vID).Error; err != nil {
		return err
	}
	ROOT := os.Getenv("GOOGLE_CLOUD_STORAGE_PUBLIC_PATH")
	DIRPATH := os.Getenv("VEGETABLE_PATH")
	vegetable.ImagePath = ROOT + DIRPATH + string(vID) + ".jpg"
	if err := r.Conn.Save(&vegetable).Error; err != nil {
		return err
	}
	return nil
}

// UpdateMyVegetable is
func (r *VegetableRepository) UpdateMyVegetable(ctx context.Context, token string, vID int64, v *pv.Vegetable) error {
	user := model.User{}
	shop := model.Shop{}
	vegetable := model.Vegetable{}
	// トークンに紐付く直売所を取得
	if err := r.Conn.Find(&user, "access_token = ?", token).Related(&user.MyShop).Error; err != nil {
		return err
	}
	// 直売所に紐付く野菜を取得
	if err := r.Conn.Find(&shop, user.MyShop.ID).Related(&shop.Vegetables).Error; err != nil {
		return err
	}
	// 直売所に紐付く野菜を取得
	if err := r.Conn.Find(&shop, user.MyShop.ID).Related(&shop.Vegetables).Error; err != nil {
		return err
	}
	// 直売所に紐付く野菜を取得
	if err := r.Conn.First(&vegetable, vID).Error; err != nil {
		return err
	}
	// データを更新し、更新
	vegetable.Name = v.Name
	vegetable.ProductionDate = v.ProductionDate
	vegetable.Fee = v.Fee
	vegetable.IsChemical = v.IsChemical
	vegetable.ImagePath = v.ImagePath
	if err := r.Conn.Save(&vegetable).Error; err != nil {
		return err
	}
	return nil
}

// DeleteMyVegetable is
func (r *VegetableRepository) DeleteMyVegetable(ctx context.Context, token string, vID int64) error {
	user := model.User{}
	shop := model.Shop{}
	vegetable := model.Vegetable{}
	// トークンに紐付く直売所を取得
	if err := r.Conn.Find(&user, "access_token = ?", token).Related(&user.MyShop).Error; err != nil {
		return err
	}
	// 直売所に紐付く野菜を取得
	if err := r.Conn.Find(&shop, user.MyShop.ID).Related(&shop.Vegetables).Error; err != nil {
		return err
	}
	// 直売所に紐付く野菜を取得
	if err := r.Conn.Find(&shop, user.MyShop.ID).Related(&shop.Vegetables).Error; err != nil {
		return err
	}
	// 直売所に紐付く野菜を取得
	if err := r.Conn.First(&vegetable, vID).Error; err != nil {
		return err
	}
	if err := r.Conn.Delete(&vegetable).Error; err != nil {
		return err
	}
	return nil
}

// VegetableToProtocol is ...
func VegetableToProtocol(v []*model.Vegetable) []*pv.Vegetable {
	result := []*pv.Vegetable{}
	for _, a := range v {
		b := pv.Vegetable{
			Name:           a.Name,
			Fee:            a.Fee,
			ImagePath:      a.ImagePath,
			ProductionDate: a.ProductionDate,
			IsChemical:     a.IsChemical,
		}
		result = append(result, &b)
	}
	return result
}

// ProtocolToVegetable is ...
func ProtocolToVegetable(v *pv.Vegetable) model.Vegetable {
	result := model.Vegetable{
		Name:           v.Name,
		Fee:            v.Fee,
		ImagePath:      v.ImagePath,
		ProductionDate: v.ProductionDate,
		IsChemical:     v.IsChemical,
	}
	return result
}
