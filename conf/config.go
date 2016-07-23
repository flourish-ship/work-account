package conf

import (
	"encoding/json"
	"io/ioutil"
)

//Config struct
type Config struct {
	API *APIConfig `json:"api"`
	DB  *DBConfig  `json:"db"`
}

// APIConfig ...
type APIConfig struct {
	Port  string      `json:"port"`
	Redis RedisConfig `json:"redis"`
}

//RedisConfig ...
type RedisConfig struct {
	Addr     string `json:"addr"`
	Password string `json:"password"`
	Database string `json:"database"`
}

//DBConfig ...
type DBConfig struct {
	DriverName string `json:"driver"`
	Addr       string `json:"addr"`
	Username   string `json:"username"`
	Password   string `json:"password"`
	Database   string `json:"database"`
}

//LoadConfig ...
func LoadConfig(path string, c *Config) error {

	file, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}

	err = json.Unmarshal(file, c)
	if err != nil {
		return err
	}

	return nil
}
