package serveSOAP

import (
    "encoding/json"
	"fmt"
	"log"
	"io/ioutil"
	"net/http"
	"net/url"
	"crypto/tls"
)

type Event struct {
	Title       string `xml:"title"`
	Description string `xml:"description"`
	Start       string `xml:"start"`
	End         string `xml:"end"`
	Location    string `xml:"location"`
	AllDay      int    `xml:"allDay"`
}


type GraphQLResponse struct {
	Data struct {
		EventsByUsername []Event `json:"eventsByUsername"`
	} `json:"data"`
}

func GetEvents(username string) []Event {
	endpoint := "https://host.docker.internal:4000/graphql"

	query := fmt.Sprintf(`
	{
		eventsByUsername(username: "%s") {
			title
			description
            start
            end
            location
            allDay
		}
	}
	`, username)

	params := url.Values{}
	params.Set("query", query)
	url := endpoint + "?" + params.Encode()

	transport := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	
	client := &http.Client{
		Transport: transport,
	}

	resp, err := client.Get(url)
	if err != nil {
		log.Fatal("Error sending HTTP request:", err)
		return nil
	}
	defer resp.Body.Close()

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Error reading response body:", err)
		return nil
	}

	var graphqlResponse GraphQLResponse
	err = json.Unmarshal(responseBody, &graphqlResponse)
	if err != nil {
		log.Fatal("Error parsing GraphQL response:", err)
		return nil
	}

	events := graphqlResponse.Data.EventsByUsername

    log.Println("Got events for", username)
	return events
}
