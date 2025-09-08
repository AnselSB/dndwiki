package rest

import (
	"fmt"
	"io"
	"net/http"
	"strings"
	"wikiBot/src/cache"
)

const BotColor = 0xa026c4

func makeRequest(endpointURL string) ([]byte, error) {
	// before we try to hit the endpoint let's take from the cache first
	data, err := cache.GetObject(endpointURL)
	// if no errors we just return the cached object instead
	if err == nil {
		return data, nil
	}
	// first we need to define the method and url we are trying to hit
	url := fmt.Sprintf("https://www.dnd5eapi.co/api/2014/%v", endpointURL)
	method := "GET"

	// now make the client and the request object with proper headers
	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println("There was an error in creating the client")
		return nil, err
	}
	req.Header.Add("Accept", "application/json")

	// now make the actual request to the third party api
	res, err := client.Do(req)

	if err != nil {
		fmt.Println("There was an error in making the request")
		return nil, err
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Error during json parsing")
		return nil, err
	}
	// now when we have this body we write to the cache with the endpoint url
	writeErr := cache.SetObject(endpointURL, body)
	if writeErr != nil {
		fmt.Printf("An Error occurred while trying to write to the cache: %v\n", writeErr)
	}
	return body, nil

}

func formatMultiValues(values []string) string {
	var formattedString strings.Builder
	for _, value := range values {
		formattedString.WriteString(value)
		formattedString.WriteString("\n")
	}
	return formattedString.String()
}
