package main

import (
	"fmt"
	"math/rand"
)

func main() {
	target := rand.Intn(100)

	for {
		var guess int
		fmt.Print("Enter a guess: ")
		_, err := fmt.Scanf("%d", &guess)
		if err != nil {
			fmt.Println("Invalid guess: err:", err)
			continue
		}

		if guess > target {
			fmt.Println("Too high!")
			continue
		}

		if guess < target {
			fmt.Println("Too low!")
			continue
		}

		fmt.Println("You win!")
		break
	}
}
