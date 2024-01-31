package main

import (
	"fmt"
	"math/big"

	//"math/rand"
	"crypto/rand"
	// "time"
)

func main() {
	fmt.Println("Welcome to maths in goalng")

	// random number using rand package
	//rand.Seed(time.Now().UnixNano()) // rand.seed is deprecated
	//fmt.Println(rand.Intn(10) + 2)

	// random number using crypto package

	myRandomNum, _ := rand.Int(rand.Reader, big.NewInt(10))
	fmt.Println(myRandomNum)

}
