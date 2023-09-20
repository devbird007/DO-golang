package main

import (
	"fmt"
	"strconv"
)

func main() {
	a := strconv.Itoa(12)
	fmt.Printf("%q\n", a)

	// In a sentence
	user := "Sammy"
	lines := 50

	fmt.Println("Congratulations, " + user + "! You just wrote " + strconv.Itoa(lines) + " lines of code.")

	// Float to a string
	fmt.Println(fmt.Sprint(421.034))

	f := 5524.53
	fmt.Println(fmt.Sprint(f))

	fmt.Println("Sammy has " + fmt.Sprint(f) + " points.")
}
