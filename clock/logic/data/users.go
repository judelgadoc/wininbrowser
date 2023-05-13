package data

import (
    "database/sql"
    "encoding/json"
    "io"
    "log"
)

type User struct {
    ID     int     `json:"id"`
}

type Users []*User

func (u *Users) ToJSON (w io.Writer) error {
    e := json.NewEncoder(w)
    return e.Encode(u)
}

func (u *User) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(u)
}

func AddUser(user_id int, db *sql.DB) {
    qry := `INSERT INTO User(id) VALUES (?)`
    res, err := db.Exec(qry, user_id)
    if err != nil {
        log.Fatal("Can't add user: ", err)
    }
    count, err := res.RowsAffected()
    if err != nil {
        log.Fatal(err)
    }
    log.Println("Added user: ", count)
}

func DeleteUser(user_id int, db *sql.DB) {
    qry := `DELETE FROM User WHERE id=?`
    res, err := db.Exec(qry, user_id)
    if err != nil {
        log.Fatal("Can't delete user: ", err)
    }
    count, err := res.RowsAffected()
    if err != nil {
        log.Fatal(err)
    }
    log.Println("Deleted user: ", count)
}
