package model

// BuyList is ...
type BuyList struct {
	BaseModel
	UserID uint
	List   []Buy
}

// SellList is ...
type SellList struct {
	BaseModel
	UserID uint
	List   []Sell
}
