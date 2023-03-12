package handlers

import (
    "database/sql"
    "context"
    "net/http"
    "log"
    "strconv"

    "github.com/satovarr/wininbrowser/clock/data"
    "github.com/gorilla/mux"
)


type Timers struct {
    l *log.Logger
    db *sql.DB
}

func NewTimers(l *log.Logger, db *sql.DB) *Timers {
    return &Timers{l, db}
}

func (t *Timers) GetTimers(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	user_id, err := strconv.Atoi(vars["user_id"])
	if err != nil {
		http.Error(rw, "Unable to convert user id", http.StatusBadRequest)
		return
	}

    t.l.Println("Handle GET Timers")

    lt := data.GetTimers(user_id, t.db)

    err = lt.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
	}
}

func (t *Timers) AddTimer(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	user_id, err := strconv.Atoi(vars["user_id"])
	if err != nil {
		http.Error(rw, "Unable to convert user id", http.StatusBadRequest)
		return
	}
	t.l.Println("Handle POST Timer")

	timer := r.Context().Value(KeyTimer{}).(data.Timer)
	data.AddTimer(user_id, &timer, t.db)
}

func (t Timers) UpdateTimer(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(rw, "Unable to convert id", http.StatusBadRequest)
		return
	}

	t.l.Println("Handle PUT Timer", id)
	timer := r.Context().Value(KeyTimer{}).(data.Timer)

	data.UpdateTimer(id, &timer, t.db)
}

func (t Timers) DeleteTimer(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(rw, "Unable to convert id", http.StatusBadRequest)
		return
	}

	t.l.Println("Handle DELETE Timer", id)

	data.DeleteTimer(id, t.db)
}

type KeyTimer struct{}

func (t Timers) MiddlewareValidateTimer(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		timer := data.Timer{}

		err := timer.FromJSON(r.Body)
		if err != nil {
			t.l.Println("[ERROR] deserializing timer", err)
			http.Error(rw, "Error reading timer", http.StatusBadRequest)
			return
		}

		// add the timer to the context
		ctx := context.WithValue(r.Context(), KeyTimer{}, timer)
		r = r.WithContext(ctx)

		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(rw, r)
	})
}
