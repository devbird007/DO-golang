package main

import (
	"fmt"
)

func main() {
	// Creating a slice
	seaCreatures := []string{"shark", "cuttlefish", "squid", "mantis shrimp", "anemone"}
	fmt.Printf("%q\n", seaCreatures)

	// Initializing an empty slice with preallocated length & memory
	oceans := make([]string, 3, 5)
	fmt.Printf("%q\n", oceans)

	// Slicing a created array into slices
	coral := [4]string{"blue coral", "foliose coral", "pillar coral", "elkhorn coral"}

	fmt.Printf("%q\n", coral[1:])

	// Copying an Array into a new replica Slice
	coralSlice := coral[:]
	coralSlice = append(coralSlice, "black coral", "antipathes", "leptopsammia")
	fmt.Printf("%q\n", coralSlice)

	// Conjoining two slices
	moreCoral := []string{"massive coral", "soft coral"}
	coralSlice = append(coralSlice, moreCoral...)
	fmt.Printf("%q\n", coralSlice)

	// Deleting an element from a slice
	/* coralSlice := []string{"blue coral", "foliose coral", "pillar coral", "elkhorn coral", "black coral", "antipathes", "leptopsammia", "massive coral", "soft coral"} */

	coralSlice = append(coralSlice[:3], coralSlice[6:]...)
	fmt.Printf("%q\n", coralSlice)

}
