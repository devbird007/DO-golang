package main

import (
	"fmt"
	"math"
)

func main() {
	i := 3.3
	fmt.Println(+i)

	j := -19
	fmt.Println(-j)

	k := 100.2
	l := 10.2
	fmt.Println(k * l)

	m := 80
	n := 6
	r := float64(m) / float64(n)
	fmt.Println(r)

	o := 85
	p := 15
	fmt.Println(o % p)

	q := 36.0
	c := 8.0
	s := math.Mod(q, c)
	fmt.Println(s)
}
