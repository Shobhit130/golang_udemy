package main

import (
	"fmt"
	"time"
)

func main() {
	timeFunc(doWork)
}

func doWork() {
	for i := 0; i < 2_000; i++ {
		fmt.Println(i)
	}
}

func timeFunc(f func()) {
	start := time.Now()
	f()
	elapsed := time.Since(start)
	fmt.Println(elapsed)
}

// timeFunc is acting as a wrapper around the provided function f. It measures the time it takes for f to execute and then prints out the elapsed time. In this context, f is referred to as a callback function because it's a function that is passed as an argument to another function and is intended to be executed at a later point in the code.
