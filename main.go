package main

import (
	"fmt"
	"log"

	db "github.com/simplytunde/tutorials-go/pkg/database"
)

func main() {
	var dbConfig *db.DBConfig
	var err error
	if dbConfig, err = db.NewDBConfigFromFile("config.yaml"); err != nil {
		print(dbConfig)
		log.Fatal("There is a problem with dbConfig")
	}
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/dbname?charset=utf8mb4&parseTime=True&loc=Local", dbConfig.Username, dbConfig.Password, dbConfig.Host, dbConfig.Port)
	fmt.Println(dsn)
	// db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
}
