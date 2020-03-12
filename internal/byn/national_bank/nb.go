package national_bank

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/zifter/currency-api/types"
)

func RequestInfo(currencyId int) (*types.NBInfo, error) {
	resp, err := http.Get("http://www.nbrb.by/API/ExRates/Rates/" + strconv.Itoa(currencyId))
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("request was failed: %v", resp.Status)
	}

	info := &types.NBInfo{}
	err = json.NewDecoder(resp.Body).Decode(&info)
	if err != nil {
		return nil, fmt.Errorf("wrong body %v, failed with: %w", resp.Body, err)
	}

	return info, nil
}
