package main

import "fmt"

func main() {
	fmt.Println("Welcome to class on pointers")

	// var ptr *int
	// fmt.Println("Value of pointer is ", ptr)

	myNumber := 24
	var ptr = &myNumber
	fmt.Println("Value of Actual Pointer is ", ptr)
	fmt.Println("Value of Actual Pointer is ", *ptr)

	*ptr = *ptr + 2
	fmt.Println("New value is: ", myNumber)
}
