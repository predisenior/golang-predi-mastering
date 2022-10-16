package config

import (
	// "rest/model"

	// model "project-3/models"
	model "project-3/graph/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)
var DB *gorm.DB

func KonekDB(){
	dsn := "root:predator@tcp(127.0.0.1:3306)/go_books?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if (err!=nil){
		panic(err.Error())
	}
	initMigrate()
}

func initMigrate(){
	DB.AutoMigrate(&model.Book{})
}
