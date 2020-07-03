package stringtofloat

import (
	"strconv"
	"strings"
)

func normalizeEurope(old string) string {
	count := strings.Count(old, ".")
	s := strings.Replace(old, ",", ".", -1)
	return strings.Replace(s, ".", "", count)

}
func normalizeAmericanBritain(old string) string {
	return strings.Replace(old, ",", "", -1)
}

// Convert determines locale there are really only two types, US/Britain and rest of Europe
func Convert(fs string) (float64, error) {
	point := strings.LastIndex(fs, ".")
	comma := strings.LastIndex(fs, ",")

	text := fs

	isEurope := point > comma

	if isEurope {
		text = normalizeAmericanBritain(text)
	} else {
		text = normalizeEurope(text)
	}

	return strconv.ParseFloat(text, 64)
}
