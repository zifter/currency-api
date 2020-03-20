package types

import "time"

type ScratchingResponse struct {
	Value       float64   `json:"value"`
	Diff        float64   `json:"diff"`
	DiffPercent float64   `json:"diff_percent"`
	Timestamp   time.Time `json:"timestamp"`
}

type BitcoinResponse ScratchingResponse
type EthereumResponse ScratchingResponse
type OilResponse ScratchingResponse
