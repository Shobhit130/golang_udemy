package main

import "fmt"

func main() {
	ch := make(chan int)

	//send
	go foo(ch)

	//receive
	//if we make both as go routines then they won't have enough time to run and the program will EXIT
	//we don't need to make this a go routine as when it is executed it will call the functions
	//and channel blocks (and our program won't exit) and waits until it gets a value from the channel,
	//now the above go routine will have time to run and it will run
	//after that foo() will put a value in channel and bar() will get a value from channel
	//and then bar() will be unblocked
	bar(ch)

	fmt.Println("about to exit")
}

//send
func foo(ch chan<- int) {
	ch <- 42
}

//receive
func bar(ch <-chan int) {
	fmt.Println(<-ch)
}
