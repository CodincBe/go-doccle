package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Configuration struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func ReadConfiguration() Configuration {
	configuration := Configuration{}
	file, _ := os.Open("config.json")
	decoder := json.NewDecoder(file)
	err := decoder.Decode(&configuration)
	if err != nil {
		fmt.Println("error:", err)
	}
	return configuration
}
