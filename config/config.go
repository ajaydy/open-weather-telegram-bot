package config

import (
	"fmt"
	"github.com/BurntSushi/toml"
)

type (
	Config struct {
		PublicUrl       string `toml:"public_url"`
		Token           string `toml:"token"`
		Port            string `toml:"port"`
		OpenweatherURL  string `toml:"openweather_url"`
		LocationiqToken string `toml:"locationiq_token"`
		ApiKey          string `toml:"api_key"`
	}
)

func Init(configFile string) (Config, error) {
	var config Config
	if _, err := toml.DecodeFile(configFile, &config); err != nil {
		fmt.Println(err)
		return Config{}, err
	}

	return config, nil
}
