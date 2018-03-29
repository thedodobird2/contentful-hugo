package main

import (
	"github.com/thedodobird2/contentful-hugo/config"
	"github.com/thedodobird2/contentful-hugo/contentful"
	"github.com/thedodobird2/contentful-hugo/hugo"
)

const (
	configFile = "config.json"
)

func main() {
	// set configuration properties
	config.GetConfig(configFile)

	// set up Contentful Space
	space := &contentful.Space{
		ID:    config.Contentful.SpaceID,
		Token: config.Contentful.AccessToken,
	}

	hugo.GenerateMarkdown(space, "testContentType")
}
