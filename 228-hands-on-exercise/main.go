package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	tasks := make(chan int)
	numWorkers := 10
	wg.Add(numWorkers)

	for i := 0; i < numWorkers; i++ {
		go send(tasks)
	}

	go func() {
		wg.Wait()
		close(tasks)
	}()

	for v := range tasks {
		fmt.Println(v)
	}
}

func send(tasks chan<- int) {
	for i := 0; i < 10; i++ {
		tasks <- i
	}
	wg.Done()
}
