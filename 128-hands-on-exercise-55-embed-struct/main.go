package main

import "fmt"

type engine struct {
	electric bool
}

type vehicle struct {
	engine
	make  string
	model string
	doors int
	color string
}

func main() {
	v1 := vehicle{
		engine: engine{
			electric: true,
		},
		make:  "Ford",
		model: "Mustang",
		doors: 4,
		color: "blue",
	}
	v2 := vehicle{
		engine: engine{
			electric: false,
		},
		make:  "Toyota",
		model: "Tundra",
		doors: 4,
		color: "black",
	}

	fmt.Println(v1)
	fmt.Println(v2)

	fmt.Println(v1.color,v1.electric,v1.engine,v1.make,v1.doors,v1.model)

}
