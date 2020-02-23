package lib

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
	CurrencyAggregation map[string]AggregatedData
}

func NewFullCurrencyInfo() *FullCurrencyInfo {
	return &FullCurrencyInfo{map[string]AggregatedData{}}
}
