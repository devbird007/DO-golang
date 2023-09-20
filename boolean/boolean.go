package main

import (
	"fmt"
)

func main() {
	x := 5
	y := 8

	fmt.Println("x == y:", x == y)
	fmt.Println("x != y:", x != y)
	fmt.Println("x < y:", x < y)
	fmt.Println("x > y:", x > y)
	fmt.Println("x <= y:", x <= y)
	fmt.Println("x >= y:", x >= y)

	Sammy := "Sammy"
	sammy := "sammy"
	alsoSammy := "Sammy"

	fmt.Println("Sammy == sammy:", Sammy == sammy)
	fmt.Println("Sammy == alsoSammy", Sammy == alsoSammy)

	t := true
	f := false

	fmt.Println("t != f:", t != f)
}
