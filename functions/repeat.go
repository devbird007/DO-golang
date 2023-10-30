package main

import "fmt"

func main() {
	val, err := repeat("Sammy", -1)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(val)
}

func repeat(word string, reps int) (string, error) {
	if reps <= 0 {
		return "", fmt.Errorf("Invalid value of %d provided for reps. Value must be greater than 0.", reps)
	}

	var value string
	for i := 0; i < reps; i++ {
		value = value + word
	}
	return value, nil
}
