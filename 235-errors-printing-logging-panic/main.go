package main

import (
	"os"
)

func main() {
	_, err := os.Open("names.txt")

	if err != nil {
		panic(err)
	}

}