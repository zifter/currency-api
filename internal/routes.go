package internal

import (
	"encoding/json"
	"fmt"
	"net/http"

	log "github.com/sirupsen/logrus"
	"github.com/zifter/currency-api/internal/byn"
	"github.com/zifter/currency-api/internal/investingcom"
	"github.com/zifter/currency-api/internal/rub"
)

func init() {
	// backward
	http.HandleFunc("/currency/", getBYNCurrency)
	http.HandleFunc("/currency/byn/", getBYNCurrency)

	http.HandleFunc("/currency/rub/", getRUBCurrency)

	http.HandleFunc("/oil/", getOil)
	http.HandleFunc("/bitcoin/", getBitcoin)
	http.HandleFunc("/ethereum/", getEthereum)
}

func getBYNCurrency(w http.ResponseWriter, req *http.Request) {
	log.Infof("Get %v", req.URL.Path)

	fullInfo, err := byn.Aggregate()
	if err != nil {
		log.Errorf("cant aggregate: %v", err)
	}
	b, err := json.Marshal(fullInfo)
	if err != nil {
		fmt.Fprintf(w, "cant aggregate: %v", err)
	} else {
		fmt.Fprint(w, string(b))
	}
}

func getRUBCurrency(w http.ResponseWriter, req *http.Request) {
	log.Infof("Get %v", req.URL.Path)

	fullInfo, err := rub.Aggregate()
	if err != nil {
		log.Errorf("cant aggregate: %v", err)
	}

	b, err := json.Marshal(fullInfo)
	if err != nil {
		fmt.Fprintf(w, "cant aggregate: %v", err)
	} else {
		fmt.Fprint(w, string(b))
	}
}

func getDataTypeResponse(dataType string, w http.ResponseWriter, req *http.Request) {
	log.Infof("Get %v", req.URL.Path)

	fullInfo, err := investingcom.Aggregate(dataType)
	if err != nil {
		log.Errorf("cant aggregate: %v", err)
	}

	b, err := json.Marshal(fullInfo)
	if err != nil {
		fmt.Fprintf(w, "cant aggregate: %v", err)
	} else {
		fmt.Fprint(w, string(b))
	}
}

func getOil(w http.ResponseWriter, req *http.Request) {
	getDataTypeResponse("brent-oil", w, req)
}

func getBitcoin(w http.ResponseWriter, req *http.Request) {
	getDataTypeResponse("bitcoin-usd", w, req)
}

func getEthereum(w http.ResponseWriter, req *http.Request) {
	getDataTypeResponse("ethereum-usd", w, req)
}
