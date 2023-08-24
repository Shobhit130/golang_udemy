package main

import (
	"log"
	"os"
)

func main() {
	_, err := os.Open("names.txt")

	if err != nil {
		log.Fatalln(err)
	}

}

//Fatalln is equivalent to Println followed by a call to os.Exit(1)
//deferred functions don't run
