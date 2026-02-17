# Go Log-Analyzer

A simple CLI tool written in Go to analyze log files.  
It reads a log file, counts the number of INFO, WARN, and ERROR messages, and outputs the results as structured JSON.

## Features

- Reads log files line by line using `bufio.Scanner`
- Counts INFO, WARN, and ERROR entries
- Tracks total number of lines
- Outputs analysis in JSON (`analysis.json`)
- Safe file handling with `defer` and proper error checks
- Easily extendable for future features

## Installation

Clone the repository:

```bash
git clone https://github.com/hssh8877/Go-practice
cd Go-practice/log-analyzer


Usage

#Run the program:

go run .
```

The programm reads sample.log in the folder and generates anaylsis.json with the results.
For Example:
{
  "lines": 9,
  "info": 5,
  "warn": 2,
  "error": 2
}


