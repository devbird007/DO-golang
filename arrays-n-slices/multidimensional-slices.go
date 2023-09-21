package main

import (
	"fmt"
)

func main() {
	seaNames := [][]string{{"shark", "octopus", "squid", "mantis shrimp"}, {"Sammy", "Jesse", "Drew", "Jamie"}, {"Cthulu", "Dagon", "Azathoth", "Yoharneth-Lahai"}}

	fmt.Println(seaNames[1][0])
	fmt.Println(seaNames[2][3])
}
