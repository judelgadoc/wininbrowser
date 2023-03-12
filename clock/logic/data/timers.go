package data

import (
    "database/sql"
    "encoding/json"
    "io"
    "log"
)

type Timer struct {
    ID     int     `json:"id"`
    Time   int     `json:"time"`
    UserID int     `json:"userId`
}

type Timers []*Timer

func (t *Timers) ToJSON (w io.Writer) error {
    e := json.NewEncoder(w)
    return e.Encode(t)
}

func (t *Timer) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(t)
}

func GetTimers(user_id int, db *sql.DB) Timers {
    var timers Timers
    qry := `SELECT * FROM Timer WHERE user_id=?`
    rows, err := db.Query(qry, user_id)
    if err != nil {
        log.Fatal("Can't query DB: ", err)
    }
    defer rows.Close()
    for rows.Next() {
        var (
            id int
            time int
            userid int
        )
        err = rows.Scan(&id, &time, &userid)
        if err != nil {
            log.Fatal("Scan: %v", err)
        }
        t := Timer{id, time, userid}
        timers = append(timers, &t)

    }
    return timers
}

func AddTimer(user_id int, t *Timer, db *sql.DB) {
    qry := `INSERT INTO Timer(time, user_id) VALUES (?, ?)`
    res, err := db.Exec(qry, t.Time, user_id)
    if err != nil {
        log.Fatal("Can't add timer: ", err)
    }
    count, err := res.RowsAffected()
    if err != nil {
        log.Fatal(err)
    }
    log.Println("Added timer: ", count)
}

func UpdateTimer(id int, t *Timer, db *sql.DB) {
    qry := `UPDATE Timer SET time=? WHERE id=?`
    res, err := db.Exec(qry, t.Time, id)
    if err != nil {
        log.Fatal("Can't update timer: ", err)
    }
    count, err := res.RowsAffected()
    if err != nil {
        log.Fatal(err)
    }
    log.Println("Updated timer: ", count)
}

func DeleteTimer(id int, db *sql.DB) {
    qry := `DELETE FROM Timer WHERE id=?`
    res, err := db.Exec(qry, id)
    if err != nil {
        log.Fatal("Can't delete timer: ", err)
    }
    count, err := res.RowsAffected()
    if err != nil {
        log.Fatal(err)
    }
    log.Println("Deleted timer: ", count)
}
