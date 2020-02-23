package types

import "time"

type RatePost struct {
	ID   string `db:"id"`
	Date string `db:"postDate"`
	Msg  string `db:"msg"`
}

func NewRatePost(msg string) *RatePost {
	return &RatePost{
		Date: time.Now().Format("2006-01-02"),
		Msg:  msg,
	}
}
