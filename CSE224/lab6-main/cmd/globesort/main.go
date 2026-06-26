package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"gopkg.in/yaml.v3"
)

type GlobeSortConfig struct {
	Nodes []struct {
		NodeID string
		Host   string
		Port   int
	}
}

func readConfig(configFilePath string) (*GlobeSortConfig, error) {
	config := GlobeSortConfig{}
	content, err := os.ReadFile(configFilePath)
	if err != nil {
		return nil, err
	}

	err = yaml.Unmarshal(content, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}

func main() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.SetOutput(os.Stdout)

	if len(os.Args) != 5 {
		fmt.Println("Usage:", os.Args[0], "<nodeID> <inputFilePath> <outputFilePath> <configFilePath>")
		os.Exit(1)
	}

	serverId, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatalf("Server ID must be an integer! Got '%s'", os.Args[1])
	}

	inputFilePath := os.Args[2]
	outputFilePath := os.Args[3]
	configFilePath := os.Args[4]

	log.Printf("serverID: %d", serverId)
	log.Printf("inputFilePath: %s", inputFilePath)
	log.Printf("outputFilePath: %s", outputFilePath)
	log.Printf("configFilePath: %s", configFilePath)

	config, err := readConfig(configFilePath)
	if err != nil {
		log.Fatalf("Failed to read config file: %v", err)
	}

	log.Printf("Configured nodes: %+v", config.Nodes)
}
