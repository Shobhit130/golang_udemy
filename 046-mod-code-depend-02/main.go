package main

import (
	"fmt"

	"github.com/Shobhit130/puppy"
)

func main() {
	puppy.From2()
	s1 := puppy.Bark()
	s2 := puppy.BigBark()
	fmt.Println(s1,s2)
}
