package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First string `json:"first"`
	Last  string `json:"last"`
	Age   int    `json:"age"`
}

func main() {
	s := `[{"First":"Shobhit","Last":"Agarwal","Age":22},{"First":"Eshaan","Last":"Mohapatra","Age":22}]`
	bs := []byte(s)

	fmt.Printf("%T\n", s)
	fmt.Printf("%T\n", bs)

	var people []person

	err := json.Unmarshal(bs, &people)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("All people data", people)

	for i, v := range people {
		fmt.Println("Person Number", i)
		fmt.Println(v.Age, v.First, v.Last)
	}
}
