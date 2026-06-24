package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	if len(os.Args) != 2 && len(os.Args) != 3 {
		log.Fatalf("Usage: %s cidr_block [ip_address]", os.Args[0])
	}

	// os.Args[1] contains the cidr_block
	// os.Args[2] optionally contains the IP address to test

	// Replace the line below and start coding your logic from here.
	fmt.Println("Hello, World!")
}
