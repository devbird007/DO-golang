package main

import (
	"fmt"
)

func main() {
	a := "Hello, 世界"
	for i, c := range a {
		fmt.Printf("%d: %s\n", i, string(c))
	}
	fmt.Println("length of 'Hello, 世界': ", len(a))
}
