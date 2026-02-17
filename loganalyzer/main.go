package main

import (
	"bufio"
	"encoding/json"
	"log"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("sample.log")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	stats := LogStats{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		stats.Lines++

		line := scanner.Text()

		if strings.Contains(line, "ERROR") {
			stats.Error++
		} else if strings.Contains(line, "WARN") {
			stats.Warn++
		} else if strings.Contains(line, "INFO") {
			stats.Info++
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	data, err := json.MarshalIndent(stats, "", "  ")
	if err != nil {
		log.Fatal(err)
	}

	err = os.WriteFile("Analysis.json", data, 0644)
	if err != nil {
		log.Fatal(err)
	}
}
