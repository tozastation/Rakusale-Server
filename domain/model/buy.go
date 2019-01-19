package model

// Buy
type Buy struct {
	BaseModel
	Vegetables []Vegetable
	Total      int //合計金額
}
