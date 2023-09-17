package main

import (
	"fmt"
)

func main() {
	s := "This string\nspans multiple\nlines."
	fmt.Println(s)
	fmt.Println("1.\tShark\n2.\tShrimp\n10.\tSquid\n")

	fmt.Println(
		`This string is on
		multiple lines
		within three single
		quotes on either side.`,
	)
	fmt.Println(
		"The string is on \n" +
			"multiple lines\n" +
			"within three single\n" +
			"quotes on either side.",
	)

	fmt.Println(`Sammy says,\"The balloon\'s is color is red.\"`)
}
