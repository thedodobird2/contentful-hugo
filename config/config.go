package config

import (
	"encoding/json"
	"log"
	"os"
)

var (
	// Contentful contains configuration properties for connecting Contentful
	Contentful contentful
	// Hugo contains configuration properties for the Hugo website
	Hugo hugo
)

type config struct {
	Contentful contentful `json:"contentful"`
	Hugo       hugo       `json:"hugo"`
}

type contentful struct {
	SpaceID     string `json:"spaceId"`
	AccessToken string `json:"accessToken"`
}

type hugo struct {
	Root string `json:"root"`
}

// GetConfig gets configuration properties from a json file
// and puts the values into the global properties
func GetConfig(file string) {
	var config config

	configJSON, err := os.Open(file)
	defer configJSON.Close()
	if err != nil {
		log.Fatal(err)
	}

	jsonParser := json.NewDecoder(configJSON)
	jsonParser.Decode(&config)

	Contentful = config.Contentful
	Hugo = config.Hugo
}
