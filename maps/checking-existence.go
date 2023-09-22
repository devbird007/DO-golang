package main

import "fmt"

func main() {
	counts := map[string]int{}
	fmt.Println(counts["sammy"])

	// Using *ok* idiom in a conditional statement
	count, ok := counts["sammy"]
	if ok {
		fmt.Printf("Sammy has a count of %d\n", count)
	} else {
		fmt.Println("Sammy was not found.")
	}

	// Repeating above but combining variable declaration and conditional checking
	if count, ok := counts["sammy"]; ok {
		fmt.Printf("Sammy has a count of %d\n", count)
	} else {
		fmt.Println("Sammy was not found.")
	}
}

// Whenever you wish to get a particular value from a map in Go,
// it is important to also verify the value's existence in the map.
