package config

import (
	"testing"
)

func TestGet(t *testing.T) {
	mockConfig := Config{}
	mockConfig.Contentful.Access_token = "mockTasdoken"
	mockConfig.Contentful.Space_id = "mockSpaceId"

	if Get("testcontent/mockConfig.json") != mockConfig {
		t.Error("File data didn't match mock:", mockConfig)
	}
}
