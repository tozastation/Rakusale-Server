package model

// Shop is ...
type Shop struct {
	BaseModel
	UserID       uint
	ImagePath    string
	Name         string `gorm:"unique;not null"`
	Latitude     float32
	Longitude    float32
	Introduction string `gorm:"type:varchar(1000);" `
	Vegetables   []Vegetable
}
