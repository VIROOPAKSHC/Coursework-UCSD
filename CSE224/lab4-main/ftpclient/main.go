package main

import (
	"flag"
	"log"
	"os"
)

func main() {
	port := flag.Int("port", 3333, "Port to accept connections on")
	host := flag.String("host", "127.0.0.1", "Host to bind to")
	flag.Parse()

	if len(os.Args) < 2 {
		log.Fatalf("Usage: %s <STOR|RETR> <filename> [rate]", os.Args[0])
	}

	log.Printf("Connecting to %s on port %d", *host, *port)
}
