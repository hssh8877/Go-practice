package main

import (
	"fmt"
	"sort"
	"time"
)

func PrintStats(results []WorkerResult) {
	// nach Dauer sortieren (aufsteigend)
	sort.Slice(results, func(i, j int) bool {
		return results[i].Duration < results[j].Duration
	})

	// Durchschnitt berechnen
	var total time.Duration
	for _, r := range results {
		total += r.Duration
	}
	avg := total / time.Duration(len(results))

	fmt.Println("================================================================")
	fmt.Println("_-=Worker Stats=-_")
	for _, r := range results {
		fmt.Printf("Worker %d finished in %.2f s\n", r.ID, r.Duration.Seconds())
	}
	fmt.Printf("\nAverage duration: %.2f s\n", avg.Seconds())
	fmt.Println("================================================================")
}
