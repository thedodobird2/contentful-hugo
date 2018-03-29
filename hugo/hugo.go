package hugo

import (
	"bufio"
	"encoding/json"
	"github.com/thedodobird2/contentful-hugo/config"
	"github.com/thedodobird2/contentful-hugo/contentful"
	"log"
	"os"
)

// GenerateMarkdown takes your Contentful Space and Content Type
// and generates a markdown file for each Entry of that Content Type
func GenerateMarkdown(space *contentful.Space, contentType string) {
	path := config.Hugo.Root + "/content/" + contentType
	entries := getEntries(space, contentType)

	createFiles(entries, path)
}

func getEntries(space *contentful.Space, contentType string) []contentful.Entry {
	var entries contentful.Entries

	err := json.Unmarshal(space.GetContentTypeEntries(contentType), &entries)
	if err != nil {
		log.Fatal(err)
	}

	return entries.Items
}

func createFiles(entries []contentful.Entry, path string) {
	for _, item := range entries {
		frontMatter := frontMatter{Contentful: item}

		entryJSON, err := json.MarshalIndent(frontMatter, "", "  ")
		if err != nil {
			log.Fatal(err)
		}

		if err := os.MkdirAll(path, 0666); err != nil {
			log.Fatal(err)
		}

		file, err := os.Create(path + "/" + frontMatter.Contentful.Sys.ID + ".md")
		if err != nil {
			log.Fatal(err)
		}

		writer := bufio.NewWriter(file)
		writer.WriteString(string(entryJSON) + "\n")
		writer.Flush()
	}
}
