package internal

import "github.com/zifter/currency-api/types"

func Aggregate() (data *types.FullCurrencyInfo) {
	data = types.NewFullCurrencyInfo()

	var usd types.AggregatedData
	usd.NationalBank = *RequestNBUSD()
	data.CurrencyAggregation["usd"] = usd

	//infoBankData := RequestInfoBankData()
	//if len(infoBankData) == 0 {
	//	log.Fatal("Can't aggregate")
	//	return
	//}
	//
	//sort.Slice(infoBankData[:], func (l, r int) bool {
	//	return infoBankData[l].USDBuy > infoBankData[r].USDBuy
	//})
	//data.USD.Buy.BankName = infoBankData[0].BankName
	//data.USD.Buy.Value = infoBankData[0].USDBuy
	//
	//sort.Slice(infoBankData[:], func (l, r int) bool {
	//	return infoBankData[l].USDSell < infoBankData[r].USDSell
	//})
	//var i int = 0
	//for {
	//	if  infoBankData[i].USDSell > 0 {
	//		data.USD.Sell.BankName = infoBankData[i].BankName
	//		data.USD.Sell.Value = infoBankData[i].USDSell
	//		break
	//	} else {
	//		i++
	//	}
	//}

	//data.USD.Sell.BankName = firstEl.BankName
	//data.EUR.Buy.BankName = firstEl.BankName
	//data.EUR.Sell.BankName = firstEl.BankName
	//data.RUB.Buy.BankName = firstEl.BankName
	//data.RUB.Sell.BankName = firstEl.BankName
	//
	//
	//data.USD.Sell.Value = firstEl.USDSell
	//data.EUR.Buy.Value = firstEl.EURBuy
	//data.EUR.Sell.Value = firstEl.EURSell
	//data.RUB.Buy.Value = firstEl.RUBBuy
	//data.RUB.Sell.Value = firstEl.RUBSell

	return
}
