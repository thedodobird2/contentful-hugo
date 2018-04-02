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
	entries, err := contentful.GetEntries(space, contentType)
	if err != nil {
		log.Fatal(err)
	}

	for _, item := range entries {
		frontMatter := frontMatter{Contentful: item, Type: contentType}
		createFile(frontMatter, path)
	}

}

func createFile(frontMatter frontMatter, path string) {
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
