package main

import "fmt"

const (
	year     = 365
	leapYear = int32(366)
)

func main() {
	hours := 24
	minutes := int32(60)
	fmt.Println(hours * year)
	fmt.Println(minutes * leapYear)
	fmt.Println(minutes * year)
}
