package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

const workerRuntime = 2 * time.Second

func main() {
	numJobs := 5
	numWorkers := 2

	// Context used to gracefully stop workers
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	jobs := make(chan Job, numJobs)
	results := make(chan Result, numJobs)

	var wg sync.WaitGroup

	// Start workers
	for w := 1; w <= numWorkers; w++ {
		wg.Add(1)
		go worker(ctx, w, jobs, results, &wg)
	}

	// Send jobs
	for j := 1; j <= numJobs; j++ {
		jobs <- Job{ID: j, Value: j}
	}
	close(jobs)

	// Let workers process for a short time, then cancel
	time.Sleep(workerRuntime)
	cancel()

	wg.Wait()
	close(results)

	for res := range results {
		fmt.Printf("Result: %+v\n", res)
	}
}
