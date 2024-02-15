package main

import (
	"fmt"
)

func main() {

	//err := errors.New("custom error occured")
	//fmt.Println(err)

	err := process(3)
	// fmt.Println(err)
	checkError(err)

	err = process(4)
	checkError(err)

}

func process(i int) error {
	if i%2 == 0 {
		// return errors.New("only odd numbers allowed")
		return fmt.Errorf("only odd numbers allowed, got: %d", i)
	}

	return nil

}

func checkError(e error) {
	if e != nil {
		fmt.Println(e)
	}

	fmt.Println("operation sucessful")

}
