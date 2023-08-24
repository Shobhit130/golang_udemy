package main

import "fmt"

func main() {
	ch := make(chan int)

	//send
	go foo(ch)

	//receive

	//Without closing the channel, the loop in the main function's for v := range ch statement would wait indefinitely, even after all the values have been sent. This would cause a deadlock because the loop would expect more values to be sent through the channel, but the sender (foo function) has finished sending values and won't be able to send the signal that it's done.
	for v := range ch {
		fmt.Println(v)
	}
	fmt.Println("about to exit")
}

//send
func foo(ch chan<- int) {
	for i := 0; i < 100; i++ {
		ch <- i
	}
	//Closing a channel is important to avoid potential deadlocks and to inform the receiving goroutine that it should stop waiting for more values once the sender is done. When a channel is closed, any further attempts to send values through that channel will result in a panic.
	close(ch)
}
