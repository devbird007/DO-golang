// This demonstrates the capacities of the join function
// when declared as a variadic function.

package main

import "fmt"

func main() {
	var line string

	line = updatedJoin(",", "Sammy", "Jessica", "Drew", "Jamie")
	fmt.Println(line)

	line = updatedJoin(",", "Sammy", "Jessica")
	fmt.Println(line)

	line = updatedJoin(",", "Sammy")
	fmt.Println(line)
}

func updatedJoin(del string, values ...string) string {
	var line string
	for i, v := range values {
		line = line + v
		if i != len(values)-1 {
			line = line + del
		}
	}
	return line
}
