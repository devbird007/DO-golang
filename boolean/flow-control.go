package main

import (
	"fmt"
)

func main() {

	var grade int
	fmt.Println("What is your score?")
	fmt.Scanln(&grade)

	if grade >= 65 {
		fmt.Println("Passing grade.")
	} else {
		fmt.Println("Failing grade.")
	}
}
