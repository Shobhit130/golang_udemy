package main

import "fmt"

func main() {
	mp := map[string]int{
		"A": 42,
		"B": 16,
		"C": 12,
	}

	fmt.Println("The age os A is ", mp["A"])

	fmt.Println(mp)
	fmt.Println(len(mp))
	fmt.Printf("%#v\n", mp)

	mp1 := make(map[string]int)

	mp1["D"] = 23
	fmt.Println(mp1)

	fmt.Printf("%#v\n", mp1)
	fmt.Println(len(mp1))

	for key, value := range mp {
		fmt.Println(key, value)
	}
	for key := range mp {
		fmt.Println(key)
	}
	for _, value := range mp {
		fmt.Println(value)
	}
	
	delete(mp,"A")
	fmt.Println(mp)

	//accessing keys that don't exist
	delete(mp,"A")
	fmt.Println(mp["A"]) //gives 0

	//comma ok idiom
	value,ok := mp["A"]

	if(ok){
		fmt.Println(value)
	}else{
		fmt.Println("Not exists")
	}
}
