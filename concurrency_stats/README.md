# Concurrency Stats – Go Project

A small Go CLI tool demonstrating **concurrency with Goroutines**.  
It simulates multiple workers running in parallel, measures their execution duration, and prints statistics.

---

## Features

- Starts multiple Goroutines (`worker`) in parallel
- Measures and prints the duration of each worker
- Calculates average execution time
- Results are displayed in a clean, readable format
- Uses WaitGroup for synchronization
- Randomized execution time for each worker (1–5 seconds)

---

## Installation

Clone the repository:

```bash
git clone https://github.com/hssh8877/Go-practice
cd Go-practice/concurrency_stats
```

## Usage

```bash
go run .
```

## Project Goals

- Practice and learn concurrency in Go
- Understand Goroutines, Channels, and WaitGroups
- Learn to measure execution time for parallel tasks

## Notes

It helped me understand how to run parallel tasks and collect results efficiently.
