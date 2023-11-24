package main

import (
	"flag"
	"fmt"

	"github.com/fatih/color"
)

func main() {
	useColor := flag.Bool("color", false, "display colorized output")
	flag.Parse()

	if *useColor {
		green := color.New(color.FgHiMagenta).Add(color.Underline)
		green.Printf("You are welcome in the name of the Lord!\n")

		return
	}
	fmt.Println("You are welcome in the name of the Lord!")
}
