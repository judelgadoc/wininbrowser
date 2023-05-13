package handlers

import (
    "database/sql"
    "net/http"
    "log"
    "strconv"

    "github.com/satovarr/wininbrowser/clock/data"
    "github.com/gorilla/mux"
)


type Users struct {
    l *log.Logger
    db *sql.DB
}

func NewUsers(l *log.Logger, db *sql.DB) *Users {
    return &Users{l, db}
}

func (u *Users) AddUser(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	user_id, err := strconv.Atoi(vars["user_id"])
	if err != nil {
		http.Error(rw, "Unable to convert user id", http.StatusBadRequest)
		return
	}
	u.l.Println("Handle POST User")

	data.AddUser(user_id, u.db)
}

func (u Users) DeleteUser(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	user_id, err := strconv.Atoi(vars["user_id"])
	if err != nil {
		http.Error(rw, "Unable to convert id", http.StatusBadRequest)
		return
	}

	u.l.Println("Handle DELETE User", user_id)

	data.DeleteUser(user_id, u.db)
}

