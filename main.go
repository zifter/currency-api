package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/zifter/currency-api/internal"
)

func getCurrency(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Get currency request")
	fullInfo := internal.Aggregate()

	b, err := json.Marshal(fullInfo)
	if err != nil {
		fmt.Fprintf(w, "Something went wrong: %v", err)
	} else {
		fmt.Fprint(w, string(b))
	}
}

func main() {
	log.Println("Currency api")

	config := internal.LoadConfig()

	http.HandleFunc("/currency/", getCurrency)

	log.Println("Start http on: ", config.API.Port)
	err := http.ListenAndServe(":"+config.API.Port, nil)
	if err != nil {
		panic("Something went wrong!")
	}
}
