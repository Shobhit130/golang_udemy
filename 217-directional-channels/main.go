package main

import "fmt"

func main() {
	//making a send-only channel
	ch := make(chan <- int,2)

	ch <- 42
	ch <- 43
	
	//error - cannot receive from send-only channel c
	// fmt.Println(<-ch)
	// fmt.Println(<-ch)

	fmt.Println("--------")
	fmt.Printf("%T\n",ch)
	
	fmt.Println("--------")
	
	//making a receive-only channel
	ch1 := make(<-chan int,2)
	
	//error - cannot send to receive-only channel ch1
	// ch1 <- 42
	// ch1 <- 43
	
	fmt.Printf("%T\n",ch1)
	
	//we can convert a channel to a send or receive only channel
	
	fmt.Println("--------")
	ch2 := make(chan int,1)
	
	fmt.Printf("%T\n",(chan<- int)(ch2)) //converting to a send-only channel
	fmt.Printf("%T\n",(<-chan int)(ch2)) //converting to a receive-only channel

}
