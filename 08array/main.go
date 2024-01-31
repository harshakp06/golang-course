package main

import "fmt"

func main() {

	fmt.Println("Welcome to file of array")

	var fruitList [6]string

	fruitList[0] = "Apple"
	fruitList[1] = "Yellow"
	fruitList[4] = "Orange"
	fruitList[5] = "Blue"

	fmt.Println("Fruit list is: ", fruitList)

	fmt.Println("Fruit list is: ", len(fruitList))

	var vegList = [4]string{"potato", "carrot", "mushroom"}

	fmt.Println("veg list is: ", vegList)
	fmt.Println("veg list len: ", len(vegList))

}
