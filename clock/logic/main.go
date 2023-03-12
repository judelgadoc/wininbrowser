package main

import (
    "database/sql"
    "net/http"
    "log"
    "os"
    "time"

    "github.com/satovarr/wininbrowser/clock/handlers"
    "github.com/go-sql-driver/mysql"
    "github.com/gorilla/mux"
)



func main() {
    l := log.New(os.Stdout, "clock", log.LstdFlags)
    db, err := createConnection()
    if err != nil {
        log.Fatal("Can't create connection: ", err)
    }
    defer db.Close()

    ah := handlers.NewAlarms(l, db)

    sm := mux.NewRouter()

    getRouter := sm.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc("/alarms", ah.GetAlarms)
    
    postRouter := sm.Methods(http.MethodPost).Subrouter()
	postRouter.HandleFunc("/alarms", ah.AddAlarm)
    postRouter.Use(ah.MiddlewareValidateAlarm)

    putRouter := sm.Methods(http.MethodPut).Subrouter()
    putRouter.HandleFunc("/alarms/{id:[0-9]+}", ah.UpdateAlarm)
    putRouter.Use(ah.MiddlewareValidateAlarm)

    deleteRouter := sm.Methods(http.MethodDelete).Subrouter()
    deleteRouter.HandleFunc("/alarms/{id:[0-9]+}", ah.DeleteAlarm)

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

func createConnection() (*sql.DB, error) {
    cfg := mysql.Config{
        User:   os.Getenv("DBUSER"),
        Passwd: os.Getenv("DBPASS"),
        Net:    "tcp",
        Addr:   os.Getenv("DBHOST") + ":3306",
        DBName: "clock_db",
        AllowNativePasswords: true,
    }

    db, err := sql.Open("mysql", cfg.FormatDSN())
    if err != nil {
        log.Fatal("Can't open DB: ", err)
    }

    err = db.Ping()
    if err != nil {
        log.Fatal("Can't ping DB: ", err)
    }
    return db, err
}
