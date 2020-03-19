package investingcom

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

type privateData struct {
	url         string
	diff        string
	diffPercent string
}

var available = map[string]privateData{
	"brent-oil": privateData{
		url:         "https://ru.investing.com/commodities/brent-oil",
		diff:        "span.arial_20.pid-8833-pc",
		diffPercent: "span.arial_20.pid-8833-pcp",
	},
	"bitcoin-usd": privateData{
		url:         "https://ru.investing.com/crypto/bitcoin/btc-usd",
		diff:        "span.arial_20.pid-945629-pc",
		diffPercent: "span.arial_20.pid-945629-pcp",
	},
	"ethereum-usd": privateData{
		url:         "https://ru.investing.com/crypto/ethereum",
		diff:        "span.arial_20.pid-1061443-pc",
		diffPercent: "span.arial_20.pid-1061443-pcp",
	},
	"tesla-usd": privateData{
		url:         "https://ru.investing.com/equities/tesla-motors",
		diff:        "span.arial_20.pid-13994-pc",
		diffPercent: "span.arial_20.pid-13994-pcp",
	},
}

func strToFloat64(text string) float64 {
	text = strings.Replace(text, ".", "", -1)
	text = strings.Replace(text, ",", ".", -1)
	text = strings.Replace(text, "%", "", -1)
	v, _ := strconv.ParseFloat(text, 64)
	return v
}

func Aggregate(dataType string) (*ScratchResponse, error) {
	investingData := available[dataType]
	req, err := http.NewRequest("GET", investingData.url, nil)
	if err != nil {
		return nil, fmt.Errorf("cant create request: %w", err)
	}
	req.Header.Set("User-Agent", "Mozilla/5.0")
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("cant perform request: %w", err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		return nil, fmt.Errorf("cant get source because of: %v", res.Status)
	}

	data, err := goquery.NewDocumentFromResponse(res)
	if err != nil {
		return nil, fmt.Errorf("cant get page from external site %v", err)
	}

	return &ScratchResponse{
		Value:       strToFloat64(data.Find("#last_last").Text()),
		Diff:        strToFloat64(data.Find(investingData.diff).Text()),
		DiffPercent: strToFloat64(data.Find(investingData.diffPercent).Text()),
		Timestamp:   time.Now(),
	}, nil
}
