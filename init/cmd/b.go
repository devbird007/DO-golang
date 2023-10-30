package main

import (
	"fmt"

	"init/message"
)

func init() {
	message.Message = "Hello"
	fmt.Println("b ->", message.Message)
}
