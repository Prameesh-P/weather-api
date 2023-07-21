package database

import (
	// "database/sql"
	
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func DBConnection() {
	connStr := os.Getenv("DSN")
	DB, err = gorm.Open(postgres.Open(connStr),&gorm.Config{})
	if err!=nil{
		log.Printf("cat't connect database got an error %v",err)
	}
}
