package handler

import (
	"fmt"
	"github.com/2018-miraikeitai-org/Rakusale-Another-Server/domain/model"
	"github.com/jinzhu/gorm"
	"os"
)

var db *gorm.DB

// OpenDBConnection is
func OpenDBConnection() *gorm.DB {
	dbType := os.Getenv("DB_TYPE")
	connectionString := os.Getenv("CONNECTION_STRING")
	db, _ := gorm.Open(dbType, connectionString)
	// [Initialize] Database
	db.AutoMigrate(&model.Vegetable{})
	fmt.Println("[Init] Vegetable Model")
	db.AutoMigrate(&model.Sell{})
	fmt.Println("[Init] Sell Model")
	db.AutoMigrate(&model.Buy{})
	fmt.Println("[Init] Buy Model")
	db.AutoMigrate(&model.SellList{})
	fmt.Println("[Init] SellList Model")
	db.AutoMigrate(&model.BuyList{})
	fmt.Println("[Init] BuyList Model")
	db.AutoMigrate(&model.Shop{})
	fmt.Println("[Init] Shop Model")
	db.AutoMigrate(&model.User{})
	fmt.Println("[Init] User Model")

	return db
}

// CloseDBConnction is
func CloseDBConnction() {
	db.Close()
}
