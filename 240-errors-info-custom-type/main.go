package main

import (
	"fmt"
	"log"
)

type ErrNegSqrt struct {
	lat string
	lon string
	err error
}

func (n ErrNegSqrt) Error() string {
	return fmt.Sprintf("a ErrNegSqrt error occured: %v %v %v", n.lat, n.lon, n.err)
}

func main() {
	_, err := sqrt(-10)

	if err != nil {
		log.Fatalln(err)
	}
}

func sqrt(f float64) (float64, error) {
	if f < 0 {
		nme := fmt.Errorf("error:square root of negative number: %v", f)
		return 0, ErrNegSqrt{"50,28898 N", "90.8374 W", nme}
	}
	return 42, nil
}
