package model

// Sell is ...
type Sell struct {
	BaseModel
	SellListID uint
	Vegetables []Vegetable
}
