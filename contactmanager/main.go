package main

import (
	"fmt"
)

type Contact struct {
	Name  string
	Email string
	Phone string
}

func main() {
	var contacts []Contact
	var choice int

	for {
		fmt.Println("===Contacts===")
		fmt.Println("1. Add contact")
		fmt.Println("2. Show contacts")
		fmt.Println("3. Search contact")
		fmt.Println("4. Delete contact")
		fmt.Println("5. Save contacts")
		fmt.Println("6. Load contacts")
		fmt.Println("7. Exit")
		fmt.Println("Please choose (1-7)")

		fmt.Scan(&choice)

		switch choice {
		case 1:
			var name, email, phone string
			fmt.Print("Name: ")
			fmt.Scan(&name)
			fmt.Print("Email: ")
			fmt.Scan(&email)
			fmt.Print("Phone: ")
			fmt.Scan(&phone)

			contacts = append(contacts, Contact{name, email, phone})
			fmt.Println("Contact added!")

		case 2:
			if len(contacts) == 0 {
				fmt.Println("No contacts found!")
			} else {
				for _, c := range contacts {
					fmt.Printf("Name: %s |Email:  %s |Phone: %s\n", c.Name, c.Email, c.Phone)
				}
			}

		case 3:
			var search string
			fmt.Print("Name? ")
			fmt.Scan(&search)
			found := false
			for _, c := range contacts {
				if c.Name == search {
					fmt.Printf("Contact found: Name: %s | Phone: %s | Email: %s\n", c.Name, c.Email, c.Phone)
					found = true
				}
			}
			if !found {
				fmt.Println("No contact found!")
			}

		case 4:
			var delName string
			fmt.Print("Name to delete? ")
			fmt.Scan(&delName)

			found := false
			for i, c := range contacts {
				if c.Name == delName {
					contacts = append(contacts[:i], contacts[i+1:]...)
					found = true
					break
				}
			}

			if found {
				fmt.Println("Contact deleted!")
			} else {
				fmt.Println("No contact found!")
			}

		case 5:
			err := saveContacts("contacts.json", contacts)
			if err != nil {
				fmt.Println("Error saving contacts: ", err)
			} else {
				fmt.Println("Contacts saved to contacts.json")
			}

		case 6:
			err := loadContacts("contacts.json", &contacts)
			if err != nil {
				fmt.Println("Error loading contacts: ", err)
			} else {
				fmt.Println("Contacts loaded from contacts.json")
			}

		case 7:
			fmt.Println("Exiting....")
			return

		default:
			fmt.Println("Invalid choice!")
		}
	}
}
