package contentful

import (
	"encoding/json"
	"fmt"
	"log"
)

const (
	spaceURL = "https://cdn.contentful.com/spaces/"
)

// Space holds the needed properties for all requests to a Contentful Space
type Space struct {
	ID, Token string
}

// Entry defines the structure of a single Contentful Entry
type Entry struct {
	Sys struct {
		ID          string `json:"id"`
		Type        string `json:"type"`
		Version     int    `json:"version"`
		Space       link   `json:"space"`
		ContentType link   `json:"contentType"`
		CreatedAt   string `json:"createdAt"`
		UpdatedAt   string `json:"updatedAt"`
		Revision    int    `json:"revision"`
	} `json:"sys"`
	Fields map[string]interface{} `json:"fields"`
}

type link struct {
	Sys struct {
		Type     string `json:"type"`
		LinkType string `json:"linkType"`
		ID       string `json:"id"`
	} `json:"sys"`
}

type entries struct {
	Sys struct {
		Type string `json:"type"`
	} `json:"sys"`
	Total int     `json:"total"`
	Skip  int     `json:"skip"`
	Limit int     `json:"limit"`
	Items []Entry `json:"items"`
}

type requestError struct {
	Sys struct {
		ID   string `json:"id"`
		Type string `json:"type"`
	}
	Message   string `json:"message"`
	RequestID string `json:"requestId"`
}

// GetEntries triggers a request to the Contentful Delivery API
// which returns all entries of the given Contentful Content Type.
//
// If Contentful returns an error based on a faulty Query, the message is returned.
func GetEntries(space *Space, contentTypeID string) ([]Entry, error) {
	var response entries
	var requestError requestError
	var err error

	entries := request(spaceURL + space.ID + "/entries?content_type=" + contentTypeID)

	if err = json.Unmarshal(entries, &response); err != nil {
		log.Fatal(err)
	}

	if response.Sys.Type == "Error" {
		if err = json.Unmarshal(entries, &requestError); err != nil {
			log.Fatal(err)
		}
		err = fmt.Errorf("The Content Type you are looking for probably doesn't exist in your Contentful"+
			"Space.\n\tContentful requestError: %s:\n\t%s", requestError.Sys.ID, requestError.Message)
	}

	return response.Items, err
}
