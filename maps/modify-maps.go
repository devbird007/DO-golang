package main

import (
	"fmt"
)

func main() {
	// Add new key-value to a map
	usernames := map[string]string{"Sammy": "sammy-shark", "Jamie": "mantisshrimp54"}

	usernames["Drew"] = "squidly"
	fmt.Println(usernames)

	// Add new value to a key
	followers := map[string]int{"drew": 305, "mary": 428, "cindy": 918}
	followers["drew"] = 342
	fmt.Println(followers)
}
