package main

import (
	"flag"
	"log"
)

func main() {
	port := flag.Int("port", 3333, "Port to accept connections on")
	host := flag.String("host", "127.0.0.1", "Host to bind to")
	flag.Parse()

	log.Printf("%s:%d", *host, *port)
}
