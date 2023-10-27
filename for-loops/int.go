// Using a RangeClause For loop to populate the values of an
// slice.
package main

import "fmt"

func main() {
	integers := make([]int, 10)
	fmt.Println(integers)

	for i := range integers {
		integers[i] = i
	}

	fmt.Println(integers)
}
