package model

// Vegetable Model
type Vegetable struct {
	BaseModel
	ShopID         uint
	SellID         uint
	BuyID          uint
	Name           string
	Fee            int64
	IsChemical     bool
	ImagePath      string
	ProductionDate string
}
