package contentful

import (
	"io/ioutil"
	"log"
	"net/http"
)

func contentfulRequest(url string, token string) []byte {
	// initialize reusable client
	client := &http.Client{}

	// initialize request
	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}
	// add token as HTTP header
	request.Header.Add("Authorization", "Bearer "+token)

	// get the response for the just-configured request
	response, err := client.Do(request)
	if err != nil {
		log.Fatal(err)
	}

	// close body once function returns
	defer response.Body.Close()
	// get the body from the response (json)
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	return body
}
