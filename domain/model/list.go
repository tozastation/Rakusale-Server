package model

// BuyList is ...
type BuyList struct {
	BaseModel
	List []Buy
}

// SellList is ...
type SellList struct {
	BaseModel
	List []Sell
}
