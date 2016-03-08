package main

import (
	"fmt"
)

func initializeDatastore() bool {
	var mds = getMongoDatasource();
	fmt.Println("Should be connecting to MongoDB at ", mds)
	return true
}
