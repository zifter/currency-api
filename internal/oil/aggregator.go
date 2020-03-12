package oil

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/zifter/currency-api/types"
)

func Aggregate() (*types.OilResponse, error) {
	req, err := http.NewRequest("GET", "https://ru.investing.com/commodities/brent-oil", nil)
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

	f := strings.Replace(data.Find("#last_last").Text(), ",", ".", -1)
	v, _ := strconv.ParseFloat(f, 32)
	t := data.Find(".pid-8833-time").Text()

	return &types.OilResponse{
		Value:     float32(v),
		Timestamp: t,
	}, nil
}
