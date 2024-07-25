package main

import (
	"fmt"
	"time"
)

func main() {
	/* Use of defer to ensure a function call is done at the end of the surrounding function */
	defer fmt.Println("This message is deferred and will be printed last")

	/* Use of var to declare a variable: Declares a variable msg and assigns it a string value */
	var msg string = "Hello, Go!"
	fmt.Println(msg)
	print("\n")

	/* Use of type to define a new data type */
	type My_int int8
	var x My_int = 5
	fmt.Printf("This x = %d value with new data type 'My_int'\n", x)
	print("\n")

	/* Use of select to wait on multiple channel operations */
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		ch1 <- "Message from channel 1"
	}()
	go func() {
		time.Sleep(2 * time.Second)
		ch2 <- "Message from channel 2"
	}()

	select {
	case msg1 := <-ch1:
		fmt.Println("Received: ", msg1)
	case msg2 := <-ch2:
		fmt.Println("Received: ", msg2)
	case <-time.After(3 * time.Second):
		fmt.Println("Timeout: No Message Received")
	}
	print("\n")

	/* Use of fallthrough in a switch statement */
	num := 1
	switch num {
	case 1:
		fmt.Println("one")
		fallthrough
	case 2:
		fmt.Println("two")
		fallthrough
	case 3:
		fmt.Println("three")
	}
	print("\n")
}
