package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
	"os"
)

var db *gorm.DB

func Init() {
	dbUri := GetDbUri()
	conn, err := gorm.Open("postgres", dbUri)
	if err != nil {
		panic("failed to connect database")
	}
	db = conn
	db.Debug().AutoMigrate(&Hashes{}) // Add .AutoMigrate(&Hashes{}) for migration
}

func GetDbUri() string {
	e := godotenv.Load()
	if e != nil {
		fmt.Print(e)
	}

	// add credentials from .env
	username := os.Getenv("db_user")
	password := os.Getenv("db_pass")
	dbName := os.Getenv("db_name")
	dbHost := os.Getenv("db_host")

	dbUri := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s", dbHost, username, dbName, password)
	return dbUri
}
