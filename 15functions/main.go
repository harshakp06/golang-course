package main

import "fmt"

func main() {
	fmt.Println("Welcome to functions in Golang")
	greeter()
	grreter2()

	resuilt := adder(3, 9)

	fmt.Println("Result is: ", resuilt)

	proResult, mymessage := proAdder(2, 5, 6, 8, 7418, 6688)

	fmt.Println("Pro resilt is: ", proResult)
	fmt.Println("Pro message is: ", mymessage)

}

func proAdder(values ...int) (int, string) {
	total := 0

	for _, val := range values {
		total += val
	}

	return total, "Hi Pro result function"

}

func adder(valOne int, valTwo int) int {
	return valOne + valTwo

}

func grreter2() {
	fmt.Println("another method")

}

func greeter() {
	fmt.Println("namaste from Golang")

}
