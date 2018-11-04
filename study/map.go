package study

import "fmt"

func MapTest() {
	countryCapitalMap := map[string]string{
		"France": "巴黎",
		"Italy":  "罗马",
		"Japan":  "东京",
		"India":  "新德里",
		"China":  "北京",
	}

	for country, capital := range countryCapitalMap {
		fmt.Printf("coutry:%s -> capital:%s\n", country, capital)
	}

	delete(countryCapitalMap, "Japan")

	for country, capital := range countryCapitalMap {
		fmt.Printf("coutry:%s -> capital:%s\n", country, capital)
	}
}
