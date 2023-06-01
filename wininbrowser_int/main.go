package main

import (
	"log"
	"net/http"

	"wininbrowser_int/handlers"
)

func main() {
	http.HandleFunc("/", handlers.SoapHandler)
	log.Println("Listening on port 55694")
	log.Fatal(http.ListenAndServe(":55694", nil))
}



