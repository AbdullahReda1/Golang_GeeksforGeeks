package main

import "fmt"

type point struct {
	x bool
	y bool
}

func main() {
	print("dec:      ", 42, "\n")
	fmt.Println("dec:\t ", 42)
	fmt.Printf("dec:\t  %d\n", 42)

	fmt.Printf("bin:\t  %b\n", 42)
	fmt.Printf("bin:\t  %#b\n", 42)

	fmt.Printf("hex:\t  %x\n", 42)
	fmt.Printf("hex:\t  %#x\n", 42)

	fmt.Printf("oct:\t  %o\n", 42)
	fmt.Printf("oct:\t  %#o\n", 42)

	//fmt.Printf("syn:\t  %+v\n", 42)

	p := point{x: true, y: false}
	fmt.Printf("bol:\t  %+v\n", p)
}
