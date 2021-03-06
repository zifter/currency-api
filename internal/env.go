package internal

import (
	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	API struct {
		Port string `yaml:"port" env:"PORT" env-default:"8082"`
	} `yaml:"logic"`
}

func LoadConfig() *Config {
	cfg := &Config{}
	err := cleanenv.ReadEnv(cfg)
	if err != nil {
		panic(err)
	}
	return cfg
}
