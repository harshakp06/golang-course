package main

import "fmt"

func main() {
	// defer key word make the line to be excuted end of the lines
	// more defers are placed
	// 1st defer is last and last defer is first
	defer fmt.Println("world")
	defer fmt.Println("one")
	defer fmt.Println("two")
	defer fmt.Println("three")
	fmt.Println("hello")

	myDefer()

}

/*
OUTPUT

hello
0
1
2
3
4
5
three
two
one
world

*/

func myDefer() {
	for i := 0; i < 6; i++ {
		fmt.Println(i)

	}

}
