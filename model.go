package main

type Measure struct {
	units []string
	measureType string
	typeEquivalent float64
}

var MeasureReference = []Measure{
	{units: []string{"tsp","teaspoon"}, measureType: "volume", typeEquivalent: 4.92892},
	{units: []string{"tbsp","tablespoon"}, measureType: "volume", typeEquivalent: 14.7868},
	{units: []string{"oz","ounce","fluid ounce","fl oz"}, measureType: "volume", typeEquivalent: 29.5735},
	{units: []string{"cup","c"}, measureType: "volume", typeEquivalent: 240},
	{units: []string{"pint","pt","p"}, measureType: "volume", typeEquivalent: 473.176},
}

var MeasureMap = make(map[string]Measure)

func RegisterMeasure(m Measure){
	for _, r := range m.units {
		MeasureMap[r] = m
	}
}
