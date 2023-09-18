package main

import (
	"fmt"
	"strings"
)

func main() {
	openSource := "Sammy contributes to open source."
	fmt.Println(len(openSource))

	// strings.Join
	fmt.Println(strings.Join([]string{"sharks", "crustaceans", "plankton"}, " "))
	fmt.Println(strings.Join([]string{"sharks", "crustaceans", "plankton"}, ", "))

	// strings.Split
	balloon := "Sammy has a balloon."
	s := strings.Split(balloon, " ")
	fmt.Printf("%q\n", s)
	fmt.Println(len(s))

	// strings.Fields
	data := " username password     	email	date"
	fields := strings.Fields(data)
	fmt.Printf("%q\n", fields)
	fmt.Println(fields)
}
