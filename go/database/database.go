package database

import (
	"log"
	"os"
    "strconv"
	"gopkg.in/yaml.v2"
)

type DBConfig struct {
	Host     string `yaml:"database:host"`
	Port     int `yaml:"database:port"`
	Username string `yaml:"database:username"`
	Password string
}

func NewDBConfig(configFile string) (*DBConfig, error) {
	var dbConfig *DBConfig
	var err error
	if _, err := os.Stat(configFile); err != nil {
		dbConfig = &DBConfig{}
	}
	if dbConfig, err = newDBConfigFromFile(configFile); err != nil {
		return nil, err
	}

	if Host, present := os.LookupEnv("DB_HOST"); dbConfig.Host == "" && present == false {
		log.Fatal("Please enter the database host")
	} else if dbConfig.Host == "" {
		dbConfig.Host = Host
	}
	if Username, present := os.LookupEnv("DB_USERNAME"); dbConfig.Username == "" && present == false {
		log.Fatal("Please enter the database username")
	} else if dbConfig.Username == "" {
		dbConfig.Username = Username
	}
	if Password, present := os.LookupEnv("DB_PASSWORD"); dbConfig.Password == "" && present == false {
		log.Fatal("Please enter the database password")
	} else if dbConfig.Password == "" {
		dbConfig.Password = Password
	}
	if Port, present := os.LookupEnv("DB_PORT"); dbConfig.Port == 0 && present == false {
		dbConfig.Port = 3306
	} else if dbConfig.Port == 0 {
		if PortInt, err := strconv.Atoi(Port); err == nil {
			dbConfig.Port = PortInt
		}
	}
	return dbConfig, nil
}
func newDBConfigFromFile(configFile string) (*DBConfig, error) {
	dbConfig := DBConfig{}
	var file *os.File
	if file, err := os.Open(configFile); err != nil {
		log.Fatal("Config file was not found")
	}else{
		defer file.Close()
	}

	d := yaml.NewDecoder(file)
	if err := d.Decode(&dbConfig); err != nil {
		return nil, err
	}
	return &dbConfig, nil
}
