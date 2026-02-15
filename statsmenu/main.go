package main

import "fmt"

func main() {
	statsObj := stats{}
	var choice int

	for {
		fmt.Println("====Menü====")
		fmt.Println("1. Add numbers")
		fmt.Println("2. Show numbers")
		fmt.Println("3. Sum numbers")
		fmt.Println("4. Highest value")
		fmt.Println("5. Lowest value")
		fmt.Println("6. Average")
		fmt.Println("7. Remove a value")
		fmt.Println("8. Save numbers")
		fmt.Println("9. Load numbers")
		fmt.Println("10. Exit")
		fmt.Println("Choose your  number (1-10)")

		fmt.Scan(&choice)

		switch choice {
		case 1:
			var x int
			fmt.Print("Which number shall be added? ")
			fmt.Scan(&x)
			statsObj.add(x)
			fmt.Printf("%d got added\n", x)

		case 2:
			if len(statsObj.numbers) == 0 {
				fmt.Println("There are no numbers")
			} else {
				fmt.Println("Numbers:", statsObj.numbers)
			}

		case 3:
			fmt.Println("Sum: ", statsObj.sum())
		case 4:
			fmt.Println("Highest Value: ", statsObj.max())
		case 5:
			fmt.Println("Lowest Value: ", statsObj.min())
		case 6:
			fmt.Printf("Average: %.2f\n", statsObj.mean())

		case 7:
			if len(statsObj.numbers) == 0 {
				fmt.Println("No numbers to remove.")
				break
			}

			var x int
			fmt.Print("Which number should be removed? ")
			fmt.Scan(&x)

			if statsObj.remove(x) {
				fmt.Printf("%d got removed\n", x)
			} else {
				fmt.Println("Number not found.")
			}
		case 8:
			err := statsObj.save("numbers.json")
			if err != nil {
				fmt.Println("Error saving:", err)
			} else {
				fmt.Println("Numbers saved.")
			}

		case 9:
			err := statsObj.load("numbers.json")
			if err != nil {
				fmt.Println("Error loading:", err)
			} else {
				fmt.Println("Numbers loaded.")
			}

		case 10:
			fmt.Println("Exiting Program...")
			return
		default:
			fmt.Println("Invalid choice. Please try again. (1-10")
		}

	}
}
