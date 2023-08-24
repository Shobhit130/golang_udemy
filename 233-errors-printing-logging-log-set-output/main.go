package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	f, err := os.Create("./log.txt")

	if err != nil {
		log.Println("error:", err)
		return
	}

	defer f.Close()
	log.SetOutput(f)

	f2, err := os.Open("names.txt")

	if err != nil {
		log.Println("error:", err)
	}

	defer f2.Close()

	fmt.Println("Check the log.txt file")

}
