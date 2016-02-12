package main

import (
	"net/http"
	"encoding/json"
	"github.com/gorilla/mux"
	"fmt"
	"log"
)

type ConversionRequest struct {
	Amount float64
	Unit string
	ConvertTo string
}

func startService(){
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", ConvertWebService)
	log.Fatal(http.ListenAndServe(":8088", router))
}

//ConvertWebService accepts a JSON body containing a float64 Amount field
//and a string Unit field describing the measure and a string ConvertTo field
//describing the requested conversion unit. 
func ConvertWebService(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("Content-Type") != "application/json" || r.Body == nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w,"Invalid Reqest")
		return
	}

	var query ConversionRequest
	dec := json.NewDecoder(r.Body)

	if err := dec.Decode(&query); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w,"Invalid Request: %s", err)
		return
	}
	
	fmt.Printf("amount = %d, unit = %s, convertTo: %s\n", query.Amount, query.Unit, query.ConvertTo)

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


