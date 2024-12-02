package main

import (
	"encoding/json"
	"fmt"
	"os"

	"music-request-api/models"
)

func GetConfigurations() models.Configuration {
	// Read the configuration file
	file, err := os.ReadFile("/app/configs/config.json")
	if err != nil {
		fmt.Printf("Read config file error: %v\n", err)
		os.Exit(1)
	}

	// Parse the JSON into the Configuration struct
	var config models.Configuration
	err = json.Unmarshal(file, &config)
	if err != nil {
		fmt.Printf("Problem in config file: %v\n", err.Error())
		os.Exit(1)
	}

	return config
}
