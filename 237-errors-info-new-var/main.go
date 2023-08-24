package main

import (
	"errors"
	"fmt"
	"log"
)

var ErrNegSqrt = errors.New("error: square root of negative number")

func main() {
	fmt.Printf("%T\n", ErrNegSqrt)
	_, err := sqrt(-10)

	if err != nil {
		log.Fatalln(err)
	}
}

func sqrt(f float64) (float64, error) {
	if f < 0 {
		return 0, ErrNegSqrt
	}
	return 42, nil
}
