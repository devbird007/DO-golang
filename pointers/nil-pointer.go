package main

import "fmt"

type Creation struct {
	Species string
}

func main() {
	var creation *Creation
	creation = &Creation{Species: "shark"}

	fmt.Printf("1) %+v\n", creation)
	changeCreation(creation)
	fmt.Printf("3) %+v\n", creation)
}

func changeCreation(creation *Creation) {
	if creation == nil {
		fmt.Println("creature is nil")
		return
	}

	creation.Species = "jellyfish"
	fmt.Printf("2) %+v\n", creation)
}
