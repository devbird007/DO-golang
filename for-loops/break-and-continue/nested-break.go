package main

import "fmt"

func main() {
	for outer := 0; outer < 5; outer++ {
		if outer == 3 {
			fmt.Println("Breaking out of outer loop")
			break // break here
		}
		fmt.Println("The value of outer is", outer)
		for inner := 0; inner < 5; inner++ {
			if inner == 2 {
				fmt.Println("Breaking out of inner loop")
				break // break here
			}
			fmt.Println("The value of inner is", inner)
		}
	}
	fmt.Println("Exiting program")
}
