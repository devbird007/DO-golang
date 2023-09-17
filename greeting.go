// This is a doc comment for greeting.go.
//   - prompt user for name.
//   - wait for name
//   - print name.
//
// This is the second paragraph of this doc comment.
// `gofmt` (and `go doc`) will insert a blank line before it.
package main

import (
	"fmt"
	"strings"
)

func main() {
	// This is not a doc comment. Gofmt will NOT format it.
	// 	- prompt user for name
	// 		- wait for name
	//			- print name
	// This is not a "second paragraph" because this is not a doc comment.
	// It's just more lines to this non-doc comment.
	fmt.Println("Please enter your name.")
	var name string
	fmt.Scanln(&name)
	name = strings.TrimSpace(name)
	fmt.Printf("Hi, %s! I'm Go!\n", name)
}
