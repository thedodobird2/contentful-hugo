package contentful

const (
	spaceURL = "https://cdn.contentful.com/spaces/"
)

// Space holds the needed properties for all requests to a Contentful Space
type Space struct {
	ID, Token string
}

// Entries defines the structure for the response from Contentful when querying for Entries
type Entries struct {
	Sys struct {
		Type string `json:"type"`
	} `json:"sys"`
	Total int     `json:"total"`
	Skip  int     `json:"skip"`
	Limit int     `json:"limit"`
	Items []Entry `json:"items"`
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

// GetContentTypeEntries triggers a request to the Contentful Delivery API
// which returns all Entries of the given Contentful Content Type
func (s *Space) GetContentTypeEntries(contentType string) []byte {
	return request(spaceURL+s.ID+"/entries?content_type="+contentType, s.Token)
}
