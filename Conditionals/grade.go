package main

import "fmt"

func main() {
	grade := 96

	if grade >= 65 {
		fmt.Println("Passing grade of: ")

		if grade >= 90 {
			if grade > 96 {
				fmt.Println("A+")

			} else if grade > 93 && grade <= 96 {
				fmt.Println("A")

			} else {
				fmt.Println("A-")
			}

		} else if grade >= 80 {
			fmt.Println("B")

		} else if grade >= 70 {
			fmt.Println("C")

		} else if grade >= 65 {
			fmt.Println("D")
		}
	} else {
		fmt.Println("Failing grade.")
	}
}
