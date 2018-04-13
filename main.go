package main

import (
	"flag"
	"github.com/thedodobird2/contentful-hugo/config"
	"github.com/thedodobird2/contentful-hugo/contentful"
	"github.com/thedodobird2/contentful-hugo/hugo"
	"log"
)

const (
	configFile = "config.json"
)

var (
	contentType string
)

func init() {
	const (
		defaultValue = ""
		usage        = "the Content Type from Contentful you wish to add to your Hugo website"
	)
	flag.StringVar(&contentType, "contentType", defaultValue, usage)
	flag.StringVar(&contentType, "c", defaultValue, usage+" (shorthand)")
}

func main() {
	flag.Parse()

	// set configuration properties
	config.GetConfig(configFile)

	// set up Contentful Space
	space := &contentful.Space{
		ID:    config.Contentful.SpaceID,
		Token: config.Contentful.AccessToken,
	}

	contentType = "exampleContentType"
	if contentType == "" {
		log.Fatal("CLI requestError: Specify the Content Type from Contentful you wish to add to Hugo with the -c flag.")
	}

	hugo.GenerateMarkdown(space, contentType)
}
