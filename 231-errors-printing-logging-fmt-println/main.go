package main

import (
	"fmt"
	"os"
)

func main() {
	_, err := os.Open("names.txt")

	if err != nil {
		fmt.Println("error:", err)
		return
	}

}
