package types

import "time"

type OilResponse struct {
	Value       float64   `json:"value"`
	Diff        float64   `json:"diff"`
	DiffPercent float64   `json:"diff_percent"`
	Legacy      string `json:"timestamp"`
	Timestamp  time.Time `json:"timestamp2"`
}
