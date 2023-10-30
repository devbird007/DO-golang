package main

import (
	"fmt"
	"strings"
)

func names() {
	fmt.Println("Enter your name:")

	var name string
	fmt.Scanln(&name)
	// Check whether the name has a vowel
	for _, v := range strings.ToLower(name) {
		if v == 'a' || v == 'e' || v == 'i' || v == 'o' || v == 'u' {
			fmt.Println("Your name contains a vowel.")
			return
		}
	}
	fmt.Println("Your name does not contain a vowel.")
}

func main() {
	names()
}
