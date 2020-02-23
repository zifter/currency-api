package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/zifter/currency/lib"
	"github.com/zifter/currency/service-api/api"
)

func getCurrency(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Get currency request")
	fullInfo := api.Aggregate()

	b, err := json.Marshal(fullInfo)
	if err != nil {
		fmt.Fprintf(w, "Something went wrong: %v", err)
	} else {
		fmt.Fprint(w, string(b))
	}
}

func main() {
	log.Println("Currency api")

	config := lib.LoadConfig()

	http.HandleFunc("/currency/", getCurrency)

	log.Println("Start http on: ", config.APIPort)
	err := http.ListenAndServe(config.APIHost+":"+config.APIPort, nil)
	if err != nil {
		panic("Something went wrong!")
	}
}
