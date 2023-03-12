package data

import (
    "database/sql"
    "encoding/json"
    "io"
    "log"

    "zgo.at/tz"
)

type Timezone struct {
    ID     int     `json:"id,omitempty"`
    Name   string  `json:"name"`
    UserID int     `json:"userId,omitempty"`
}

type Timezones []*Timezone

func (t *Timezones) ToJSON (w io.Writer) error {
    e := json.NewEncoder(w)
    return e.Encode(t)
}

func (t *Timezone) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(t)
}

func GetTimezones(user_id int, db *sql.DB) Timezones {
    var timezones Timezones
    qry := `SELECT * FROM Timezone WHERE user_id=?`
    rows, err := db.Query(qry, user_id)
    if err != nil {
        log.Fatal("Can't query DB: ", err)
    }
    defer rows.Close()
    for rows.Next() {
        var (
            id int
            name string
            userid int
        )
        err = rows.Scan(&id, &name, &userid)
        if err != nil {
            log.Fatal("Scan: %v", err)
        }
        t := Timezone{id, name, userid}
        timezones = append(timezones, &t)
    }
    return timezones
}

func GetAllTimezones() Timezones {
    var timezones Timezones
    for _, z := range tz.Zones {
        var t Timezone
        t.Name = z.Zone
        timezones = append(timezones, &t)
    }
    return timezones
}

func AddTimezone(user_id int, t *Timezone, db *sql.DB) {
    qry := `INSERT INTO Timezone(name, user_id) VALUES (?, ?)`
    res, err := db.Exec(qry, t.Name, user_id)
    if err != nil {
        log.Fatal("Can't add timezone: ", err)
    }
    count, err := res.RowsAffected()
    if err != nil {
        log.Fatal(err)
    }
    log.Println("Added timezone: ", count)
}

func UpdateTimezone(id int, t *Timezone, db *sql.DB) {
    qry := `UPDATE Timezone SET name=? WHERE id=?`
    res, err := db.Exec(qry, t.Name, id)
    if err != nil {
        log.Fatal("Can't update timezone: ", err)
    }
    count, err := res.RowsAffected()
    if err != nil {
        log.Fatal(err)
    }
    log.Println("Updated timezone: ", count)
}

func DeleteTimezone(id int, db *sql.DB) {
    qry := `DELETE FROM Timezone WHERE id=?`
    res, err := db.Exec(qry, id)
    if err != nil {
        log.Fatal("Can't delete timezone: ", err)
    }
    count, err := res.RowsAffected()
    if err != nil {
        log.Fatal(err)
    }
    log.Println("Deleted timezone: ", count)
}
