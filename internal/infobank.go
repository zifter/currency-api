package internal

import (
	"encoding/json"
	"log"
	"net/http"
)

type InfoBankData struct {
	ID         int     `json:"FID"`
	BankID     int     `json:"FBANKID"`
	BankName   string  `json:"FBANKNAME"`
	USDBuy     float32 `json:"FRATEBUYNAL1"`
	USDSell    float32 `json:"FRATESELLNAL1"`
	EURBuy     float32 `json:"FRATEBUYNAL2"`
	EURSell    float32 `json:"FRATESELLNAL2"`
	RUBBuy     float32 `json:"FRATEBUYNAL3"`
	RUBSell    float32 `json:"FRATESELLNAL3"`
	UpdateDate string  `json:"FLASTD"`
	UpdateTime string  `json:"FLASTT"`
}

func RequestInfoBankData() (bankInfoList []InfoBankData) {
	resp, err := http.Get("https://infobank.by/modules/Ajax/CreateCardTable.aspx?Action=crttbl")
	if err != nil {
		log.Println(err)
		return
	}
	if resp.StatusCode != 200 {
		log.Println("status: ", resp.Status)
		return
	}

	err = json.NewDecoder(resp.Body).Decode(&bankInfoList)
	if err != nil {
		log.Println("Wrong body:\n", err, resp.Body)
		return
	}

	return
}
