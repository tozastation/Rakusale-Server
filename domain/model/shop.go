package model

// Shop is ...
type Shop struct {
	BaseModel
	ImagePath    string
	Name         string
	Latitude     float32
	Longitude    float32
	Introduction string `gorm:"type:varchar(1000);" `
	Vegetables   []Vegetable
}
