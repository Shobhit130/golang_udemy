package main

import "fmt"

//All go routines are asleep - Deadlock.
//And that's because channels blocked. And so my code starts at package mains where my program enters goes to func main. 
//That's the entry point for my program.
//It makes a channel and then it tries to put 42 on the channel and it blocks.
//Because when you send and receive on a channel, it's like relay racers, racers in a track race that
//they have to pass a baton and they have to pass it hand to hand.
//It can't occur until both send and receive can happen at the same time.
//And if they can't happen at the same time, it blocks the send and receive blocks until the receiver is ready to pull it off.

func main() {
	//this is blocking because my program enters main, makes a channel
	ch := make(chan int)
	//gets right here and there's nothing to pull it off.
	ch <- 42
	//You can't jump down here because this can't get on until something can pull it off.
	fmt.Println(<-ch)
}
