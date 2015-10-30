package config


import (
	"os"
	"encoding/json"
	log "github.com/cihub/seelog"
)

type config struct {
	Port       string
	Connection string
}


func GetConfig() (currConfig config, err error) {

	//Open the config file. Defer to closing it when the function goes out of scope
	configFile, err := os.Open("config/env.json")
	defer configFile.Close()

	//If we had an error log it and return our error
	if err != nil {
		log.Error("Problem grabbing configuration", err)
		return
	}

	// Create a json parser for this file
	jsonParser := json.NewDecoder(configFile)

	//This is a good golang pattern to use:
	// grab the error from the Decode call and if the err is not null log
	if err = jsonParser.Decode(&currConfig); err != nil {
		log.Error("Problem parsing configuration", err)
		return
	}

	return currConfig, err

}