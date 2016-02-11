package main

import (
	"fmt"
)

// A Unit is a measure or weight unit including a slice of
// equivalent labels (the first member is the default label),
// the type of measure ('volume' or 'mass')
// and the type equivalency for the unit. Type equivalency
// is the amount of reference units in this unit. Reference
// units are milliliters for volume and grams for mass.
type Unit struct {
	labels []string
	measureType string
	typeEquivalent float64
}

// A Measure is a measurement amount with an assigned
// unit.
type Measure struct {
  Unit
  amount float64
}

var UnitReference = []Unit{
	{labels: []string{"tsp","teaspoon"}, measureType: "volume", typeEquivalent: 4.92892},
	{labels: []string{"tbsp","tablespoon"}, measureType: "volume", typeEquivalent: 14.7868},
	{labels: []string{"fl oz","fluid ounce"}, measureType: "volume", typeEquivalent: 29.5735},
	{labels: []string{"cup","c"}, measureType: "volume", typeEquivalent: 236.5882365},
	{labels: []string{"pint","pt","p"}, measureType: "volume", typeEquivalent: 473.176},
	{labels: []string{"l","litre"}, measureType: "volume", typeEquivalent: 1000},
	{labels: []string{"ml","mils","millilitre"}, measureType: "volume", typeEquivalent: 1},
	{labels: []string{"oz","ounce"}, measureType: "mass", typeEquivalent: 28.3495},
	{labels: []string{"lb","pound"}, measureType: "mass", typeEquivalent: 453.592},
	{labels: []string{"stone"}, measureType: "mass", typeEquivalent: 6350.29},
	{labels: []string{"kg","kilogram"}, measureType: "mass", typeEquivalent: 1000},
	{labels: []string{"g","gram"}, measureType: "mass", typeEquivalent: 1},
}

var UnitMap = make(map[string]Unit)

// RegisterUnit adds a Unit struct to the Map of all valid units
func RegisterUnit(u Unit){
  
	for _, r := range u.labels {
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
	return fmt.Sprintf("%.2f %s", m.amount, m.labels[0])
}

// Convert a Measure to a target Unit and return a new Measure
// of the converted amount and unit
func (m Measure) Convert(target string) (new Measure,ok bool) {
	if ValidUnit(target) {
		newUnit := UnitMap[target]
		if m.measureType != newUnit.measureType {
			ok = false
		} else {
			new.amount = (m.amount * m.typeEquivalent) / newUnit.typeEquivalent
			new.Unit = newUnit
			ok = true
		}
	} else {
		ok = false
	}
	return
}
  
