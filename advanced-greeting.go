package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Heyyy, what's your name?")
	var name string
	fmt.Scanln(&name)
	name = strings.TrimSpace(name)
	fmt.Printf("It's nice to meet you %s.\n", name)

	fmt.Printf("%s, what's your favorite color?\n", name)
	var color string
	fmt.Scanln(&color)
	fmt.Printf("No way %s, %s is a very nice color. \nMy own favorite color is red. \n", name, color)
}
