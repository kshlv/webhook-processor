package main

import (
	"log"
)

func main() {
	config := NewConfig()
	if err := Start(config); err != nil {
		log.Fatal(err)
	}
}
