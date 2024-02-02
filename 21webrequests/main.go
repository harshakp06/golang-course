package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	fmt.Println("Welcome to web verb video")
	//PerformGetRequest()
	// PerformPostRequest()

	PerformPostFormRequest()
}

func PerformGetRequest() {
	const myurl = "http://localhost:8000/"

	response, err := http.Get(myurl)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	fmt.Println("Status code: ", response.StatusCode)

	fmt.Println("Content length is: ", response.ContentLength)

	var responseString strings.Builder
	content, _ := ioutil.ReadAll(response.Body)
	byteCount, _ := responseString.Write(content)

	// fmt.Println(string(content))

	fmt.Println("Bytecount is; ", byteCount)
	fmt.Println(responseString.String())

}

func PerformPostRequest() {
	const myurl = "http://localhost:8000/post"

	//fake json payload

	requestBody := strings.NewReader(`
		{
			"coursename":"Let's go with Golang",
			"price": 0,
			"platform":"learnCodeOnline.in"
		}
	`)

	reponse, err := http.Post(myurl, "application/json", requestBody)

	if err != nil {
		panic(err)
	}
	defer reponse.Body.Close()

	content, _ := ioutil.ReadAll(reponse.Body)
	fmt.Println(string(content)) // priting form above psot method not using any get method
}

func PerformPostFormRequest() {
	const myurl = "http://localhost:8000/postform"

	//form data

	data := url.Values{}
	data.Add("firtsname", "harsha")
	data.Add("lastsname", "choudary")
	data.Add("email", "harsha@go.dev")

	response, err := http.PostForm(myurl, data)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()
	content, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(content))

}
