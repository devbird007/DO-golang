/* This demonstrates using an anonymous function to reduce boilerplate
for multiple return value functions */

package main

import (
	"errors"
	"fmt"
	"strings"
)

func capitalize(name string) (string, int, error) {
	handle := func(err error) (string, int, error) {
		return "", 0, err
	}

	if name == "" {
		return handle(errors.New("no name provided"))
	}

	return strings.ToTitle(name), len(name), nil
}

func main() {
	name, size, err := capitalize("samurai jack")
	if err != nil {
		fmt.Println("An error occurred:", err)
	}

	fmt.Printf("Capitalized name: %s, length: %d\n", name, size)
}
