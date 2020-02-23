package lib

import (
	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	API struct {
		Port string `yaml:"port" env:"API_PORT" env-default:"8080"`
		Host string `yaml:"host" env:"API_HOST" env-default:"localhost"`
	} `yaml:"logic"`

	Logic struct {
		Port string `yaml:"port" env:"LOGIC_PORT" env-default:"8081"`
		Host string `yaml:"host" env:"LOGIC_HOST" env-default:"localhost"`
	} `yaml:"logic"`

	Chatbot struct {
		Port string `yaml:"port" env:"CHATBOT_PORT" env-default:"8082"`
		Host string `yaml:"host" env:"CHATBOT_HOST" env-default:"localhost"`

		TelegramBotToken string `yaml:"telegram_bot_token" env-upd:"" env:"TELEGRAM_BOT_TOKEN"`
		TelegramGroupID  int64  `yaml:"telegram_group_id" env-upd:"" env:"TELEGRAM_GROUP_ID"`
	} `yaml:"chatbot"`
}

func LoadConfig() *Config {
	cfg := &Config{}
	err := cleanenv.ReadConfig("secrets.yml", cfg)
	if err != nil {
		panic(err)
	}
	cleanenv.UpdateEnv(cfg)
	return cfg
}
