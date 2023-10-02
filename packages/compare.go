package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func main() {
	// User must pass in number of integers to generate
	if len(os.Args) < 2 {
		println("Usage:\n")
		println(" go run compare.go <numberOfInts>")
		return
	}
	n, err := strconv.Atoi(os.Args[1])
	if err != nil {
		// Maybe they didn't pass an integer
		panic(err)
	}

	// Phase 1 - Using math/rand
	// Initialize the byte slice
	b1 := make([]byte, n)

	// Get the time
	start := time.Now()

	// Generate the pseudo-random numbers
	for i := 0; i < 5; i++ {
		b1[i] = byte(rand.Intn(256)) // Where the magic happens!
	}

	// Compute the elapsed time
	elapsed := time.Now().Sub(start)

	// In case n is very large, only print a few numbers
	for i := 0; i < 5; i++ {
		println(b1[i])
	}
	fmt.Printf("Time to generate %v pseudo-random numbers with math/rand: %v\n", n, elapsed)
}
