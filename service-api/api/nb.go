package api

import (
	"encoding/json"
	"github.com/zifter/currency/lib"
	"log"
	"net/http"
	"strconv"
)

func requestNB(currencyId int) (info *lib.NBInfo) {
	resp, err := http.Get("http://www.nbrb.by/API/ExRates/Rates/" + strconv.Itoa(currencyId))
	if err != nil {
		log.Println(err)
		return
	}
	if resp.StatusCode != 200 {
		log.Println("status: ", resp.Status)
		return
	}

	info = &lib.NBInfo{}
	err = json.NewDecoder(resp.Body).Decode(&info)
	if err != nil {
		log.Println("Wrong body:\n", err, resp.Body)
		return
	}

	return
}

func RequestNBUSD() *lib.NBInfo  {
	return requestNB(145)
}

func RequestNBEUR() *lib.NBInfo  {
	return requestNB(19)
}

func RequestNBRUB() *lib.NBInfo  {
	return requestNB(141)
}