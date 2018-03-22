package main

import (
	"fmt"
	"github.com/thedodobird2/contentful-hugo/contentful"
)

const (
	// Content Delivery API
	cdaUrl   = ""
	cdaToken = ""
	// Content Preview API
	cpaBaseUrl = ""
	cpaToken   = ""
	spaceId = ""
)

func main() {
	fmt.Println("Contentful, meet Hugo")

	space := &contentful.Space{
		spaceId,
		cdaToken,
	}

	fmt.Println(string(space.GetEntries()))
}
