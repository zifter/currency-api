package main

import (
	"net/http"

	log "github.com/sirupsen/logrus"
	"github.com/zifter/currency-api/internal"
)

func main() {
	log.Infof("Currency api")

	config := internal.LoadConfig()

	log.Println("Start http on: ", config.API.Port)
	err := http.ListenAndServe(":"+config.API.Port, nil)
	if err != nil {
		log.Fatalf("Something went wrong: %v", err)
	}
}
