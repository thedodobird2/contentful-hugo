package hugo

import "github.com/thedodobird2/contentful-hugo/contentful"

// frontMatter defines the structure of the Front Matter
// which is used for all generated content files for Hugo
type frontMatter struct {
	Contentful  contentful.Entry `json:"contentful"`
	Type        string           `json:"type"`
	Url         string           `json:"url"`
	Draft       bool             `json:"draft"`
	PublishDate string           `json:"publishdate"`
	ExpiryDate  string           `json:"expirydate"`
}
