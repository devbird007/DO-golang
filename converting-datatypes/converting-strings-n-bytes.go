package main

import (
	"fmt"
)

func main() {
	a := "my string"

	b := []byte(a)

	c := string(b)

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}
