package main

import (
	"errors"
	"fmt"
	"strings"
)

func capitalize(name string) (string, int, error) {
	if name == "" {
		return "", 0, errors.New("no name provided")
	}
	return strings.ToTitle(name), len(name), nil
}

func main() {
	name, digit, err := capitalize("sammy")
	if err != nil {
		fmt.Println("Could not capitalize:", err)
		return
	}

	fmt.Printf("Capitalized name: %s, length: %d\n", name, digit)
}
