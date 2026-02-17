package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"log"
	"os"
	"strings"
)

func main() {
	inputFile := flag.String("file", "tasks.txt", "Input tasks file")
	outputFile := flag.String("output", "tasks.json", "Output tasks file")
	flag.Parse()

	file, err := os.Open(*inputFile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	stats := TaskStats{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "TODO") {
			stats.TODO++
		} else if strings.HasPrefix(line, "DONE") {
			stats.DONE++
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	data, err := json.MarshalIndent(stats, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	err = os.WriteFile(*outputFile, data, 0644)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Analysis written to %s", *outputFile)

}
