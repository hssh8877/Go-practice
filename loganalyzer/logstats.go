package main

type LogStats struct {
	Lines int `json:"lines"`
	Info  int `json:"info"`
	Warn  int `json:"warn"`
	Error int `json:"error"`
}
