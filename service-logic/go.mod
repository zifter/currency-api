module github.com/zifter/currency/service-logic

replace github.com/zifter/currency/lib => ../lib

require (
	bitbucket.org/liamstask/goose v0.0.0-20150115234039-8488cc47d90c // indirect
	github.com/go-telegram-bot-api/telegram-bot-api v4.6.4+incompatible
	github.com/golang/mock v1.4.0
	github.com/ilyakaznacheev/cleanenv v1.2.0
	github.com/jmoiron/sqlx v1.2.0
	github.com/kylelemons/go-gypsy v0.0.0-20160905020020-08cad365cd28 // indirect
	github.com/lib/pq v1.0.0
	github.com/mattn/go-sqlite3 v1.9.0
	github.com/myusuf3/goose v0.0.0-20191015002438-f2d09616c248 // indirect
	github.com/pressly/goose v2.6.0+incompatible
	github.com/rk/go-cron v0.0.0-20130419213454-4a45dd81c5db
	github.com/sirupsen/logrus v1.2.0
	github.com/spf13/viper v1.6.2
	github.com/stretchr/testify v1.4.0
	github.com/zifter/currency/lib v1.0.0
	github.com/ziutek/mymysql v1.5.4 // indirect
)

go 1.13
