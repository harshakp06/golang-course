package main

import "fmt"

func main() {
	fmt.Println("Structs in Golang")

	harsha := User{"Harsha", "harsha@go.dev", true, 25}
	fmt.Println(harsha)
	fmt.Printf("harsha details are: %+v\n", harsha) // values with structure and naming convention
	fmt.Printf("harsha details are: %v\n", harsha)  //just prints values
	fmt.Printf("Nmae is %v and email is %v\n", harsha.Name, harsha.Email)

	harsha.GetStatus()
	harsha.NewMail() // this one won't change the actial email

	// no change sin the email
	fmt.Printf("Nmae is %v and email is %v\n", harsha.Name, harsha.Email)

}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) GetStatus() {
	fmt.Println("Is User active: ", u.Status)

}

func (u User) NewMail() {
	u.Email = "test01@go1.dev"
	fmt.Println("Email of this user is: ", u.Email)

}
