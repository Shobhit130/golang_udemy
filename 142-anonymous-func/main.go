package main

import "fmt"

func main() {
	//anonymous function

	func() {
		fmt.Println("Anonymous")
	}()

	func(s string){
		fmt.Println("Hi I am ",s)
	}("Shobhit")
}
