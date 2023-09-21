package main

import (
	"fmt"
)

func main() {
	// Populating a slice with a for loop that appends a new element
	numbers := []int{}
	for i := 0; i < 4; i++ {
		numbers = append(numbers, i)
	}
	fmt.Println(numbers)

	// Populating a slice by pre-allocating capacity with the cap() function.
	// This is more memory efficient than the above
	digits := make([]int, 4)
	for i := 0; i < cap(digits); i++ {
		digits[i] = i
	}
	fmt.Println(digits)
}
