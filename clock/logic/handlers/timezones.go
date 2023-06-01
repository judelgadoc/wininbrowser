package handlers

import (
    "database/sql"
    "context"
    "net/http"
    "log"
    "strconv"

	"wininbrowser_clock_ms/data"
    "github.com/gorilla/mux"
)


type Timezones struct {
    l *log.Logger
    db *sql.DB
}

func NewTimezones(l *log.Logger, db *sql.DB) *Timezones {
    return &Timezones{l, db}
}

func (t *Timezones) GetTimezones(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	user_id, err := strconv.Atoi(vars["user_id"])
	if err != nil {
		http.Error(rw, "Unable to convert user id", http.StatusBadRequest)
		return
	}

    t.l.Println("Handle GET Timezones")

    lt := data.GetTimezones(user_id, t.db)

    err = lt.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
	}
}

func (t *Timezones) GetAllTimezones(rw http.ResponseWriter, r *http.Request) {
    t.l.Println("Handle GET All Timezones")

    lt := data.GetAllTimezones()
    err := lt.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
	}
}

func (t *Timezones) AddTimezone(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	user_id, err := strconv.Atoi(vars["user_id"])
	if err != nil {
		http.Error(rw, "Unable to convert user id", http.StatusBadRequest)
		return
	}
	t.l.Println("Handle POST Timezone")

	timezone := r.Context().Value(KeyTimezone{}).(data.Timezone)
	data.AddTimezone(user_id, &timezone, t.db)
}

func (t Timezones) UpdateTimezone(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(rw, "Unable to convert id", http.StatusBadRequest)
		return
	}

	t.l.Println("Handle PUT Timezone", id)
	timezone := r.Context().Value(KeyTimezone{}).(data.Timezone)

	data.UpdateTimezone(id, &timezone, t.db)
}

func (t Timezones) DeleteTimezone(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(rw, "Unable to convert id", http.StatusBadRequest)
		return
	}

	t.l.Println("Handle DELETE Timezone", id)

	data.DeleteTimezone(id, t.db)
}

type KeyTimezone struct{}

func (t Timezones) MiddlewareValidateTimezone(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		timezone := data.Timezone{}

		err := timezone.FromJSON(r.Body)
		if err != nil {
			t.l.Println("[ERROR] deserializing timezone", err)
			http.Error(rw, "Error reading timezone", http.StatusBadRequest)
			return
		}

		// add the Timezone to the context
		ctx := context.WithValue(r.Context(), KeyTimezone{}, timezone)
		r = r.WithContext(ctx)

		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(rw, r)
	})
}
