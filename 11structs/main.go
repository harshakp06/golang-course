package main

import "fmt"

func main() {
	fmt.Println("Structs in Golang")

	harsha := User{"Harsha", "harsha@go.dev", true, 25}
	fmt.Println(harsha)
	fmt.Printf("harsha details are: %+v\n", harsha) // values with structure and naming convention
	fmt.Printf("harsha details are: %v\n", harsha)  //just prints values
	fmt.Printf("Nmae is %v and email is %v\n", harsha.Name, harsha.Email)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
