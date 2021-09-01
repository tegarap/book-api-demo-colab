package config

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"book-api-demo/models"
)

var Db *gorm.DB

func init() {

	var err error
	dsn := "root:12345@tcp(127.0.0.1:3306)/alterra?charset=utf8mb4&parseTime=True&loc=Local"
	Db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	Db.AutoMigrate(&models.Users{})

}
