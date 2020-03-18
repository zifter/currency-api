package types

import "time"

type BitcoinResponse struct {
	Value       float64   `json:"value"`
	Diff        float64   `json:"diff"`
	DiffPercent float64   `json:"diff_percent"`
	Timestamp   time.Time `json:"timestamp"`
}
