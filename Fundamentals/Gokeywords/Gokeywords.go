package main

import "fmt"

func main() {
	// Use of var to declare a variable: Declares a variable msg and assigns it a string value.
	var msg string = "Hello, Go!"
	fmt.Println(msg)

	// Use of map: Declares a map `fruits` to store key-value pairs of strings.
	fruits := map[string]string{
		"a": "Apple",
		"b": "Banana",
		"c": "Cherry",
	}
	fmt.Println(fruits)
}

/*
	chan
	defer
	fallthrough
	go
	interface
	map
	range
	select
	type
*/
