package main

import (
	"fmt"
	"github.com/thedodobird2/contentful-hugo/config"
	"github.com/thedodobird2/contentful-hugo/contentful"
)

var (
	properties = config.Get("config.json")

	accessToken = properties.Contentful.Access_token
	spaceId     = properties.Contentful.Space_id
)

func main() {
	fmt.Println("Contentful, meet Hugo")

	space := &contentful.Space{
		spaceId,
		accessToken,
	}

	fmt.Println(string(space.GetEntries()))
}
