package main

import (
	"net/http"
	"encoding/json"
	"github.com/gorilla/mux"
	"fmt"
	"log"
	"bufio"
)

type ConversionRequest struct {
	Amount float64
	Unit string
	ConvertTo string
}

func startService(){
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/measure/convert", ConvertWebService)
	router.HandleFunc("/foodweight/index", IndexFoodWeightService)
	log.Fatal(http.ListenAndServe(":8088", router))
}

//ConvertWebService accepts a JSON body containing a float64 Amount field
//and a string Unit field describing the measure and a string ConvertTo field
//describing the requested conversion unit. 
func ConvertWebService(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("Content-Type") != "application/json" || r.Body == nil || r.Method != "POST" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w,"Invalid Reqest: this service requests a POST request in applicaiton/json format")
		return
	}

	var query ConversionRequest
	dec := json.NewDecoder(r.Body)

	if err := dec.Decode(&query); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w,"Invalid Request: %s", err)
		return
	}
	

	if !ValidUnit(query.Unit) || !ValidUnit(query.ConvertTo) {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w,"Invalid unit specified")
		return
	}

	var original Measure
	original.Amount = query.Amount
	original.Unit = UnitMap[query.Unit]

	var conversion, _ = original.Convert(query.ConvertTo)
	response, _ := json.Marshal(conversion)

	w.Write(response)
	return 
}

func IndexFoodWeightService(w http.ResponseWriter, r *http.Request){
	if r.Header.Get("Content-Type") != "text/plain" || r.Body == nil || r.Method != "POST" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w,"Invalid request: this service requires a POST request in text/plain format")
		return
	}

	var reqBody = bufio.NewScanner(r.Body)

	for reqBody.Scan() {
		foodWeight,err := ParseUSDAFoodWeight(reqBody.Text())
		if err != nil {
			log.Printf("Error: Could not parse food weight record - ", err)
		}		
	  err = IndexUSDAFoodWeight(foodWeight)
	}	

	
}
