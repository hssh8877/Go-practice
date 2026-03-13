Buffered Worker Pool – Go Project

A small Go project demonstrating a simple buffered worker pool in Go. It simulates multiple workers processing tasks in parallel and collects the results in a channel.

Features

Starts multiple Goroutines (workers) in parallel

Processes tasks from a buffered channel

Collects results in a separate channel

WaitGroup ensures all workers finish

Demonstrates concurrent processing with non-deterministic order

## Installation

Clone the repository:
```
git clone https://github.com/hssh8877/Go-practice
cd Go-practice/buffered_worker_demo
```
## Usage
```
go run .
```
Observe workers processing tasks.
Results are printed from the results channel.
All tasks are processed before the program exits.

Project Goals

Practice and learn Go concurrency patterns

Understand Goroutines, buffered Channels, and WaitGroups

Learn basic task distribution in a worker pool
