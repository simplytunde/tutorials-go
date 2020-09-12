package main

import (
	"fmt"

	db "github.com/simplytunde/database"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dbConfig,err := db.NewDBConfig("config.yaml")
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/dbname?charset=utf8mb4&parseTime=True&loc=Local", dbUsername, dbPassword, dbHost, dbPort)
	
	// db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
}
