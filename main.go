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
	// set configuration properties
	config.GetConfig(configFile)

	// setup Contentful Space
	space := &contentful.Space{
		ID:    config.Contentful.SpaceID,
		Token: config.Contentful.AccessToken,
	}

	fmt.Println(string(space.GetContentTypeEntries("testContentType")))
}
