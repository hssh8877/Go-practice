# Task Analyzer (Go)

A simple CLI tool written in Go that analyzes a task list file.
It counts TODO and DONE entries and outputs the result as structured JSON.

## Features

- Reads a text file line by line using `bufio.Scanner`
- Counts TODO and DONE tasks
- Outputs analysis as formatted JSON
- Supports CLI flags for flexible input and output files
- Safe file handling with `defer` and proper error checks
- Easy to extend

## Installation

Clone the repository:

```bash
git clone https://github.com/hssh8877/Go-practice
cd Go-practice/task-analyzer
```

## Usage

Run with default files (tasks.txt → tasks.json):

```bash
go run .
```

Specify custom input and output files:

```bash
go run . --file=tasks.txt --output=tasks.json
```
