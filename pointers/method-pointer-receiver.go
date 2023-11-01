package main

import "fmt"

type Animal struct {
	Species string
}

func (a *Animal) Reset() {
	a.Species = ""
}

func main() {
	var animal Animal = Animal{Species: "shark"}

	fmt.Printf("1) %+v\n", animal)
	animal.Reset()
	fmt.Printf("2) %+v\n", animal)
}
