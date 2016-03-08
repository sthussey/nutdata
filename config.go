package main

import (
	"github.com/spf13/viper"
	"fmt"
)

func initializeConfig() {
	viper.SetConfigName("nutdata-config")
	viper.AddConfigPath("/etc/nutdata/")
	viper.AddConfigPath(".")
	viper.SetConfigType("yaml")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error reading config file: %s \n", err))
	}
}

func getMongoDatasource() string {
  return viper.GetString("datasource.mongodb.url")
}

func getElasticsearchDatasource() string {
  return viper.GetString("datasource.elasticsearch.url")
}

