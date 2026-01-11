package main

import (
	"database/sql"
	"fmt"
	"time"

	_ "modernc.org/sqlite"
)

type Event struct {
	IP   string
	Type string
}

type Alert struct {
	IP       string
	Severity string
	Type     string
}

var db *sql.DB

func ThreatEngineBasic(e Event) Alert {
	if e.Type == "SSH_LOGIN_FAIL" {
		alert := Alert{
			IP:       e.IP,
			Severity: "HIGH",
			Type:     e.Type,
		}

		return alert
	}

	return Alert{
		IP:       e.IP,
		Severity: "LOW",
		Type:     "UNKNOWN",
	}
}

func CreateTable(db *sql.DB) error {
	_, err := db.Exec(`CREATE TABLE IF NOT EXISTS ALERT(id INTEGER PRIMARY KEY AUTOINCREMENT,
		source_ip TEXT,
		severity TEXT,
		type TEXT,
		created_at TEXT);`)

	return err

}

func SaveToDB(db *sql.DB, a Alert) error {
	_, err := db.Exec("INSERT INTO ALERT (source_ip,severity,type,created_at) VALUES(?,?,?,?)", a.IP, a.Severity, a.Type, time.Now().Format(time.RFC3339))
	return err

}

func ReadFromDB(db *sql.DB) error {
	rows, err := db.Query(
		`SELECT id, source_ip, severity, type, created_at FROM alert`,
	)
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var ip, sev, msg, created string

		rows.Scan(&id, &ip, &sev, &msg, &created)

		fmt.Println(id, ip, sev, msg, created)
	}
	return nil

}

func main() {
	var err error
	db, err = sql.Open("sqlite", "alerts.db")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer db.Close()

	if err := CreateTable(db); err != nil {
		panic(err)
	}

	e := Event{
		IP:   "192.168.1.50",
		Type: "SSH_LOGIN_FAIL",
	}

	alert := ThreatEngineBasic(e)

	if err := SaveToDB(db, alert); err != nil {
		panic(err)
	}

	fmt.Println("Attack Detected!!")
	ReadFromDB(db)
}
