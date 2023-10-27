// A basic demo of a nested for loop
package main

import "fmt"

func main() {
	numList := []int{1, 2, 3}
	alphaList := []string{"a", "b", "c"}

	for _, i := range numList {
		fmt.Println(i)
		for _, letter := range alphaList {
			fmt.Println(letter)
		}
	}
}
