package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	fmt.Println("Welcome to json Golang")
	EncodeJson()

}

type course struct {
	Name     string `json:"coursename"`
	Price    int
	Platform string   `json:"website"`
	Password string   `json:"-"`
	Tags     []string `json:"tags,omitempty"` // nil is not printed - omitempty removes
}

func EncodeJson() {
	lcoCourses := []course{
		{"ReactJs Bootcamp", 299, "LearnCodeOnline.in", "abc12340", []string{"web-dev", "js"}},
		{"MERN Bootcamp", 299, "LearnCodeOnline.in", "abcsdn123408745", []string{"fullstack", "js"}},
		{"Angular Bootcamp", 299, "LearnCodeOnline.in", "abc12340", nil},
	}

	// pakage this data as JSON data

	finalJson, err := json.MarshalIndent(lcoCourses, " ", "\t")

	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", finalJson)
}
