package config

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db *gorm.DB
)

func ConnectToDB() {
	fmt.Println("Connecting to database...")
	d, err := gorm.Open("mysql", "skincare_user:skincare2022@tcp(127.0.0.1:3306)/skincare_ingredients?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("Connection to database failed. " + err.Error())
	} else {
		fmt.Println("Successfully connected to DB.")
	}
	db = d

	// defer db.Close()
}

func GetDB() *gorm.DB {
	return db
}
