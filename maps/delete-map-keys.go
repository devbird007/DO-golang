package main

import (
	"fmt"
)

func main() {
	permissions := map[int]string{1: "read", 2: "write", 4: "delete", 8: "create", 16: "modify"}

	// To delete a select key from the map
	delete(permissions, 16)
	fmt.Println(permissions)

	// To clear the map of all its values
	permissions = map[int]string{}
	fmt.Println(permissions)
}
