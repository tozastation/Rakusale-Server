package model

// Vegetable Model
type Vegetable struct {
	BaseModel
	Name           string
	Fee            int64
	IsChemical     bool
	ImagePath      string
	ProductionDate string
}
