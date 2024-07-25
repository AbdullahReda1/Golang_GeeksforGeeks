package main

import "fmt"

func main() {
	/* Use of map: Declares a map `fruits` to store key-value pairs of strings */

	// Empty map
	emptymap := map[int]string{}

	// Using var keyword
	var varmap map[int]int

	// Valuable map
	Valuablemap := map[int]string{
		1: "Apple",
		2: "Banana",
		3: "Cherry",
	}

	// Using make function
	var makemap = make(map[int]float32)
	makemap[0] = 0.0
	makemap[1] = 0.1
	makemap[2] = 0.2
	makemap[3] = 0.3
	makemap[4] = 0.4
	makemap[5] = 0.5

	fmt.Println(emptymap, "It's an empty map ._.")
	fmt.Println(varmap, "=", nil)
	fmt.Println(Valuablemap)
	fmt.Println(makemap)

	// Iterate over map using for range "forr", The value of this loop may vary because the map is an unordered collection
	for id, value := range makemap {
		fmt.Println(id, "\t", value)
	}

	// Adding and changing key value pairs
	makemap[6] = 0.6
	makemap[0] = 0.01

	fmt.Println(makemap)
}
