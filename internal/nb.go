package internal

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/zifter/currency-api/types"
)

func requestNB(currencyId int) (info *types.NBInfo) {
	resp, err := http.Get("http://www.nbrb.by/API/ExRates/Rates/" + strconv.Itoa(currencyId))
	if err != nil {
		log.Println(err)
		return
	}
	if resp.StatusCode != 200 {
		log.Println("status: ", resp.Status)
		return
	}

	info = &types.NBInfo{}
	err = json.NewDecoder(resp.Body).Decode(&info)
	if err != nil {
		log.Println("Wrong body:\n", err, resp.Body)
		return
	}

	return
}

func RequestNBUSD() *types.NBInfo {
	return requestNB(145)
}

func RequestNBEUR() *types.NBInfo {
	return requestNB(19)
}

func RequestNBRUB() *types.NBInfo {
	return requestNB(141)
}
