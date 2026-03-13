package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Task struct
type Task struct {
	ID    int
	Value int
}

func worker(id int, tasks <-chan Task, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for task := range tasks {
		duration := time.Duration(rand.Intn(3)+1) * time.Second
		fmt.Printf("Worker %d processing task %d for %v\n", id, task.ID, duration)
		time.Sleep(duration)
		results <- task.Value * 2
		fmt.Printf("Worker %d finished task %d\n", id, task.ID)
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	fmt.Println("Buffered Worker Pool Demo")

	numTasks := 10
	numWorkers := 3
	taskChan := make(chan Task, 5)
	resultChan := make(chan int, numTasks)

	var wg sync.WaitGroup

	for w := 1; w <= numWorkers; w++ {
		wg.Add(1)
		go worker(w, taskChan, resultChan, &wg)
	}

	go func() {
		for i := 1; i <= numTasks; i++ {
			taskChan <- Task{ID: i, Value: i * 10}
			fmt.Printf("Task %d sent to queue\n", i)
		}
		close(taskChan)
	}()

	go func() {
		wg.Wait()
		close(resultChan)
	}()

	for res := range resultChan {
		fmt.Println("Result received:", res)
	}

	fmt.Println("All tasks processed, main finished")
}
