package internal

type currencyTechDescr struct {
	humanReadable  string
	nationalBankID int
}

var (
	currencies = []currencyTechDescr{
		{
			humanReadable:  "USD",
			nationalBankID: 145,
		},
		{
			humanReadable:  "EUR",
			nationalBankID: 292,
		},
		{
			humanReadable:  "RUB",
			nationalBankID: 298,
		},
	}
)
