package main

import (
	"errors"
	"fmt"
	"time"
)

func main() {
	// Using errors.New
	err := errors.New("barnacles")
	fmt.Println("Sammy says:", err)

	// Using fmt.Errorf
	ess := fmt.Errorf("error occured at %v", time.Now())
	fmt.Println("An error happened:", ess)
}
