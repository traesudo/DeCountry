package DeCountry

import (
	"DetermineCountry/source"
	"github.com/tidwall/gjson"
	"strings"
)

func IsCountryNameE(str string) bool {
	str1 := strings.ToLower(str)
	result := gjson.Get(source.StringInfo, "..#.nameEN")
	for _, name := range result.Array() {
		if str1 == name.Str {
			return true
		}
	}
	return false
}
func IsCountryCode(str string) bool {
	str1 := strings.ToLower(str)
	result := gjson.Get(source.StringInfo, "..#.code")
	for _, name := range result.Array() {
		if str1 == name.Str {
			return true
		}
	}
	return false
}

func IsCountryNameC(str string) bool {
	str1 := strings.ToLower(str)
	result := gjson.Get(source.StringInfo, "..#.nameZH")
	for _, name := range result.Array() {
		if str1 == name.Str {
			return true
		}
	}
	return false
}
