package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"
)

type Usen struct {
	Name          string    `json:"name"`
	Password      string    `json:"-"`
	PreferredFish []string  `json:"preferredFish,omitempty"`
	CreatedAt     time.Time `json:"createdAt"`
}

func main() {
	u := &Usen{
		Name:      "Sammy the shark",
		Password:  "fisharegreat",
		CreatedAt: time.Now(),
	}

	out, err := json.MarshalIndent(u, "", "  ")
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	fmt.Println(string(out))
}
