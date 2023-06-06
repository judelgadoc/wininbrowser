package main

import (
	"log"
	"net/http"

	"wininbrowser_int/serveSOAP"
	"wininbrowser_int/consumeSOAP"
)



func main() {

	soapMux := http.NewServeMux()
	soapMux.HandleFunc("/", serveSOAP.SoapHandler)

	restMux := http.NewServeMux()
	restMux.HandleFunc("/", consumeSOAP.RestHandler)

	// Start the servers
	go func() {
		log.Println("SOAP server listening on port 55694...")
		err := http.ListenAndServe(":55694", soapMux)
		if err != nil {
			log.Fatal("SOAP server (55694) failed: ", err)
		}
	}()

	go func() {
		log.Println("REST server listening on port 29162...")
		err := http.ListenAndServe(":29162", restMux)
		if err != nil {
			log.Fatal("REST server (29162) failed: ", err)
		}
	}() 

	// Wait indefinitely
	select {}
}