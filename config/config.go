package config

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"crudgoeg/models"
)

var DB *gorm.DB

func MigrateDB() {
	DB.AutoMigrate(&models.User{})
}

func InitDB() {
	connectionString := "root:BavarianRestu2002@tcp(127.0.0.1:3306)/user?charset=utf8mb4&parseTime=True&loc=Local"

	var err error
	DB, err = gorm.Open(mysql.Open(connectionString), &gorm.Config{})

	if err != nil {
		panic(err)
	}

}

