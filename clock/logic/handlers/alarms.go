package handlers

import (
    "context"
    "net/http"
    "log"
    "strconv"

    "github.com/satovarr/wininbrowser/clock/data"
    "github.com/gorilla/mux"
)


type Alarms struct {
    l *log.Logger
}

func NewAlarms(l *log.Logger) *Alarms {
    return &Alarms{l}
}

func (p *Alarms) GetAlarms(rw http.ResponseWriter, r *http.Request) {
    p.l.Println("Handle GET Alarms")

    la := data.GetAlarms()

    err := la.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
	}
}

func (a *Alarms) AddAlarm(rw http.ResponseWriter, r *http.Request) {
	a.l.Println("Handle POST Alarm")

	alarm := r.Context().Value(KeyAlarm{}).(data.Alarm)
	data.AddAlarm(&alarm)
}

func (a Alarms) UpdateAlarm(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(rw, "Unable to convert id", http.StatusBadRequest)
		return
	}

	a.l.Println("Handle PUT Alarm", id)
	alarm := r.Context().Value(KeyAlarm{}).(data.Alarm)

	data.UpdateAlarm(id, &alarm)
}

func (a Alarms) DeleteAlarm(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(rw, "Unable to convert id", http.StatusBadRequest)
		return
	}

	a.l.Println("Handle DELETE Alarm", id)

	data.DeleteAlarm(id)
}

type KeyAlarm struct{}

func (a Alarms) MiddlewareValidateAlarm(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		alarm := data.Alarm{}

		err := alarm.FromJSON(r.Body)
		if err != nil {
			a.l.Println("[ERROR] deserializing alarm", err)
			http.Error(rw, "Error reading alarm", http.StatusBadRequest)
			return
		}

		// add the alarm to the context
		ctx := context.WithValue(r.Context(), KeyAlarm{}, alarm)
		r = r.WithContext(ctx)

		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(rw, r)
	})
}
