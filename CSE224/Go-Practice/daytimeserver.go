package main

import (
	"log"
	"net"
	"time"
)

func daytimeservermain() {
	service := ":13000"

	listener, err := net.Listen("tcp", service)
	if err != nil {
		log.Fatalf("Fatal error: %s\n", err.Error())
	}

	for {
		conn, err := listener.Accept()

		if err != nil {
			continue
		}

		now := time.Now()
		daytime := now.Format("Monday, January 2, 2006 15:04:05-MST")

		_, err = conn.Write([]byte(daytime))

		if err != nil {
			log.Printf("write error: %v", err)
			_ = conn.Close()
			continue
		}
		_ = conn.Close()
	}
}
