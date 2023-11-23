package main

import (
	"app/build"
	"fmt"
)

var Version = "development"

func main() {
	fmt.Println("Version:\t", Version)
	fmt.Println("build.Time:\t", build.Time)
	fmt.Println("build.User:\t", build.User)
}
