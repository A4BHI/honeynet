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
	ln.Accept()
	ip := ln.Addr().String()
	fmt.Println(ip)
}
