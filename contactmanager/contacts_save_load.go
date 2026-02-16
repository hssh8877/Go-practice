package main

import (
	"encoding/json"
	"os"
)

func saveContacts(filename string, contacts []Contact) error {
	data, err := json.MarshalIndent(contacts, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(filename, data, 0644)
}

func loadContacts(filename string, contacts *[]Contact) error {
	data, err := os.ReadFile(filename)
	if err != nil {
		return err
	}
	return json.Unmarshal(data, contacts)
}
