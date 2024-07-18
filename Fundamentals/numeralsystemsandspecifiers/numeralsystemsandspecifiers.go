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

	num := 12345.6789
	smallNum := 0.00012345
	bigNum := 123456789.0

	fmt.Printf("flt:\t  %f\n", num)
	fmt.Printf("flt:\t  %F\n", num)
	fmt.Printf("flt:\t  %g\n", num)
	fmt.Printf("flt:\t  %G\n", num)
	fmt.Printf("flt:\t  %f\n", smallNum)
	fmt.Printf("flt:\t  %F\n", smallNum)
	fmt.Printf("flt:\t  %g\n", smallNum)
	fmt.Printf("flt:\t  %G\n", smallNum)
	fmt.Printf("flt:\t  %f\n", bigNum)
	fmt.Printf("flt:\t  %F\n", bigNum)
	fmt.Printf("flt:\t  %g\n", bigNum)
	fmt.Printf("flt:\t  %G\n", bigNum)

	fmt.Printf("bin:\t  %b\n", 42)
	fmt.Printf("bin:\t  %#b\n", 42)

	fmt.Printf("hex:\t  %x\n", 42)
	fmt.Printf("hex:\t  %#x\n", 42)
	fmt.Printf("hex:\t  %X\n", 42)
	fmt.Printf("hex:\t  %#X\n", 42)

	fmt.Printf("oct:\t  %o\n", 42)
	fmt.Printf("oct:\t  %#o\n", 42)

	str := "世界"
	fmt.Printf("chr:\t  %c\n", 'A')
	fmt.Printf("str:\t  %s\n", str)

	p := point{x: true, y: false}
	fmt.Printf("bol:\t  %+v\n", p)
	fmt.Printf("def:\t  %v\n", p)
	fmt.Printf("syn:\t  %#v\n", p)

	fmt.Printf("typ:\t  %T\n", p)
	fmt.Printf("typ:\t  %T\n", 42)
	fmt.Printf("typ:\t  %T\n", 4.2)

	var r rune = 'a'
	fmt.Printf("uni:\t  %U\n", r)
	fmt.Printf("uni:\t  %U\n", str[0])
	fmt.Printf("uni:\t  %U\n", str[1])
}
