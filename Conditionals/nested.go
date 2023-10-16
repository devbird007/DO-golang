package main

import "fmt"

func main() {
	if statement1 { // outer if
		fmt.Println("hello world")

		if nested_statement1 { // first nested if
			fmt.Println("yes")
		} else if nested_statement2 { // first nested else if
			fmt.Println("maybe")
		} else { // first nested else
			fmt.Println("no")
		}
	} else if statement2 { // outer else if
		fmt.Println("hello galaxy")

		if nested_statement3 { // second nested if
			fmt.Println("yes")
		} else if nested_statement4 { // second nested else if
			fmt.Println("maybe")
		} else { // second nested else
			fmt.Println("no")
		}
	} else { // outer else
		statement("hello universe")
	}
}
