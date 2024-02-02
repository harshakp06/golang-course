package main

import (
	"fmt"
	"net/url"
)

const myurl string = "https://lco.dev:3000/learn?coursename=reactjs&paymentd=kjsdfnksdn"

func main() {
	fmt.Println("Welcome to handling URLs in Golang")

	fmt.Println(myurl)

	//parsing
	result, err := url.Parse(myurl)

	if err != nil {
		panic(err)
	}

	// fmt.Println(result.Scheme)
	// fmt.Println(result.Host)
	// fmt.Println(result.Path)
	// fmt.Println(result.Port())
	fmt.Println(result.RawQuery)

	qprams := result.Query()
	fmt.Printf("The type of query parms are: %T \n", qprams)

	fmt.Println(qprams["coursename"])
	fmt.Println(qprams["paymentd"])

	for _, val := range qprams {
		fmt.Println("param is: ", val)
	}

	for id, val := range qprams {
		fmt.Println("param is: ", val)
		fmt.Println("id is: ", id)

	}

	partsOfUrl := &url.URL{
		Scheme:  "https",
		Host:    "lco.dev",
		Path:    "/tutcss",
		RawPath: "user=hitesh",
	}

	anotherURL := partsOfUrl.String()
	fmt.Println(anotherURL)

}
