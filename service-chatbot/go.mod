module github.com/zifter/currency/service-chatbot

replace github.com/zifter/currency/lib => ../lib

require (
	github.com/go-telegram-bot-api/telegram-bot-api v4.6.4+incompatible
	github.com/golang/mock v1.4.0
	github.com/ilyakaznacheev/cleanenv v1.2.0
	github.com/rk/go-cron v0.0.0-20130419213454-4a45dd81c5db
	github.com/sirupsen/logrus v1.2.0
	github.com/spf13/viper v1.6.2
	github.com/stretchr/testify v1.4.0
	github.com/zifter/currency/lib v0.0.0-00010101000000-000000000000
)

go 1.13
