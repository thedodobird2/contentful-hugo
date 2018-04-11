package contentful

import (
	"io/ioutil"
	"log"
	"net/http"
	"github.com/thedodobird2/contentful-hugo/config"
)

func request(url string) []byte {
	client := &http.Client{}

	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}

	request.Header.Add("Authorization", "Bearer "+config.Contentful.AccessToken)

	response, err := client.Do(request)
	if err != nil {
		log.Fatal(err)
	}

	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	return body
}
