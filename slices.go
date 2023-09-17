package main

import (
	"fmt"
)

func main() {
	a := []int{-3, -2, -1, 0, 1, 2, 3}
	b := []float64{3.14, 9.23, 111.11, 312.12, 1.05}
	c := []string{"shark", "cuttlefish", "squid", "mantis shrimp"}

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

	c = append(c, "seahorse")

	fmt.Println(c)
}
