package handler

import (
	"fmt"
	"github.com/2018-miraikeitai-org/Rakusale-Another-Server/domain/model"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"os"
	"time"
)

// OpenDBConnection is
func OpenDBConnection() *gorm.DB {
	dbType := os.Getenv("DB_TYPE")
	connectionString := os.Getenv("CONNECTION_STRING")
	DB, err := gorm.Open(dbType, connectionString)
	if err != nil {
		panic(err)
	}
	// [Initialize] Database
	DB.AutoMigrate(&model.Vegetable{})
	fmt.Println("[Init] Vegetable Model")
	DB.AutoMigrate(&model.Sell{})
	fmt.Println("[Init] Sell Model")
	DB.AutoMigrate(&model.Buy{})
	fmt.Println("[Init] Buy Model")
	DB.AutoMigrate(&model.SellList{})
	fmt.Println("[Init] SellList Model")
	DB.AutoMigrate(&model.BuyList{})
	fmt.Println("[Init] BuyList Model")
	DB.AutoMigrate(&model.Shop{})
	fmt.Println("[Init] Shop Model")
	DB.AutoMigrate(&model.User{})
	fmt.Println("[Init] User Model")

	return DB
}

// CloseDBConnction is
// func CloseDBConnction() {
// 	DB.Close()
// }
