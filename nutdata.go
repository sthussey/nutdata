// The API server handles incoming requests for nutrition data
package main

import (
	"fmt"
)

func main(){
	for _, u := range UnitReference {
		RegisterUnit(u)
	}
	m1 := Measure { amount:1 }
	m1.Unit = UnitMap["cup"]
	if m2, ok := m1.Convert("tsp"); ok {
		fmt.Println(m1,"=",m2)
	}
  if m3, ok := m1.Convert("oz"); ok {
		fmt.Println(m1,"=",m3)
	}
}
