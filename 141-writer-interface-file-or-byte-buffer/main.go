package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
)

type person struct {
	first string
}

func (p person) writeOut(w io.Writer) error {
	_, err := w.Write([]byte(p.first))
	return err
}

func main() {

	// An `io.Writer` is an interface that wraps the basic `Write` method. Both `os.File` and
	// `bytes.Buffer` implement `io.Writer` interface which allows us to write to either a file or a byte
	// buffer in a similar way.
	
	p := person{
		first: "Shobhit",
	}
	f, err := os.Create("output.txt")
	if err != nil {
		log.Fatalf("error %s", err)
	}
	defer f.Close()

	var b bytes.Buffer

	p.writeOut(f)
	p.writeOut(&b)

	fmt.Println(b.String())
}
