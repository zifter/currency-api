package byn

import (
	"fmt"
	"sync"

	"github.com/sirupsen/logrus"
	"github.com/zifter/currency-api/internal/byn/infobank"
	"github.com/zifter/currency-api/internal/byn/national_bank"
	"github.com/zifter/currency-api/types"
)

var log = logrus.New().WithFields(logrus.Fields{
	"name": "currency-api",
})

func Aggregate() (*types.FullCurrencyInfo, error) {
	data := types.NewFullCurrencyInfo()

	wg := sync.WaitGroup{}
	wg.Add(len(currencies))
	for i := range currencies {
		data.CurrencyAggregation[currencies[i].humanReadable] = &types.AggregatedData{}

		go func(cur *currencyTechDescr) {
			defer wg.Done()

			nb, err := request(cur)
			if err != nil {
				log.Errorf("Failed while requsting %v: %v", cur, err)
			} else {
				data.CurrencyAggregation[cur.humanReadable].SetNBInfo(nb)
			}
		}(&currencies[i])
	}

	wg.Add(1)

	infoBank := []infobank.InfoBankData{}
	go func() {
		defer wg.Done()
		var err error
		infoBank, err = infobank.RequestInfoBankData()
		if err != nil {
			log.Errorf("Failed to read infobank %v", err)
		}
	}()

	wg.Wait()

	if len(infoBank) > 1 {
		for _, name := range infobank.SupportedCurrency {
			sellMin := 0
			buyMax := 0
			for i := range infoBank {
				if infoBank[i].Rates[name].Sell > infoBank[sellMin].Rates[name].Sell {
					sellMin = i
				}

				if infoBank[i].Rates[name].Buy > infoBank[buyMax].Rates[name].Buy {
					buyMax = i
				}
			}

			data.CurrencyAggregation[name].SetBankBest(&types.BestInfo{
				Sell: types.Rate{
					BankName: infoBank[sellMin].BankName,
					Value:    infoBank[sellMin].Rates[name].Sell,
				},
				Buy: types.Rate{
					BankName: infoBank[buyMax].BankName,
					Value:    infoBank[buyMax].Rates[name].Buy,
				},
			})
		}
	}

	return data, nil
}

func request(descr *currencyTechDescr) (*types.NBInfo, error) {
	info, err := national_bank.RequestInfo(descr.nationalBankID)
	if err != nil {
		return nil, fmt.Errorf("failed to get national bank: %w", err)
	}

	return info, nil
}
