// The API server handles incoming requests for nutrition data
package main

func main(){
	for _, u := range UnitReference {
		RegisterUnit(u)
	}
	startService()
}
