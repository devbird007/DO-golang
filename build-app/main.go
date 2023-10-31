package main

import "fmt"

var features = []string{
	"Free Features #1",
	"Free Features #2",
}

func main() {
	for _, f := range features {
		fmt.Println(">", f)
	}
}
