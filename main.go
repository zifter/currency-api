package main

import (
	"fmt"
	"net/http"

	log "github.com/sirupsen/logrus"
	"github.com/zifter/currency-api/internal"
)

func init() {
	http.HandleFunc("/", index)
}

func index(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "currency api")
}

func main() {
	log.Infof("Currency api")

	config := internal.LoadConfig()

	log.Printf("Start http on: localhost:%v", config.API.Port)
	err := http.ListenAndServe(":"+config.API.Port, nil)
	if err != nil {
		log.WithError(err).Panicf("Something went wrong")
	}
}
