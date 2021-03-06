package model

import "time"

// BaseModel is ...
type BaseModel struct {
	ID        int64      `gorm:"primary_key" json:"id"`
	CreatedAt time.Time  `json:"-"`
	UpdatedAt time.Time  `json:"-"`
	DeletedAt *time.Time `sql:"index" json:"-" `
}

// http://doc.gorm.io/associations.html
