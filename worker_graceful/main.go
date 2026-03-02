package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

type Job struct {
	ID    int
	Value int
}

func worker(ctx context.Context, id int, jobs <-chan Job, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		select {

		case <-ctx.Done():
			fmt.Printf("Worker %d shutting down\n", id)
			return

		case job, ok := <-jobs:
			if !ok {
				fmt.Printf("Worker %d finished (no more jobs)\n", id)
				return
			}

			fmt.Printf("Worker %d processing job %d\n", id, job.ID)

			select {
			case <-time.After(2 * time.Second):
				results <- job.Value * 2
				fmt.Printf("Worker %d finished job %d\n", id, job.ID)

			case <-ctx.Done():
				fmt.Printf("Worker %d interrupted during job %d\n", id, job.ID)
				return
			}
		}
	}
}

func main() {
	fmt.Println("Worker Pool with Graceful Shutdown")

	numWorkers := 3
	jobs := make(chan Job)
	results := make(chan int)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var wg sync.WaitGroup

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

	go func() {
		sig := <-sigChan
		fmt.Println("\nReceived signal:", sig)
		cancel()
	}()

	for w := 1; w <= numWorkers; w++ {
		wg.Add(1)
		go worker(ctx, w, jobs, results, &wg)
	}

	go func() {
		id := 1
		for {
			select {
			case <-ctx.Done():
				close(jobs)
				return
			default:
				fmt.Println("Sending job", id)
				jobs <- Job{ID: id, Value: id * 10}
				id++
				time.Sleep(1 * time.Second)
			}
		}
	}()

	go func() {
		wg.Wait()
		close(results)
	}()

	for res := range results {
		fmt.Println("Result:", res)
	}

	fmt.Println("Clean shutdown complete")
}
