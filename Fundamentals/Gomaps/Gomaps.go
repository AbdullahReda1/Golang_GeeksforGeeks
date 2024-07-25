package main

import "fmt"

func main() {
	/* Use of map: Declares a map `fruits` to store key-value pairs of strings */

	// Empty map
	emptymap := map[int]string{}

	// Using var keyword
	var varmapresult bool
	var varmap map[int]int
	if varmap == nil {
		varmapresult = true
	} else {
		varmapresult = false
	}

	// Valuable map
	Valuablemap := map[string]string{
		"a": "Apple",
		"b": "Banana",
		"c": "Cherry",
	}

	fmt.Println(emptymap)
	fmt.Println(varmap, " = ", varmapresult, "its nil")
	fmt.Println(Valuablemap)
}
