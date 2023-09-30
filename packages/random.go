package main

import (
	"math/rand"
	"time"
)

func main() {
	now := time.Now()
	println("Numbers seeded using current date/time:", now.Format(time.StampNano))
	for i := 0; i < 5; i++ {
		println(rand.Intn(10))
	}
}
