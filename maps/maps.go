package main

import (
	"fmt"
	"sort"
)

func main() {
	sammy := map[string]string{"name": "Sammy", "animal": "shark", "color": "blue", "location": "ocean"}

	fmt.Println(sammy["name"])
	fmt.Println(sammy["location"])

	// Range and iterate through the map
	for key, value := range sammy {
		fmt.Printf("%q is the key for the value %q.\n", key, value)
	}

	// Retrieve just the keys in the map
	keys := []string{}

	for key := range sammy {
		keys = append(keys, key)
	}
	sort.Strings(keys) // This sorts it out alphabetically
	fmt.Printf("%q\n", keys)

	// Retrieve just the values in the map
	values := []string{}

	for _, value := range sammy {
		values = append(values, value)
	}
	fmt.Printf("%q\n", values)

}
