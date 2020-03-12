package rub

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/zifter/currency-api/types"
)

type cbrCurrencyInfo struct {
	ID       string  `json:"ID"`
	NumCode  string  `json:"NumCode"`
	CharCode string  `json:"CharCode"`
	Nominal  int     `json:"Nominal"`
	Name     string  `json:"Name"`
	Value    float32 `json:"Value"`
	Previous float32 `json:"Previous"`
}

type cbrResponse struct {
	Valute map[string]cbrCurrencyInfo
}

func Aggregate() (*types.FullCurrencyInfo, error) {
	resp, err := requestCbrData()
	if err != nil {
		return nil, fmt.Errorf("cant request cbr data: %w", err)
	}

	data := types.NewFullCurrencyInfo()
	data.Version = 1
	for k, v := range resp.Valute {
		data.CurrencyAggregation[k] = &types.AggregatedData{
			NationalBank: types.NBInfo{
				ID:           -1,
				Abbreviation: k,
				Name:         v.Name,
				OfficialRate: v.Value,
			},
		}
	}

	return data, nil
}

func requestCbrData() (*cbrResponse, error) {
	resp, err := http.Get("https://www.cbr-xml-daily.ru/daily_json.js")
	if err != nil {
		return nil, fmt.Errorf("failed to get: %w", err)
	}
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("request was failed: %v", resp.Status)
	}

	data := &cbrResponse{}
	err = json.NewDecoder(resp.Body).Decode(data)
	if err != nil {
		return nil, fmt.Errorf("wrong body %v, failed with: %w", resp.Body, err)
	}

	return data, nil
}
