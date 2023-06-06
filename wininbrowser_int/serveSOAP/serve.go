package serveSOAP

import (
	"encoding/xml"
	"net/http"
)


type GetEventsRequest struct {
	Username string `xml:"username"`
}


type GetEventsResponse struct {
	Events []Event `xml:"events>event"`
}

func SoapHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Read the SOAP request body
	decoder := xml.NewDecoder(r.Body)
	var requestEnvelope struct {
		Body struct {
			GetEventsRequest GetEventsRequest `xml:"getEventsRequest"`
		} `xml:"Body"`
	}
	err := decoder.Decode(&requestEnvelope)
	if err != nil {
		http.Error(w, "Failed to parse SOAP request", http.StatusBadRequest)
		return
	}

	// Process the SOAP request and prepare the response
	responseEnvelope := struct {
		XMLName xml.Name `xml:"Envelope"`
		Body    struct {
			GetEventsResponse GetEventsResponse `xml:"getEventsResponse"`
		} `xml:"Body"`
	}{
		XMLName: xml.Name{Local: "soap:Envelope"},
		Body: struct {
			GetEventsResponse GetEventsResponse `xml:"getEventsResponse"`
		}{
			GetEventsResponse: GetEventsResponse{Events: GetEvents(requestEnvelope.Body.GetEventsRequest.Username)},
		},
	}

	// Marshal the response into XML format
	responseXML, err := xml.MarshalIndent(responseEnvelope, "", "    ")
	if err != nil {
		http.Error(w, "Failed to marshal SOAP response", http.StatusInternalServerError)
		return
	}

	// Set the SOAP response headers
	w.Header().Set("Content-Type", "text/xml; charset=utf-8")
	w.WriteHeader(http.StatusOK)

	// Write the SOAP response
	_, err = w.Write([]byte(xml.Header + string(responseXML)))
	if err != nil {
		http.Error(w, "Failed to write SOAP response", http.StatusInternalServerError)
		return
	}
}