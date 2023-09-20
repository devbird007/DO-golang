package main

import (
	"fmt"
	"log"
	"strconv"
)

func main() {
	lines_yesterday := "50"
	lines_today := "108"

	yesterday, err := strconv.Atoi(lines_yesterday)
	if err != nil {
		log.Fatal(err)
	}

	today, err := strconv.Atoi(lines_today)
	if err != nil {
		log.Fatal(err)
	}
	lines_more := today - yesterday

	fmt.Println(lines_more)

	// Converting a string that is not a number to view the error
	a := "not a number"
	b, err := strconv.Atoi(a)
	fmt.Println(b)
	fmt.Println(err)
}
