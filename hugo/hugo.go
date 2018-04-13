package hugo

import (
	"encoding/json"
	"github.com/thedodobird2/contentful-hugo/config"
	"github.com/thedodobird2/contentful-hugo/contentful"
	"log"
	"os"
	"strings"
	"bufio"
	"strconv"
	"github.com/gohugoio/hugo/parser"
	"fmt"
)

// GenerateMarkdown takes your Contentful Space and Content Type
// and generates a markdown file for each Entry of that Content Type
func GenerateMarkdown(space *contentful.Space, contentTypeID string) {
	entries, err := contentful.GetEntries(space, contentTypeID)
	if err != nil {
		log.Fatal(err)
	}

	for _, item := range entries {
		filename, dir := generateHugoCompliantPath(item)
		frontMatter := frontMatter{Contentful: item, Type: item.Sys.ContentType.Sys.ID}
		generateFile(filename, dir, frontMatter)
	}
}

func generateHugoCompliantPath(entry contentful.Entry) (string, string) {
	dir := config.Hugo.Root + "/content/" + entry.Sys.ContentType.Sys.ID
	filename := entry.GetEntryTitle()

	if strings.Contains(filename, " ") {
		filename = strings.Replace(filename, " ", "-", -1)
	}

	return strings.ToLower(filename), strings.ToLower(dir)
}

func generateFile(filename string, dir string, fm frontMatter) {
	entryJSON, err := json.MarshalIndent(fm, "", "  ")
	if err != nil {
		log.Fatal(err)
	}

	if err := os.MkdirAll(dir, 0666); err != nil {
		log.Fatal(err)
	}

	file, err := os.Create(generateFileName(fm.Contentful.Sys.ID, filename, dir))
	if err != nil {
		log.Fatal(err)
	}

	writer := bufio.NewWriter(file)
	writer.WriteString(string(entryJSON) + "\n")
	writer.Flush()
}

func generateFileName(id string, filename string, dir string) string {
	path := dir + "/" + filename + ".md"

	i := 1
	for {
		if _, err := os.Stat(path); os.IsNotExist(err) {
			break
		} else {
			file, err := os.Open(path)
			if err != nil {
				log.Fatal(err)
			}

			reader := bufio.NewReader(file)
			page, err := parser.ReadFrom(reader)
			if err != nil {
				log.Fatal(err)
			}
			var fileFrontMatter frontMatter

			fmt.Println(page.FrontMatter())

			if err := json.Unmarshal(page.FrontMatter(), fileFrontMatter); err != nil {
				log.Fatal(err)
			}

			if id != fileFrontMatter.Contentful.Sys.ID {
				path = dir + "/" + filename + "-" + strconv.Itoa(i) + ".md"
				i++
			}
		}
	}

	return path
}
