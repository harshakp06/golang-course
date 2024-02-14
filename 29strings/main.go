package main

import (
	"fmt"
	"strings"
)

func main() {

	// text := "Go is fun"

	// fmt.Println(text)

	learning := "learning standard library in Go"
	// fun := "library in Go"
	// fun := "learning in Go"
	fun := "learning"

	result := strings.Contains(learning, fun) // here comapring 2 strings if fun is present in learning
	fmt.Println(result)                       // will give boolean value after comapring both strings

	result2 := strings.Count(learning, fun) // will give number of times fun string present in learning string
	fmt.Println(result2)

	result3 := strings.ReplaceAll(learning, fun, "library in GO") //replacing fun in learning string with the given string
	fmt.Println(result3)

}
