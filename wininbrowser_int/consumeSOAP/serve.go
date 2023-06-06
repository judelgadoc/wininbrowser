package consumeSOAP

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"io/ioutil"
	"net/http"
)

type Envelope struct {
	XMLName xml.Name `xml:"http://schemas.xmlsoap.org/soap/envelope/ Envelope"`
	Body    Body     `xml:"Body"`
}

type Body struct {
	SayHelloResponse SayHelloResponse `xml:"say_helloResponse"`
}

type SayHelloResponse struct {
	SayHelloResult SayHelloResult `xml:"say_helloResult"`
}

type SayHelloResult struct {
	ScheduledPayments []ScheduledPayment `xml:"ScheduledPayment"`
}

type ScheduledPayment struct {
	UserId           int    `xml:"UserId"`
	Name             string `xml:"Name"`
	CategoryId       int    `xml:"CategoryId"`
	AccountId        int    `xml:"AccountId"`
	PaymentMethod    string `xml:"PaymentMethod"`
	Recipient        string `xml:"Recipient"`
	Frequency        string `xml:"Frequency"`
	StartDate        string `xml:"StartDate"`
	NotificationTime string `xml:"NotificationTime"`
}

func RestHandler(w http.ResponseWriter, r *http.Request) {
	// Read the request body
	xmlBody := `<soapenv:Envelope xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/" xmlns:spy="spyne.examples,hello.http">
	<soapenv:Header/>
	<soapenv:Body>
	   <spy:say_hello/>
	</soapenv:Body>
	</soapenv:Envelope>`

	// Make HTTP POST request
	resp, err := http.Post("https://21e8-186-98-91-20.ngrok-free.app", "application/xml", bytes.NewBufferString(xmlBody))
	if err != nil {
		http.Error(w, "HTTP POST request failed", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	// Read the response body
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, "Failed to read response body", http.StatusInternalServerError)
		return
	}

	// Unmarshal XML into struct
	var envelope Envelope
	err = xml.Unmarshal(respBody, &envelope)
	if err != nil {
		http.Error(w, "Failed to unmarshal XML", http.StatusInternalServerError)
		return
	}

	// Convert struct to JSON
	jsonData, err := json.Marshal(envelope.Body.SayHelloResponse.SayHelloResult.ScheduledPayments)
	if err != nil {
		http.Error(w, "Failed to marshal JSON", http.StatusInternalServerError)
		return
	}

	// Set the response headers
	w.Header().Set("Content-Type", "application/json")

	// Write the JSON response
	w.Write(jsonData)
}
