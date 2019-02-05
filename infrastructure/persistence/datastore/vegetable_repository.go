package datastore

import (
	"context"
	"fmt"
	"github.com/2018-miraikeitai-org/Rakusale-Another-Server/domain/model"
	"github.com/2018-miraikeitai-org/Rakusale-Another-Server/domain/repository"
	"github.com/2018-miraikeitai-org/Rakusale-Another-Server/interfaces/server/handler"
	pv "github.com/2018-miraikeitai-org/Rakusale-Another-Server/interfaces/server/rpc/vegetable"
	"github.com/2018-miraikeitai-org/Rakusale-Another-Server/interfaces/server/util"
	"github.com/jinzhu/gorm"
	"os"
	"strconv"
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
func (r *VegetableRepository) FindMyBoughtVegetables(ctx context.Context, token string) ([]*pv.ResponseVegetable, error) {
	user := model.User{}
	buyList := model.BuyList{}
	vegetables := []model.Vegetable{}
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
			vegetables = append(vegetables, b)
		}
	}
	return VegetableToProtocol(vegetables), nil
}

// FindMySoldVegetables is ...
func (r *VegetableRepository) FindMySoldVegetables(ctx context.Context, token string) ([]*pv.ResponseVegetable, error) {
	user := model.User{}
	shop := model.Shop{}
	vegetable := []model.Vegetable{}
	v := []model.Vegetable{}
	// ユーザ取得
	if err := r.Conn.Find(&user, "access_token = ?", token).Error; err != nil {
		return nil, err
	}
	// 直売所取得
	if err := r.Conn.Model(&user).Related(&shop).Error; err != nil {
		return nil, err
	}
	// 直売所取得
	if err := r.Conn.Model(&shop).Related(&vegetable).Error; err != nil {
		return nil, err
	}
	fmt.Println(vegetable)
	// パース
	for _, a := range vegetable {
		v = append(v, a)
	}
	return VegetableToProtocol(v), nil
}

// FindAllVegetables is
func (r *VegetableRepository) FindAllVegetables(ctx context.Context) ([]*pv.ResponseVegetable, error) {
	a := []model.Vegetable{}
	if err := r.Conn.Limit(100).Find(&a).Error; err != nil {
		return nil, err
	}
	return VegetableToProtocol(a), nil
}

// FindSingleShopAllVegetables is
func (r *VegetableRepository) FindSingleShopAllVegetables(ctx context.Context, sID int64) ([]*pv.ResponseShopVegetable, error) {
	v := []model.Vegetable{}
	s := model.Shop{}
	// トークンに紐付く直売所を取得
	if err := r.Conn.Find(&s, sID).Error; err != nil {
		return nil, err
	}
	if err := r.Conn.Model(&s).Related(&v).Error; err != nil {
		return nil, err
	}
	buf, _ := model.CategorizeVegetable(v)
	return ShopVegetableToProtocol(buf), nil
}

// AddMyVegetable is
func (r *VegetableRepository) AddMyVegetable(ctx context.Context, token string, p *pv.PostMyVegetableRequest) error {
	user := model.User{}
	shop := model.Shop{}
	fmt.Println(p.Vegetable.Category)
	fmt.Println("[RUN] AddMyVegetable")
	// トークンに紐付く直売所を取得
	if err := r.Conn.Find(&user, "access_token = ?", token).Error; err != nil {
		return err
	}
	if err := r.Conn.Model(&user).Related(&shop).Error; err != nil {
		return err
	}
	// 直売所に紐付く野菜を取得
	fmt.Println(shop.Introduction)
	vegetable := ProtocolToVegetable(p.GetVegetable())
	shop.Vegetables = append(shop.Vegetables, vegetable)
	// データを更新し、更新
	if err := r.Conn.Save(&shop).Error; err != nil {
		return err
	}
	fmt.Println(shop.Vegetables)
	// Regist Vegetable Image
	vID := shop.Vegetables[len(shop.Vegetables)-1].ID
	fmt.Println(vID)
	vImage := p.GetImage().GetData()
	err := handler.SendImage(vImage, strconv.FormatInt(vID, 10), "VEGETABLE_PATH")
	if err != nil {
		fmt.Println(err)
	}
	if err := r.Conn.Find(&vegetable, vID).Error; err != nil {
		return err
	}
	ROOT := os.Getenv("GOOGLE_CLOUD_STORAGE_PUBLIC_PATH")
	DIRPATH := os.Getenv("VEGETABLE_PATH")
	vegetable.ImagePath = ROOT + DIRPATH + strconv.FormatInt(vID, 10) + ".jpg"
	if err := r.Conn.Save(&vegetable).Error; err != nil {
		return err
	}
	return nil
}

// UpdateMyVegetable is
func (r *VegetableRepository) UpdateMyVegetable(ctx context.Context, token string, vID int64, v *pv.RequestVegetable) error {
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
	vegetable.Name = fmt.Sprint(v.Category)
	vegetable.ProductionDate = v.ProductionDate
	vegetable.Fee = v.Fee
	vegetable.IsChemical = v.IsChemical
	fmt.Println("Category is")
	fmt.Println(v.Category)
	//vegetable.Category = v.Category
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
func VegetableToProtocol(v []model.Vegetable) []*pv.ResponseVegetable {
	result := []*pv.ResponseVegetable{}
	for _, a := range v {
		b := pv.ResponseVegetable{
			Id:             a.ID,
			Name:           a.Name,
			Fee:            a.Fee,
			ImagePath:      a.ImagePath,
			ProductionDate: a.ProductionDate,
			IsChemical:     a.IsChemical,
			Category:       pv.VegetableType(pv.VegetableType_value[a.Category]),
		}
		result = append(result, &b)
	}
	return result
}

// VegetableToProtocol is ...
func ShopVegetableToProtocol(v []model.CategorizedVegetable) []*pv.ResponseShopVegetable {
	result := []*pv.ResponseShopVegetable{}
	for _, a := range v {
		b := pv.ResponseShopVegetable{
			Name:           a.Name,
			Fee:            a.Fee,
			ImagePath:      a.ImagePath,
			ProductionDate: a.ProductionDate,
			IsChemical:     a.IsChemical,
			Category:       pv.VegetableType(pv.VegetableType_value[a.Category]),
			Amount:         a.Amount,
		}
		result = append(result, &b)
	}
	return result
}

// ProtocolToVegetable is ...
func ProtocolToVegetable(v *pv.RequestVegetable) model.Vegetable {
	category := fmt.Sprint(v.Category)
	result := model.Vegetable{
		Name:           util.ConvertEnumVegetable(category),
		Fee:            v.Fee,
		ImagePath:      "",
		ProductionDate: v.ProductionDate,
		IsChemical:     v.IsChemical,
		Category:       category,
	}
	return result
}
