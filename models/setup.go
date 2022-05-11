package models

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectDB() *gorm.DB {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	db_name := os.Getenv("DB_NAME")
	db_user := os.Getenv("DB_USER")
	db_password := os.Getenv("DB_PASSWORD")

	conn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", db_user, db_password, host, port, db_name)

	db, err := gorm.Open(mysql.Open(conn), &gorm.Config{})

	if err != nil {
		panic("Unable to connect to Database")
	}

	return db
}
