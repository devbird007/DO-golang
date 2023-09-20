package main

import (
	"fmt"
)

func main() {
	sammy := map[string]string{"name": "Sammy", "animal": "shark", "color": "blue", "location": "ocean"}

	items := make([]string, len(sammy))

	var i int

	for _, v := range sammy {
		items[i] = v
		i++
	}

	fmt.Printf("%q\n", items)
}
