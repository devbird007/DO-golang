package main

import (
	"fmt"
)

func main() {
	fmt.Println((9 > 7) && (2 < 4))   // Both are true
	fmt.Println((8 == 8) || (6 != 6)) // One of them is true
	fmt.Println(!(3 <= 1))            // The statement is false

	fmt.Println((-0.2 > 1.4) && (0.8 < 3.1))
	fmt.Println((7.5 == 8.9) || (9.2 != 9.2))
	fmt.Println(!(-5.7 <= 0.3))

	fmt.Println(!((-0.2 > 1.4) && ((0.8 < 3.1) || (0.1 == 0.1))))

}
