package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Welcome to files in Golang")
	content := "This needs to go in file - Hello Okhjnmb m jm"

	file, err := os.Create("./filenow")
	fmt.Println(file.Name())
	//if err != nil {
	//	panic(err)
	//}
	checkNilErr(err)

	length, err := io.WriteString(file, content)
	checkNilErr(err)

	fmt.Println("length is: ", length)

	defer file.Close()

	readFile(file.Name())

}

func readFile(filename string) {
	datatype, err := ioutil.ReadFile(filename)
	checkNilErr(err)

	fmt.Println("Text data inside the file is \n", datatype)         // here datatype will give the raw data in byte
	fmt.Println("Text data inside the file is \n", string(datatype)) // here it is converted tonormal readable

}

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
