package model

// Buy
type Buy struct {
	BaseModel
	BuyListID  uint
	Vegetables []Vegetable
	Total      int //合計金額
}
