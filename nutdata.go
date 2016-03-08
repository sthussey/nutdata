// The API server handles incoming requests for nutrition data
package main

var FractionRunes = map[float64]rune {
	.25: 188,
	.50: 189,
	.75: 190,
	.33: 8531,
	.66: 8532,
	.20: 8533,
	.40: 8534,
	.60: 8535,
	.80: 8536,
	.17: 8537,
	.83: 8538,
	.13: 8539,
	.38: 8540,
	.63: 8541,
	.88: 8542,
}

func main(){
	initializeConfig()
	initializeDatastore()
	for _, u := range UnitReference {
		RegisterUnit(u)
	}
	startService()
}
