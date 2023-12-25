package main

import (
	"fmt"
)

func main() {
	names := []string{
		"lobster",
		"sea urchin",
		"sea cucumber",
	}
	fmt.Println("My favorite sea creature is:", names[len(names)])
}
