package main

import (
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/rk/go-cron"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/zifter/currency/lib"
	repository "github.com/zifter/currency/service-logic/internal/repository"
	service "github.com/zifter/currency/service-logic/internal/service"
)

var log = logrus.New().WithFields(logrus.Fields{
	"service-name": "logic",
})

func main() {
	log.Info("Chat bot logic")
	viper.SetDefault("DATABASE_URL", "postgres://postgres:changeme@localhost/postgres?sslmode=disable")

	config := lib.LoadConfig()
	dbURL := viper.GetString("DATABASE_URL")
	conn, err := sqlx.Connect("postgres", dbURL)
	if err != nil {
		panic(err)
	}

	addr := fmt.Sprintf("http://%s:%s", config.Chatbot.Host, config.Chatbot.Port)
	w := &service.CurrencyWatcher{
		CurrencyAPI: fmt.Sprintf("http://%s:%s", config.API.Host, config.API.Port),
		Chat:        service.NewChatGateImpl(addr),
		Repo:        repository.NewDBRepo(conn),
	}

	cron.NewDailyJob(cron.ANY, cron.ANY, cron.ANY, func(time.Time) {
		log.Info("Trigger update check")
		err := w.CheckUpdate()
		if err != nil {
			log.Errorf("Failed to check update: %v", err)
		}
	})

	for {

	}
}
