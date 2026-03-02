# Worker Pool with Graceful Shutdown

This Go program demonstrates a **worker pool** with:

- parallel job processing  
- `context.Context` for graceful shutdown  
- signal handling (`Ctrl+C`) for clean termination of running jobs  
- proper closing of all workers and channels

## Features

- Dynamic job distribution to multiple workers  
- Context allows interrupting ongoing tasks  
- Results are printed even if some jobs are interrupted  
- Great learning project for Go concurrency patterns

## How to Run

Clone the repository:
```bash
git clone https://github.com/hssh8877/Go-practice
cd Go-practice/worker_graceful
```

## Usage
```
go run .
```

## Notes


Channel sizes, number of workers, and job durations can be easily adjusted. 
Ctrl+C (graceful shutdown) works reliably in a terminal.
In some IDEs (like Flatpak/Snap GoLand), signal handling might behave differently.
