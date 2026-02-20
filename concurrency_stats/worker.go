package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type WorkerResult struct {
	ID       int
	Duration time.Duration
}

func worker(id int, wg *sync.WaitGroup, results chan<- WorkerResult) {
	defer wg.Done()

	start := time.Now()
	duration := rand.Intn(5) + 1
	fmt.Printf("Worker %d started, duration: %.2f s\n", id, float64(duration))

	time.Sleep(time.Duration(duration) * time.Second)

	elapsed := time.Since(start)
	results <- WorkerResult{ID: id, Duration: elapsed}
}
