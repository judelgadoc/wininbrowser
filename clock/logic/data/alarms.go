package data

import (
    "database/sql"
    "encoding/json"
    "io"
    "log"
    "os"

    "github.com/go-sql-driver/mysql"
)

type Alarm struct {
    ID     int     `json:"id"`
    Title  string  `json:"title"`
    Time   string  `json:"time"`
    UserID int     `json:"userId`
}

type Alarms []*Alarm

func (a *Alarms) ToJSON (w io.Writer) error {
    e := json.NewEncoder(w)
    return e.Encode(a)
}

func (a *Alarm) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(a)
}

func GetAlarms() Alarms {
    db, err := createConnection()
    if err != nil {
        log.Fatal("Can't create connection: ", err)
    }
    defer db.Close()
    var alarms Alarms
    qry := `SELECT * FROM Alarm`
    rows, err := db.Query(qry)
    if err != nil {
        log.Fatal("Can't query DB: ", err)
    }
    defer rows.Close()
    for rows.Next() {
        var (
            id int
            title string
            time string
            userid int
        )
        err = rows.Scan(&id, &title, &time, &userid)
        if err != nil {
            log.Fatal("Scan: %v", err)
        }
        a := Alarm{id, title, time, userid}
        alarms = append(alarms, &a)

    }
    return alarms
}

func AddAlarm(a *Alarm) {
    db, err := createConnection()
    if err != nil {
        log.Fatal("Can't create connection: ", err)
    }
    defer db.Close()
    qry := `INSERT INTO Alarm(title, time, user_id) VALUES (?, ?, ?)`
    res, err := db.Exec(qry, a.Title, a.Time, a.UserID)
    if err != nil {
        log.Fatal("Can't add alarm: ", err)
    }
    count, err := res.RowsAffected()
    if err != nil {
        log.Fatal(err)
    }
    log.Println("Added alarm: ", count)
}

func UpdateAlarm(id int, a *Alarm) {
    db, err := createConnection()
    if err != nil {
        log.Fatal("Can't create connection: ", err)
    }
    defer db.Close()
    qry := `UPDATE Alarm SET title=?, time=? WHERE id=?`
    res, err := db.Exec(qry, a.Title, a.Time, id)
    if err != nil {
        log.Fatal("Can't update alarm: ", err)
    }
    count, err := res.RowsAffected()
    if err != nil {
        log.Fatal(err)
    }
    log.Println("Updated alarm: ", count)
}

func DeleteAlarm(id int) {
    db, err := createConnection()
    if err != nil {
        log.Fatal("Can't create connection: ", err)
    }
    defer db.Close()
    qry := `DELETE FROM Alarm WHERE id=?`
    res, err := db.Exec(qry, id)
    if err != nil {
        log.Fatal("Can't delete alarm: ", err)
    }
    count, err := res.RowsAffected()
    if err != nil {
        log.Fatal(err)
    }
    log.Println("Deleted alarm: ", count)
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

