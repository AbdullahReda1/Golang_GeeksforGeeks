package main

import "fmt"

type point struct {
	x bool
	y bool
}

func main() {
	print("dec:              ", 42, "\n") // Printing a decimal number using print
	fmt.Println("dec:\t         ", 42)    // Printing a decimal number using fmt.Println
	fmt.Printf("dec(%%d):\t  %d\n", 42)   // Printing a decimal number using fmt.Printf with %d

	num := 12345.6789
	smallNum := 0.00012345
	bigNum := 123456789.0

	// Printing floating-point numbers using different verbs
	fmt.Printf("flt(%%f):\t  %f\n", num)      // Using %f
	fmt.Printf("flt(%%F):\t  %F\n", num)      // Using %F
	fmt.Printf("flt(%%g):\t  %g\n", num)      // Using %g
	fmt.Printf("flt(%%G):\t  %G\n", num)      // Using %G
	fmt.Printf("flt(%%f):\t  %f\n", smallNum) // Using %f
	fmt.Printf("flt(%%F):\t  %F\n", smallNum) // Using %F
	fmt.Printf("flt(%%g):\t  %g\n", smallNum) // Using %g
	fmt.Printf("flt(%%G):\t  %G\n", smallNum) // Using %G
	fmt.Printf("flt(%%f):\t  %f\n", bigNum)   // Using %f
	fmt.Printf("flt(%%F):\t  %F\n", bigNum)   // Using %F
	fmt.Printf("flt(%%g):\t  %g\n", bigNum)   // Using %g
	fmt.Printf("flt(%%G):\t  %G\n", bigNum)   // Using %G

	// Printing integer in binary format
	fmt.Printf("bin(%%b):\t  %b\n", 42)   // Using %b
	fmt.Printf("bin(%%#b):\t  %#b\n", 42) // Using %#b

	// Printing integer in hexadecimal format
	fmt.Printf("hex(%%x):\t  %x\n", 42)   // Using %x
	fmt.Printf("hex(%%#x):\t  %#x\n", 42) // Using %#x
	fmt.Printf("hex(%%X):\t  %X\n", 42)   // Using %X
	fmt.Printf("hex(%%#X):\t  %#X\n", 42) // Using %#X

	// Printing integer in octal format
	fmt.Printf("oct(%%o):\t  %o\n", 42)   // Using %o
	fmt.Printf("oct(%%#o):\t  %#o\n", 42) // Using %#o

	str := "世界"
	// Printing character and string
	fmt.Printf("chr(%%c):\t  %c\n", 'A') // Using %c
	fmt.Printf("str(%%s):\t  %s\n", str) // Using %s

	// Printing floating-point number in scientific notation
	fmt.Printf("sci(%%e):\t  %e\n", 4.2) // Using %e
	fmt.Printf("sci(%%E):\t  %E\n", 4.2) // Using %E

	p := point{x: true, y: false}
	// Printing struct
	fmt.Printf("bol(%%+v):\t  %+v\n", p) // Using %+v
	fmt.Printf("def(%%v):\t  %v\n", p)   // Using %v
	fmt.Printf("syn(%%#v):\t  %#v\n", p) // Using %#v

	// Printing type
	fmt.Printf("typ(%%T):\t  %T\n", p)   // Using %T for struct
	fmt.Printf("typ(%%T):\t  %T\n", 42)  // Using %T for integer
	fmt.Printf("typ(%%T):\t  %T\n", 4.2) // Using %T for float

	var r rune = 'a'
	// Printing Unicode code point
	fmt.Printf("uni(%%U):\t  %U\n", r)      // Using %U for rune
	fmt.Printf("uni(%%U):\t  %U\n", str[0]) // Using %U for first byte of string
	fmt.Printf("uni(%%U):\t  %U\n", str[1]) // Using %U for second byte of string

	// Printing pointer
	fmt.Printf("ptr(%%p):\t  %p\n", &num)  // Using %p
	fmt.Printf("ptr(%%p):\t  %#p\n", &num) // Using %#p
}
