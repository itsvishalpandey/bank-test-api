package config

import (
	model "bank-test-api/Models"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {

	// username := os.Getenv("DB_USERNAME")
	// password := os.Getenv("DB_PASSWORD")
	// host := os.Getenv("DB_HOST")
	// port := os.Getenv("DB_PORT")
	// dbName := os.Getenv("DB_NAME")

	username := "vishal"
	password := "VishalP@1601"
	host := "127.0.0.1"
	port := "3306"
	dbName := "bank_store"

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", username, password, host, port, dbName)

	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Error while connecting DB: " + err.Error())
	}

	if err := DB.AutoMigrate(&model.BankMaster{}); err != nil {
		panic("Error while Migrating DB: " + err.Error())
	}

	fmt.Println("Successfully connected to DB")

}
