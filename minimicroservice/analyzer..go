package main

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func Analyze(filename string) *OverallStats {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	stats := &OverallStats{
		Users: make(map[string]*UserStats),
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Fields(line)
		if len(parts) != 2 {
			continue
		}

		action := strings.ToUpper(parts[0])
		user := parts[1]

		if _, exists := stats.Users[user]; !exists {
			stats.Users[user] = &UserStats{}
		}

		switch action {
		case "LOGIN":
			stats.Users[user].Login++
			stats.TotalLogins++
		case "LOGOUT":
			stats.Users[user].Logout++
			stats.TotalLogouts++
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return stats
}
