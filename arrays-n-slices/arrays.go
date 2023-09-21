package main

import (
	"fmt"
)

func main() {
	coral := [4]string{"blue coral", "staghorn coral", "pillar coral", "elkhorn coral"}
	fmt.Printf("%q\n", coral)

	// Calling forth the elements in the array
	fmt.Println(coral[0])
	fmt.Println(coral[1])
	fmt.Println(coral[2])
	fmt.Println(coral[3])

	// Concatenating a string element in the array with other strings
	fmt.Println("Sammy loves " + coral[0] + ".")

	// Modifying a string particular string element in the array
	coral[1] = "foliose coral"
	fmt.Printf("%q\n", coral)

	// Length of coral
	fmt.Println(len(coral))

	// It is impossible to modify the length of an array by
	// appending or deleting elements in it.
	// The below will return an error.
	coral = append(coral, "black coral")
}
