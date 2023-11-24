package main

import (
	"flag"
	"fmt"
)

var color = flag.String("color", "blue", "The color of the text")

func main() {
	flag.Parse()
	fmt.Printf("\033[%sm%s\033[0m\n", *color, "Hello, world!")
}
