// The API server handles incoming requests for nutrition data
package main

import (
	"fmt"
)

func main(){
	for _, m := range MeasureReference {
		RegisterMeasure(m)
	}
	fmt.Print(MeasureMap)
}
