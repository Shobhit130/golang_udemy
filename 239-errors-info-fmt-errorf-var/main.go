package main

import (
	"fmt"
	"log"
)

func main() {
	_, err := sqrt(-10)

	if err != nil {
		log.Fatalln(err)
	}
}

func sqrt(f float64) (float64, error) {
	if f < 0 {
		ErrNegSqrt := fmt.Errorf("error:square root of negative number: %v", f)
		return 0, ErrNegSqrt
	}
	return 42, nil
}
