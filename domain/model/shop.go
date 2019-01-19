package model

// Shop is ...
type Shop struct {
	BaseModel
	ImagePath    string
	Name         string
	Latitude     float64
	Longitude    float64
	Introduction string `gorm:"type:varchar(1000);" `
	Vegetables   []Vegetable
}
