package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	fmt.Println("Main go routine")

	wg.Add(2)
	go foo()
	go bar()
	wg.Wait()
}

func foo(){
	for i:=0;i<10;i++{
		fmt.Println("foo:",i)
	}
	wg.Done()
}
func bar(){
	for i:=0;i<10;i++{
		fmt.Println("bar:",i)
	}
	wg.Done()
}
