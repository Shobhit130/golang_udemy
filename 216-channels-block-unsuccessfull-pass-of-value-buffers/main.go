package main

import "fmt"

func main() {
	ch := make(chan int,1)
	ch <- 42
	//fatal error: all goroutines are asleep - deadlock!
	//we have room for one value only in buffer
	ch <- 43
	fmt.Println(<-ch)
}
