package main

import (
	"fmt"
	"sort"
)

func main() {

	fmt.Println("welcome to slices of array")

	var fruitList = []string{"Apple", "Tomato", "Peach"}
	fmt.Printf("Type of fruitlist is %T \n", fruitList)

	fruitList = append(fruitList, "Mango", "Banana")
	fmt.Println(fruitList)
	fmt.Println(len(fruitList))

	fruitList1 := append(fruitList[1:3])
	fmt.Println(fruitList)
	fmt.Println(len(fruitList1))

	fmt.Println(fruitList1)

	fruitList1 = append(fruitList1[:3])
	fmt.Println(fruitList1)
	fmt.Println(len(fruitList1))

	highScores := make([]int, 4)

	highScores[0] = 456
	highScores[1] = 874
	highScores[2] = 8746
	highScores[3] = 741
	//	highScores[4] = 7413

	/* using append method we can add the more data
	it will automatically increase the range of the slice*/

	highScores = append(highScores, 555, 666, 874966, 5522)

	fmt.Println(highScores)

	// sorting the data using sort packages

	fmt.Println(sort.IntsAreSorted(highScores))
	sort.Ints(highScores)
	fmt.Println(highScores)

	// how to remove a value from slices

	var courses = []string{"reactjs", "javascript", "swift", "python", "ruby"}
	fmt.Println(courses)
	var index int = 2
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses)

}
