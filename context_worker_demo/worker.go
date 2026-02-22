package main

import (
	"context"
	"fmt"
	"sync"
)

func worker(ctx context.Context, id int, jobs <-chan Job, results chan<- Result, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		select {
		case job, ok := <-jobs:
			if !ok {
				fmt.Printf("Worker %d: jobs channel closed\n", id)
				return
			}

			output := job.Value * 2
			fmt.Printf("Worker %d: processed job %d\n", id, job.ID)

			results <- Result{
				JobID:  job.ID,
				Output: output,
			}

		case <-ctx.Done():
			fmt.Printf("Worker %d: stopped by context\n", id)
			return
		}
	}
}
