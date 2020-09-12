package database

import (
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

type DBConfig struct {
	dbHost     string `yaml:"database:host"`
	dbPort     string `yaml:"database:port"`
	dbUsername string `yaml:"database:username"`
	dbPassword int
}

func NewDBConfig(configFile string) (DBConfig, error) {
	var dbConfig DBConfig
	if file, err := os.Stat(configFile); err != nil {
		dbConfig = DBConfig{}
	}
	if config, err := newDBConfigFromFile(configFile); err != nil {
		return DBConfig{}, err
	}

	if dbHost, present := os.LookupEnv("DB_HOST"); dbConfig.dbHost == nil && present == false {
		log.Fatal("Please enter the database host")
	} else if dbConfig.dbHost == nil {
		dbConfig.dbHost = dbHost
	}
	if dbUsername, present := os.LookupEnv("DB_USERNAME"); dbConfig.dbUsername == nil && present == false {
		log.Fatal("Please enter the database username")
	} else if dbConfig.dbUsername == nil {
		dbConfig.dbUsername = dbUsername
	}
	if dbPassword, present := os.LookupEnv("DB_PASSWORD"); dbConfig.dbPassword == nil && present == false {
		log.Fatal("Please enter the database password")
	} else if dbConfig.dbPassword == nil {
		dbConfig.dbPassword = dbPassword
	}
	if dbPort, present := os.LookupEnv("DB_PORT"); dbConfig.dbPort == nil && present == false {
		dbPort = 3306
	} else if dbConfig.dbPort == nil {
		dbConfig.dbPort = dbPort
	}
	return dbConfig, nil
}
func newDBConfigFromFile(configFile string) (DBConfig, error) {
	dbConfig := DBConfig{}
	if file, err := os.Open(configFile); err != nil {
		log.Print("Config file was not found")
	}
	defer file.Close()
	d := yaml.NewDecoder(file)
	if err := d.Decode(&dbConfig); err != nil {
		return nil, err
	}
	return dbConfig, nil
}
