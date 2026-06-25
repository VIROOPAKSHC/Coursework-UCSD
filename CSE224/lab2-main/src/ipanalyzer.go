package main

import (
	"fmt"
	"log"
	"math"
	"net"
	"os"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	if len(os.Args) != 2 && len(os.Args) != 3 {
		log.Fatalf("Usage: %s cidr_block [ip_address]", os.Args[0])
	}

	// os.Args[1] contains the cidr_block
	// os.Args[2] optionally contains the IP address to test
	cidr := os.Args[1]

	_, ipnet, err := net.ParseCIDR(cidr)
	if err != nil {
		panic("Error while parsing the provided CIDR notation")
	}

	mask := []byte(ipnet.Mask)

	broadcast := net.IP(make([]byte, 4))
	for i := 0; i < 4; i++ {
		broadcast[i] = ipnet.IP[i] | ^mask[i]
	}
	if len(os.Args) <= 2 {
		ones, bits := ipnet.Mask.Size()
		fmt.Println("Analyzing network :", cidr)
		fmt.Println()
		fmt.Println("Network address:", ipnet.IP)
		fmt.Println("Broadcast Address:", broadcast.String())
		fmt.Println("Subnet mask:", net.IP.String(mask))
		fmt.Println("Number of usable hosts:", math.Pow(2, float64(bits-ones))-2)
	} else {
		testIP := os.Args[2]
		ip2 := net.ParseIP(testIP)
		fmt.Println(ipnet.Contains(ip2))
	}
}
