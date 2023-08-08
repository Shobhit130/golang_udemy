package main

import "fmt"

func main(){
	fmt.Println("Hello World")
	a := 42
	fmt.Printf("42 as binary %b\n",a)
	fmt.Printf("42 as hexadecimal %x\n",a)

	b,c,d,e,f := 1,2,3,4,5

	fmt.Printf("%v \t %b \t %x\n",b,b,b)
	fmt.Printf("%v \t %b \t %x\n",c,c,c)
	fmt.Printf("%v \t %b \t %x\n",d,d,d)
	fmt.Printf("%v \t %b \t %x\n",e,e,e)
	fmt.Printf("%v \t %b \t %x\n",f,f,f)
}