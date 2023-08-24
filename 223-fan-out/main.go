package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	tasks := make(chan int)
	results := make(chan int)

	//send
	go send(tasks)

	// Create worker goroutines
	numWorkers := 3
	wg.Add(numWorkers)

	for i := 0; i < numWorkers; i++ {
		//receive
		go process(tasks, results)
	}

	go func(){
		wg.Wait()
		close(results)
	}()
	
	for result := range results {
		fmt.Println("Processed result:", result)
	}

}

func send(tasks chan<- int) {
	for i := 0; i < 10; i++ {
		tasks <- i
	}
	close(tasks)
}

func process(tasks <-chan int, results chan<- int) {
	for v := range tasks {
		results <- v * 2
	}
	wg.Done()
}

//In Go, "fan-out" is a concept used in concurrent programming to distribute tasks or workloads across multiple worker goroutines that operate concurrently. The idea behind fan-out is to take a single input channel and distribute its incoming data to multiple worker goroutines for processing. This allows for efficient utilization of available resources and can lead to improved performance by processing tasks in parallel.
