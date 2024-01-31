package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Welcome to user input rating")

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the rating for Samosa:")

	// comma ok || err err

	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for Rating, ", input)
	fmt.Printf("Type of this rating is %T \n", input)

}
