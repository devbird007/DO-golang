package main

import "fmt"

func main() {
	var f float64 = 390.8
	var i int = int(f)

	fmt.Printf("f = %.2f\n", f)
	fmt.Printf("i = %d\n", i)

	b := 125.0
	c := 390.8

	fmt.Println(int(b))
	fmt.Println(int(c))
}
