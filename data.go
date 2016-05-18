package main

import (
	"fmt"
	"gopkg.in/mgo.v2"
)

var dbSession *mgo.Session

func initializeDatastore() bool {
	var mds = getMongoDatasource()

  var err error
	dbSession, err = mgo.Dial(mds)

  if err != nil {
		fmt.Printf("Could not initialize persistence store, exitting: %s\n", err)
		return false
	}
		
	return true
}

func PersistUSDAFoodWeight(record USDAFoodWeight) error {
	err := dbSession.DB("nutdata").C("foodWeight").Insert(record)

	if err != nil {
		return fmt.Errorf("Error persisting Food Weight %s: %s", record.FoodID, err)
	}

  return nil
}

func PersistUSDAFoodDesc(record USDAFoodDesc) error {
	err := dbSession.DB("nutdata").C("foodDesc").Insert(record)

  if err != nil {
		return fmt.Errorf("Error persisting Food Desc: %s", record.FoodID, err)
	}

	return nil
}

func IndexUSDAFoodWeight(record USDAFoodWeight) error {
  return fmt.Errorf("Error indexing Food Weight %s", record.FoodID)
}

func IndexUSDAFoodDesc(record USDAFoodDesc) error {
	return fmt.Errorf("Error indexing Food Desc %s", record.FoodID)
}
