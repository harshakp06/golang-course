package main

import "fmt"

// numb := 885333 // can't declare without var keywword

var numb = 859977456 // declaring var

const LogicToken string = "xgghhaertyu" // pubilc can be used all over file

func main() {
	fmt.Println("Variables")

	var username string = "harsha"
	fmt.Println(username)
	fmt.Printf("Variable is of type: %T \n", username)
	//fmt.Println("Variable is of type: %T", username) // using println can't print type %T can't be interpreted so Printf to be used

	var isLoggedIn bool = false
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type: %T \n", isLoggedIn)

	var smallval uint8 = 255 // uint consider upto 0-255 2^8
	fmt.Println(smallval)
	fmt.Printf("Variable is of type: %T \n", smallval)

	var smallFloat float64 = 255.5555565888
	fmt.Println(smallFloat)
	fmt.Printf("Variable is of type: %T \n", smallFloat)

	var anotherVariable int // will give 0 as output if not intialised with any value for int
	fmt.Println(anotherVariable)
	fmt.Printf("Variable is of type: %T \n", anotherVariable)

	var anotherString string // will give empty noting as output if not intialised with any value for string
	fmt.Println(anotherString)
	fmt.Printf("Variable is of type: %T \n", anotherString)

	// implicit type

	var website = "harshakp06.in" // no need to instialize with data type go lexer will taker of the intialization
	fmt.Println(website)
	fmt.Printf("Variable is of type: %T \n", website)
	// website = 5  // cannot change data type once intialised for the variable
	website = "kjsdfmn.ksldmf,,jaskdmf,=jksdmn" // chnage value of the string
	fmt.Println(website)

	// no var style // not declaration

	numberOfUser := 9745662
	fmt.Println(numberOfUser)

	fmt.Println(numb)

	fmt.Println(LogicToken)
	fmt.Printf("Variable is of type: %T \n", LogicToken)

}
