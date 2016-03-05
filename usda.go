package main

import (
	"strings"
	"fmt"
	"strconv"
)

func ParseUSDAFoodWeight(line string) (USDAFoodWeight,error) {
	var parsed USDAFoodWeight;

	var fields = strings.Split(line, "^")

	if len(fields) >= 5 {
		temp, _ := strconv.ParseInt(fields[2],10,32)
		parsed.Amount = int32(temp)
		parsed.FoodID = fields[0]
		parsed.Measure = fields[3]
		parsed.MassEq, _ = strconv.ParseFloat(fields[4],64)
		return parsed,nil
	} else {
		err := fmt.Errorf("Invalid Food Weight record cannot be parsed.")
		return parsed, err
	}
}

func IndexUSDAFoodWeight(record USDAFoodWeight) error {
	fmt.Printf("Indexing %d measure of %s for food %s is equivalent to %f grams.\n", record.Amount, record.Measure, record.FoodID, record.MassEq)
	return nil
}
