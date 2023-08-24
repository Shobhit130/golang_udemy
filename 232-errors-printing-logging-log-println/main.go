package main

import (
	"log"
	"os"
)

func main() {
	_, err := os.Open("names.txt")

	if err != nil {
		log.Println("error:", err)
		return
	}

}
