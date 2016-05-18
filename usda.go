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
		parsed.FoodID = strings.Replace(fields[0],"~","",-1)
		parsed.Measure = strings.Replace(fields[3],"~","",-1)
		parsed.MassEq, _ = strconv.ParseFloat(fields[4],64)
		return parsed,nil
	} else {
		err := fmt.Errorf("Invalid Food Weight record cannot be parsed.")
		return parsed, err
	}
}

func ParseUSDAFoodDesc(line string) (USDAFoodDesc, error) {
	var parsed USDAFoodDesc

	var fields = strings.Split(line, "^")

	if len(fields) >= 4 {
		parsed.FoodID = strings.Replace(fields[0], "~", "", -1)
		parsed.LongDesc = strings.Replace(fields[2], "~", "", -1)
		parsed.ShortDesc = strings.Replace(fields[3], "~", "", -1)
		parsed.Aliases = strings.Replace(fields[4], "~", "", -1)
		parsed.Manufacturer = strings.Replace(fields[5], "~", "", -1)
		return parsed, nil
	} else {
		err := fmt.Errorf("Invalid Food Description record cannot be parsed.")
		return parsed, err
	}
}
