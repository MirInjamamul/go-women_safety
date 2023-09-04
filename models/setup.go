package models

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {

	// dsn := fmt.Sprintf(
	// 	"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
	// 	"localhost",
	// 	"postgres",
	// 	"postgres",
	// 	"safety",
	// 	"5432",
	// )

	dsn := "root:workspacify2023@tcp(127.0.0.1:3306)/safety?charset=utf8mb4&parseTime=True&loc=Local"

	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to Connect to Database")
	}

	database.AutoMigrate(&User{})

	DB = database
}
