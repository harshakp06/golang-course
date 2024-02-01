package main

import (
	"fmt"
)

func main() {
	fmt.Println("Welcome to loops in Golang")

	days := []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Friday", "Saturday"}

	fmt.Println(days)

	//	for d := 0; d < len(days); d++ {
	//		fmt.Println(days[d])

	//	}

	//for i := range days {
	//	fmt.Println(days[i])
	//}

	//for index, day := range days {
	//	fmt.Printf("Index is %v and value is %v\n", index, day)
	//}

	rougueValue := 1

	for rougueValue < 10 {

		if rougueValue == 2 {
			goto kp
		}

		if rougueValue == 5 {
			//break
			rougueValue++
			continue // skips value 5
		}

		if rougueValue == 3 {
			goto kp1
		}

		fmt.Println("Value is ", rougueValue)
		rougueValue++
	}

kp:
	fmt.Println("Using goto for value 2")

kp1:
	fmt.Println("Using kp1 goto for value 3")

}
