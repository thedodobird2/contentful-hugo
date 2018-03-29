package config

import (
	"testing"
)

const (
	mockConfigFile = "testcontent/mockConfig.json"
)

func TestGetConfig(t *testing.T) {
	GetConfig(mockConfigFile)

	var tests = []struct {
		mock     string
		expected string
	}{
		{"mockToken", Contentful.AccessToken},
		{"mockSpaceId", Contentful.SpaceID},
		{"mockPath", Hugo.Root},
	}

	for _, test := range tests {
		if test.mock != test.expected {
			t.Errorf("Test Failed: {%s} expected, received: {%s}", test.mock, test.expected)
		}
	}
}
