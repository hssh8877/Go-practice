package main

import (
	"encoding/json"
	"flag"
	"log"
	"os"
)

func main() {

	inputFile := flag.String("file", "user_actions.txt", "Input file with user actions")
	outputFile := flag.String("output", "analysis.json", "Output JSON file")
	flag.Parse()

	stats := Analyze(*inputFile)

	data, err := json.MarshalIndent(stats, " ", "  ")
	if err != nil {
		log.Fatal(err)
	}

	err = os.WriteFile(*outputFile, data, 0644)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Analysis written to %s", *outputFile)
}
