# User Actions Analyzer – Go Project X

*My hardest project so far*

A small Go CLI tool to analyze user login/logout actions from a text file.  

It counts how often each user logs in or out and produces a JSON summary.


## Features

- Reads a plain text file line by line (`bufio.Scanner`)
- Tracks login and logout per user
- Calculates total logins and logouts
- Outputs results as JSON
- CLI flags for input and output files
- Safe error handling with `log.Fatal`

---

## Installation

Clone the repository:

```bash
git clone https://github.com/hssh8877/Go-practice
cd Go-practice/project-x
```

## Usage

```bash
go run .
```

You can override files with CLI flags:
```bash
go run . -file=my_actions.txt -output=my_analysis.json
```

Designed as a mini backend/microservice exercise

Great practice for handling files, maps, structs, and JSON in Go

CLI flags make it flexible for different input/output locations
