package main

type UserStats struct {
	Login  int `json:"login"`
	Logout int `json:"logout"`
}

type OverallStats struct {
	Users        map[string]*UserStats `json:"users"`
	TotalLogins  int                   `json:"total_logins"`
	TotalLogouts int                   `json:"total_logouts"`
}
