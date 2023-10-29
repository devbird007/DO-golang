// Exploding a slice so it can be parsed through a variadic
// function.

package main

import "fmt"

func main() {
	var line string

	line = expJoin(",", []string{"Sammy", "Jessica", "Drew", "Jamie"}...)
	fmt.Println(line)

	line = expJoin(",", "Sammy", "Jessica", "Drew", "Jamie")
	fmt.Println(line)

	line = expJoin(",", "Sammy", "Jessica")
	fmt.Println(line)

	line = expJoin(",", "Sammy")
	fmt.Println(line)

}

func expJoin(del string, values ...string) string {
	var line string
	for i, v := range values {
		line = line + v
		if i != len(values)-1 {
			line = line + del
		}
	}
	return line
}
