package infobank

import (
	"encoding/json"
	"fmt"
	"net/http"
)

var (
	SupportedCurrency = []string{
		"USD",
		"EUR",
		"RUB",
	}
)

type Rate struct {
	Buy  float64
	Sell float64
}

type InfoBankData struct {
	ID         int
	BankID     int
	BankName   string
	UpdateDate string
	UpdateTime string

	Rates map[string]Rate
}

type rawInfoBankData struct {
	ID         int     `json:"FID"`
	BankID     int     `json:"FBANKID"`
	BankName   string  `json:"FBANKNAME"`
	USDBuy     float64 `json:"FRATEBUYNAL1"`
	USDSell    float64 `json:"FRATESELLNAL1"`
	EURBuy     float64 `json:"FRATEBUYNAL2"`
	EURSell    float64 `json:"FRATESELLNAL2"`
	RUBBuy     float64 `json:"FRATEBUYNAL3"`
	RUBSell    float64 `json:"FRATESELLNAL3"`
	UpdateDate string  `json:"FLASTD"`
	UpdateTime string  `json:"FLASTT"`
}

func RequestInfoBankData() ([]InfoBankData, error) {
	raw, err := requestRawInfoBankData()
	if err != nil {
		return nil, err
	}

	data := make([]InfoBankData, len(raw))
	for i := range raw {
		data[i] = InfoBankData{
			ID:         raw[i].ID,
			BankID:     raw[i].BankID,
			BankName:   raw[i].BankName,
			UpdateDate: raw[i].UpdateDate,
			UpdateTime: raw[i].UpdateTime,
			Rates: map[string]Rate{
				"USD": {raw[i].USDBuy, raw[i].USDSell},
				"EUR": {raw[i].EURBuy, raw[i].EURSell},
				"RUB": {raw[i].RUBBuy, raw[i].RUBSell},
			},
		}
	}

	return data, nil
}

func requestRawInfoBankData() ([]rawInfoBankData, error) {
	resp, err := http.Get("https://infobank.by/modules/Ajax/CreateCardTable.aspx?Action=crttbl")
	if err != nil {
		return nil, fmt.Errorf("failed to get: %w", err)
	}
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("request was failed: %v", resp.Status)
	}

	bankInfoList := []rawInfoBankData{}
	err = json.NewDecoder(resp.Body).Decode(&bankInfoList)
	if err != nil {
		return nil, fmt.Errorf("wrong body %v, failed with: %w", resp.Body, err)
	}

	return bankInfoList, nil
}
