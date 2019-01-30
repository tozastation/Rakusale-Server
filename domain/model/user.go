package model

import (
	"golang.org/x/crypto/bcrypt"
)

// User is
type User struct {
	BaseModel
	Name        string `gorm:"not null"`
	Email       string `gorm:"unique;not null"`
	Birthday    string
	IsSaler     bool
	IsBuyer     bool
	Password    []byte   `gorm:"not null"`
	AccessToken string   `gorm:"type:varchar(1000);" `
	MyShop      Shop     //直売所リスト
	BuyList     BuyList  //購入リスト
	SellList    SellList //販売リスト
}

// Hashed is encrypting my password
func Hashed(pass string) ([]byte, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	return hash, nil
}

// CheckHash is checking my password & hashed data
func CheckHash(hash []byte, pass string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(pass))
	if err != nil {
		return err
	}
	return nil
}
