package handler

import (
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

// OpenDBConnection is
func OpenDBConnection() (*gorm.DB, error) {
	db, err := gorm.Open("mysql", "root:@tcp(db:3306)/hoge")
	if err != nil {
		return nil, err
	}
	return db, nil
}

// CloseDBConnction is
func CloseDBConnction() {
	db.Close()
}
