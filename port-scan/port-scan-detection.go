package portscan

import (
	"fmt"
	"log"
	"net"
)

func AcceptConnRequest() {
	ln, err := net.Listen("tcp", ":2222")
	if err != nil {
		log.Fatal(err)
	}
	conn, err := ln.Accept()

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connection Accepted!!")

	adr := conn.RemoteAddr()
	ip, port, err := net.SplitHostPort(adr.String())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("IP:", ip)
	fmt.Println("PORT:", port)

}
