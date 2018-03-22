package config

import (
	"os"
	"log"
	"encoding/json"
)

type Config struct {
	Contentful struct {
		Space_id     string `json:"space_id"`
		Access_token string `json:"access_token"`
	} `json:"contentful"`
}

func Get(fileName string) Config {
	var config Config

	file, err := os.Open(fileName)
	defer file.Close()
	if err != nil {
		log.Fatal(err)
	}

	jsonParser := json.NewDecoder(file)
	jsonParser.Decode(&config)

	return config
}
