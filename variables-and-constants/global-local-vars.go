package main

import "fmt"

var num1 = 5

func printNumbers() {
	num1 := 10
	num2 := 7

	fmt.Println(num1)
	fmt.Println(num2)
}

func main() {
	printNumbers()
	fmt.Println(num1)
}
