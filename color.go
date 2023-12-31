package main

import "fmt"

const favColor string = "blue" // Could have chosen any color

func main() {
	var guess string
	// Create an input loop
	for {
		// Ask the user to guess my favorite color
		fmt.Println("Guess my favorite color:")

		// Try to read a line of input from the user.
		// Print out an error and exit, if there is one.
		if _, err := fmt.Scanln(&guess); err != nil {
			fmt.Printf("%s\n", err)
			return
		}

		// Did they guess the correct color?
		if favColor == guess {
			// They guessed it!
			fmt.Printf("%q is my favorite color!\n", favColor)
			return
		}

		// Wrong! Have them guess again.
		fmt.Printf("Sorry, %q is not my favorite color. Guess again.\n", guess)
	}
}
