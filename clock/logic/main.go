package main

import (
    "database/sql"
    "net/http"
    "log"
    "os"
    "time"

    "wininbrowser_clock_ms/handlers"
    "github.com/go-sql-driver/mysql"
    "github.com/gorilla/mux"
)



func main() {
    l := log.New(os.Stdout, "", log.LstdFlags)
    db, err := createConnection()
    if err != nil {
        log.Fatal("Can't create connection: ", err)
    }
    defer db.Close()

    ah := handlers.NewAlarms(l, db)
    th := handlers.NewTimers(l, db)
    tzh := handlers.NewTimezones(l, db)
    uh := handlers.NewUsers(l, db)

    sm := mux.NewRouter()


    getRouterAlarms := sm.Methods(http.MethodGet).Subrouter()
	getRouterAlarms.HandleFunc("/{user_id:[0-9]+}/alarms", ah.GetAlarms)

    postRouterAlarms := sm.Methods(http.MethodPost).Subrouter()
	postRouterAlarms.HandleFunc("/{user_id:[0-9]+}/alarms", ah.AddAlarm)
    postRouterAlarms.Use(ah.MiddlewareValidateAlarm)

    putRouterAlarms := sm.Methods(http.MethodPut).Subrouter()
    putRouterAlarms.HandleFunc("/alarms/{id:[0-9]+}", ah.UpdateAlarm)
    putRouterAlarms.Use(ah.MiddlewareValidateAlarm)

    deleteRouterAlarms := sm.Methods(http.MethodDelete).Subrouter()
    deleteRouterAlarms.HandleFunc("/alarms/{id:[0-9]+}", ah.DeleteAlarm)

    getRouterTimers := sm.Methods(http.MethodGet).Subrouter()
	getRouterTimers.HandleFunc("/{user_id:[0-9]+}/timers", th.GetTimers)

    postRouterTimers := sm.Methods(http.MethodPost).Subrouter()
	postRouterTimers.HandleFunc("/{user_id:[0-9]+}/timers", th.AddTimer)
    postRouterTimers.Use(th.MiddlewareValidateTimer)

    putRouterTimers := sm.Methods(http.MethodPut).Subrouter()
    putRouterTimers.HandleFunc("/timers/{id:[0-9]+}", th.UpdateTimer)
    putRouterTimers.Use(th.MiddlewareValidateTimer)

    deleteRouterTimers := sm.Methods(http.MethodDelete).Subrouter()
    deleteRouterTimers.HandleFunc("/timers/{id:[0-9]+}", th.DeleteTimer)

    getRouterTimezones := sm.Methods(http.MethodGet).Subrouter()
	getRouterTimezones.HandleFunc("/{user_id:[0-9]+}/timezones", tzh.GetTimezones)
	getRouterTimezones.HandleFunc("/timezones", tzh.GetAllTimezones)

    postRouterTimezones := sm.Methods(http.MethodPost).Subrouter()
	postRouterTimezones.HandleFunc("/{user_id:[0-9]+}/timezones", tzh.AddTimezone)
    postRouterTimezones.Use(tzh.MiddlewareValidateTimezone)

    putRouterTimezones := sm.Methods(http.MethodPut).Subrouter()
    putRouterTimezones.HandleFunc("/timezones/{id:[0-9]+}", tzh.UpdateTimezone)
    putRouterTimezones.Use(tzh.MiddlewareValidateTimezone)

    deleteRouterTimezones := sm.Methods(http.MethodDelete).Subrouter()
    deleteRouterTimezones.HandleFunc("/timezones/{id:[0-9]+}", tzh.DeleteTimezone)

    postRouterUsers := sm.Methods(http.MethodPost).Subrouter()
	postRouterUsers.HandleFunc("/users/{user_id:[0-9]+}", uh.AddUser)

    deleteRouterUsers := sm.Methods(http.MethodDelete).Subrouter()
    deleteRouterUsers.HandleFunc("/users/{user_id:[0-9]+}", uh.DeleteUser)

    s := http.Server{
        Addr:         ":9090",  // configure the bind address
		Handler:      sm,                // set the default handler
		ErrorLog:     l,                 // set the logger for the server
		ReadTimeout:  5 * time.Second,   // max time to read request from the client
		WriteTimeout: 10 * time.Second,  // max time to write response to the client
		IdleTimeout:  120 * time.Second, // max time for connections using TCP Keep-Alive
    }
    
    l.Println("Clock running on port 9090")
    s.ListenAndServe()
}

func createConnection() (*sql.DB, error) {
    log.Println(os.Getenv("DBUSER"), os.Getenv("DBPASS"), os.Getenv("DBHOST"))
    cfg := mysql.Config{
        User:   os.Getenv("DBUSER"),
        Passwd: os.Getenv("DBPASS"),
        Net:    "tcp",
        Addr:   os.Getenv("DBHOST") + ":55000",
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
