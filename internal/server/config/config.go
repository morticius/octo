package config

import (
	"github.com/BurntSushi/toml"
	"log"
)

type Config struct {
	Addr            string `toml:"addr"`
	Secret          string `toml:"secret"`
	Issure          string `toml:"issure"`
	BtcRateEndpoint string `toml:"btc_rate_endpoint"`
}

var config *Config

func InitConfig(path string) {
	config = &Config{}
	if _, err := toml.DecodeFile(path, config); err != nil {
		log.Fatalf("cannot load config: %s", err)
	}
}

func Get() *Config {
	return config
}
