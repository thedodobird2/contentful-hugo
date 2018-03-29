package main

import (
	"fmt"
	"github.com/thedodobird2/contentful-hugo/config"
	"github.com/thedodobird2/contentful-hugo/contentful"
)

const (
	configFile = "config.json"
)

func main() {
	fmt.Println("Contentful, meet Hugo")

	config.GetConfig(configFile)

	space := &contentful.Space{
		Id:    config.Contentful.SpaceID,
		Token: config.Contentful.AccessToken,
	}

	fmt.Println(string(space.GetEntries()))
}
