package main

import "fmt"

func main() {
	//buffered channel - a channel that will allow certain values to sit in that channel 
	//regardless of whether or not something's ready to pull it off.
	ch := make(chan int,1)
	ch <- 42
	fmt.Println(<-ch)
}
