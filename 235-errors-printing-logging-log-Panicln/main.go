package main

import (
	"log"
	"os"
)

func main() {
	_, err := os.Open("names.txt")

	if err != nil {
		log.Panicln(err)
	}

}

//Panicln is equivalent to Println followed by a call to panic


//The panic built-in function stops normal execution of the current goroutine. When a function F calls panic, normal execution of F stops immediately. Any functions whose execution was deferred by F are run in the usual way, and then F returns to its caller.
