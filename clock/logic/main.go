package main

import (
    "net/http"
    "log"
    "os"
    "time"

    "github.com/satovarr/wininbrowser/clock/handlers"
    "github.com/gorilla/mux"
)



func main() {
    
    l := log.New(os.Stdout, "clock", log.LstdFlags)

    ah := handlers.NewAlarms(l)

    sm := mux.NewRouter()

    getRouter := sm.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc("/", ah.GetAlarms)
    
    postRouter := sm.Methods(http.MethodPost).Subrouter()
	postRouter.HandleFunc("/", ah.AddAlarm)
    postRouter.Use(ah.MiddlewareValidateAlarm)

    putRouter := sm.Methods(http.MethodPut).Subrouter()
    putRouter.HandleFunc("/{id:[0-9]+}", ah.UpdateAlarm)
    putRouter.Use(ah.MiddlewareValidateAlarm)

    deleteRouter := sm.Methods(http.MethodDelete).Subrouter()
    deleteRouter.HandleFunc("/{id:[0-9]+}", ah.DeleteAlarm)

    s := http.Server{
        Addr:         "localhost:9090",  // configure the bind address
		Handler:      sm,                // set the default handler
		ErrorLog:     l,                 // set the logger for the server
		ReadTimeout:  5 * time.Second,   // max time to read request from the client
		WriteTimeout: 10 * time.Second,  // max time to write response to the client
		IdleTimeout:  120 * time.Second, // max time for connections using TCP Keep-Alive
    }

    s.ListenAndServe()
}

