package main

var Config = make(map[string]string)

func initializeConfig() error {
	Config["mongoDS"] = "mongodb://data:27017/nutdata"

	return nil
}
func getMongoDatasource() string {
  return Config["mongoDS"]
}
