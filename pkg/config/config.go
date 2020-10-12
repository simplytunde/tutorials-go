package config

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password "`
}

func (c *Config) LoadConfigFromFile(configFile string) error {
	buf, err := ioutil.ReadFile(configFile)
	if err != nil {
		fmt.Println("Config file not found")
	}
	if err := yaml.Unmarshal(buf, c); err != nil {
		return err
	}
	return nil
}
