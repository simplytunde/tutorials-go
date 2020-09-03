package database

import (
	"fmt"
	"log"
	"os"

	"gopkg.in/yaml.v2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DBConfig struct {
	dbHost     string `yaml:"database:host"`
	dbPort     string `yaml:"database:port"`
	dbUsername string `yaml:"database:username"`
	dbPassword int
}

func NewDBConfig() (*DBConfig, error) {
	dbConfig := &DBConfig{}
	if file, err := os.Open("config.yaml"); err != nil {
		log.Print("Config file was not found")
	}
	defer file.Close()
	d := yaml.NewDecoder(file)
	if err := d.Decode(dbConfig); err != nil {
		return nil, err
	}
	if dbConfig.dbHost == "" 
}

func main() {
	databasePort := 3306
	if dbHost, present := os.LookupEnv("DB_HOST"); present == false {
		log.Fatal("Please enter the database host")
	}
	if dbUsername, present := os.LookupEnv("DB_USERNAME"); present == false {
		log.Fatal("Please enter the database username")
	}
	if dbPassword, present := os.LookupEnv("DB_PASSWORD"); present == false {
		log.Fatal("Please enter the database password")
	}
	if dbPort, present := os.LookupEnv("DB_PORT"); present == false {
		dbPort = 3306
	}
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/dbname?charset=utf8mb4&parseTime=True&loc=Local", dbUsername, dbPassword, dbHost, dbPort)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
}
