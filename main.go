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

func ThreatEngineBasic(e Event) {
	if e.Type == "SSH_LOGIN_FAIL" {
		alert := Alert{
			IP:       e.IP,
			Severity: "HIGH",
			Type:     e.Type,
		}
		fmt.Println("Attack Detected!!")
		fmt.Println(alert)
	}
}

func main() {}
