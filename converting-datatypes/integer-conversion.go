package main

import "fmt"

func main() {
	var big int64 = 129

	var little int8

	little = int8(big)

	fmt.Println(little)
}
