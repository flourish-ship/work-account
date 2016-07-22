package conf

import (
	"encoding/json"
	"io/ioutil"
)

//Config struct
type Config struct {
	Redis RedisConfig `json:"redis"`
}

//RedisConfig ...
type RedisConfig struct {
	Addr     string `json:"addr"`
	Password string `json:"password"`
	Database string `json:"database"`
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
