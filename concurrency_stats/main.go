package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	var wg sync.WaitGroup
	results := make(chan WorkerResult, 10)

	for i := 1; i <= 10; i++ {
		wg.Add(1)
		go worker(i, &wg, results)
	}

	wg.Wait()
	close(results)

	var stats []WorkerResult
	for r := range results {
		stats = append(stats, r)
	}
	PrintStats(stats)

	fmt.Println("----All workers finished. Main Finished----")
}
