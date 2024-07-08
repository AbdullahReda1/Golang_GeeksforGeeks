// `package main` tells the compiler that package must compile in the executable program rather than a shared library.
// It is the starting point of the program and also contains the main function in it.
package main

// `import` keyword is used to import packages in program and `fmt` package is used to implement formatted Input/Output with functions.
import (
	"fmt"
)

// `func` It is used to create a function in Go language.
// `main` It is the main function in Go language.
func main() {
	// `Print` This method is present in fmt package and it is used for displaying without new line.
	fmt.Print("\n")
	// `Println` This method is present in fmt package and it is used for displaying with new line.
	fmt.Println("!...Hello world...!")
	fmt.Print("\n")
}
