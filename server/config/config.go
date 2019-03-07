package config

import (
	"encoding/json"
	"io/ioutil"
)

type Config struct {
	Database Database `json:"Database"`
	Smite    Smite
}

type Database struct {
	Host     string `json:Host`
	Name     string `json:Name`
	Username string `json:Username`
	Password string `json:Password`
}

type Smite struct {
	DevId   string `json:DevId`
	AuthKey string `json:AuthKey`
}

func NewConfig(path string) (*Config, error) {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	config := new(Config)

	err = json.Unmarshal([]byte(file), config)
	if err != nil {
		return nil, err
	}
	return config, nil
}
