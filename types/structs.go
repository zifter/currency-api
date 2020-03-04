package types

import (
	"encoding/json"
	"errors"
)

type Rate struct {
	BankName string
	Value    float32
}

type BestInfo struct {
	Sell Rate
	Buy  Rate
}

type NBInfo struct {
	ID           int     `json:"Cur_ID"`
	Abbreviation string  `json:"Cur_Abbreviation"`
	Name         string  `json:"Cur_Name"`
	OfficialRate float32 `json:"Cur_OfficialRate"`
}

type AggregatedData struct {
	NationalBank NBInfo
	BankBest     BestInfo
}

type FullCurrencyInfo struct {
	Version             int32
	CurrencyAggregation map[string]*AggregatedData
}

func NewFullCurrencyInfo() *FullCurrencyInfo {
	return &FullCurrencyInfo{0, map[string]*AggregatedData{}}
}

func (data *AggregatedData) SetNBInfo(info *NBInfo) {
	data.NationalBank = *info
}

func (data *AggregatedData) SetBankBest(info *BestInfo) {
	data.BankBest = *info
}

func (info *FullCurrencyInfo) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}

	return json.Unmarshal(b, &info)
}
