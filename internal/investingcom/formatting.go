package investingcom

import (
	"fmt"
	"strings"

	"github.com/zifter/currency-api/stringtofloat"
)

func strToFloat64(initialText string) (float64, error) {
	fmt.Println(initialText)

	text := strings.Replace(initialText, "%", "", -1)
	text = strings.Replace(text, "+", "", -1)

	return stringtofloat.Convert(text)
}