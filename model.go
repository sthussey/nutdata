package main

import (
	"fmt"
	"math"
)

// A Unit is a measure or weight unit including a slice of
// equivalent labels (the first member is the default label),
// the type of measure ('volume' or 'mass')
// and the type equivalency for the unit. Type equivalency
// is the amount of reference units in this unit. Reference
// units are milliliters for volume and grams for mass.
type Unit struct {
	Labels []string `json:"labels"`
	MeasureType string `json:"measureType,omitempty"`
	TypeEquivalent float64 `json:"typeEquivalent"`
}

// A FoodUnit is a description of colloqial measures for a particluar
// food and a equivalent weight in grams
type FoodUnit struct {
	Unit
	FoodId string `json:"foodID"`
}

// A Measure is a measurement amount with an assigned
// unit.
type Measure struct {
  Unit	`json:"unit"`
  Amount float64 `json:"amount"`
}


// This struct represents a single record of the USDA
// Weight File as described at the below. Some fields are omitted.
// http://www.ars.usda.gov/SP2UserFiles/Place/80400525/Data/SR/SR28/sr28_doc.pdf
type USDAFoodWeight struct {
	FoodID string `json:"foodID"`
	Amount int32 `json:"amount"`
	Measure string `json:"measureLabel"` //label or description
	MassEq float64	`json:"massEq"` //mass of the described measure, in grams
}	

var UnitReference = []Unit{
	{Labels: []string{"tsp","teaspoon"}, MeasureType: "volume", TypeEquivalent: 4.92892},
	{Labels: []string{"tbsp","tablespoon"}, MeasureType: "volume", TypeEquivalent: 14.7868},
	{Labels: []string{"fl oz","fluid ounce"}, MeasureType: "volume", TypeEquivalent: 29.5735},
	{Labels: []string{"cup","c"}, MeasureType: "volume", TypeEquivalent: 236.5882365},
	{Labels: []string{"pint","pt","p"}, MeasureType: "volume", TypeEquivalent: 473.176},
	{Labels: []string{"l","litre"}, MeasureType: "volume", TypeEquivalent: 1000},
	{Labels: []string{"ml","mils","millilitre"}, MeasureType: "volume", TypeEquivalent: 1},
	{Labels: []string{"oz","ounce"}, MeasureType: "mass", TypeEquivalent: 28.3495},
	{Labels: []string{"lb","pound"}, MeasureType: "mass", TypeEquivalent: 453.592},
	{Labels: []string{"stone"}, MeasureType: "mass", TypeEquivalent: 6350.29},
	{Labels: []string{"kg","kilogram"}, MeasureType: "mass", TypeEquivalent: 1000},
	{Labels: []string{"g","gram"}, MeasureType: "mass", TypeEquivalent: 1},
}

var UnitMap = make(map[string]Unit)

// RegisterUnit adds a Unit struct to the Map of all valid units
func RegisterUnit(u Unit){
  
	for _, r := range u.Labels {
		UnitMap[r] = u 
	}
}

// Check if a label describes a valid Unit 
func ValidUnit(label string) bool {
  if _, ok := UnitMap[label]; ok {
		return true
	} else {
		return false
  }
}

// Stringer for Measure
func (m Measure) String() string {
	return fmt.Sprintf("%.2f %s", m.Amount, m.Labels[0])
}

// Convert a Measure to a target Unit and return a new Measure
// of the converted amount and unit. The converted value is rounded
// to two decimal places.
func (m Measure) Convert(target string) (new Measure,ok bool) {
	if ValidUnit(target) {
		newUnit := UnitMap[target]
		if m.MeasureType != newUnit.MeasureType {
			ok = false
		} else {
			new.Amount = math.Floor((m.Amount * m.TypeEquivalent * 100 + .5) / newUnit.TypeEquivalent) / 100
			new.Unit = newUnit
			ok = true
		}
	} else {
		ok = false
	}
	return
}
