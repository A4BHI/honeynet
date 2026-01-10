package main

import "fmt"

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

func main() {
	e := Event{
		IP:   "192.168.1.50",
		Type: "SSH_LOGIN_FAIL",
	}

	alert := ThreatEngineBasic(e)

	fmt.Println("Attack Detected!!")
	fmt.Println(alert)
}
