package main

import "fmt"

func main() {
	fmt.Println("Welcome to file on pinters")

	var ptr *int

	fmt.Println("value of pointer is ", ptr) // intializing pointer without value won't store anything

	myNumber := 27

	var ptr1 = &myNumber

	fmt.Println("Value of actual pointer is ", ptr1)
	fmt.Println("Value of actual pointer is ", *ptr1)

	*ptr1 = *ptr1 + 5 // here value of *ptr1 is actual value of mynumber, so the value of the mynumber changes

	fmt.Println("Value of mynumber is ", myNumber)
	fmt.Println("Value of *ptr1 is ", *ptr1)

}
