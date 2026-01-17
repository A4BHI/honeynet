package portscan

import (
	"fmt"
	"log"
	"net"
)

func AcceptConnRequest() {
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}
	conn, err := ln.Accept()
	if err != nil {
		log.Fatal(err)
	}

	adr := conn.RemoteAddr()
	fmt.Print(adr)
}
