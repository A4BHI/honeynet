package main

import (
	"database/sql"
	"fmt"

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
		message TEXT,
		created_at TEXT);`)

	return err

}

func main() {
	e := Event{
		IP:   "192.168.1.50",
		Type: "SSH_LOGIN_FAIL",
	}

	alert := ThreatEngineBasic(e)

	fmt.Println("Attack Detected!!")
	fmt.Println(alert)
}
