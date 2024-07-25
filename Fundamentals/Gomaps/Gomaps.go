package main

import "fmt"

func main() {
	// Empty map
	emptymap := map[int]string{}

	// Using var keyword
	var varmap map[int]int

	// Valuable map
	Valuablemap := map[int]string{
		1: "first",
		2: "second",
		3: "third",
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
	print("\n")
	fmt.Println(varmap, "=", nil)
	print("\n")
	fmt.Println(Valuablemap)
	print("\n")
	fmt.Println(makemap)
	print("\n")

	// Iterate over map using for range "forr", The value of this loop may vary because the map is an unordered collection
	for id, value := range makemap {
		fmt.Println(id, "\t", value)
	}
	print("\n")

	// Adding and changing key value pairs
	makemap[6] = 0.6
	makemap[0] = 0.01

	fmt.Println(makemap)
	print("\n")

	// Retrieve a value related to a key in the maps
	fmt.Println("Retrieved value related to a key\t", makemap[0])
	print("\n")

	// Check the key is available or not
	somevalue0, boolflag0 := makemap[6]
	somevalue1, boolflag1 := makemap[7]

	fmt.Println(boolflag0, "\t", somevalue0)
	fmt.Println(boolflag1, "\t", somevalue1)
	print("\n")

	// Check the key is available or not, Using blank identifier
	_, tboolflag := makemap[1]
	_, fboolflag := makemap[10]

	fmt.Println(tboolflag)
	fmt.Println(fboolflag)
	print("\n")

	// Delete key from the map
	delete(makemap, 0)
	fmt.Println(makemap)
	print("\n")

	// Modifying map
	themap := map[string]float32{"true": 1, "false": 0}
	fmt.Println(themap)
	modmap := themap
	fmt.Println(modmap)
	modmap["fuzzy"] = 0.5
	fmt.Println(modmap)
	fmt.Println(themap)
}
