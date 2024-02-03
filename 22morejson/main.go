package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	fmt.Println("Welcome to json Golang")
	// EncodeJson()
	DecodeJson()
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

func DecodeJson() {
	jsonDataFromWeb := []byte(`
	{
		"coursename": "ReactJs Bootcamp",
		"Price": 299,
		"website": "LearnCodeOnline.in",
		"tags": [
				"web-dev","js"]
	}
	`)

	var lcoCourse course

	checkValid := json.Valid(jsonDataFromWeb)

	if checkValid {
		fmt.Println("JSON was valid")
		json.Unmarshal(jsonDataFromWeb, &lcoCourse)
		fmt.Printf("%#v\n", lcoCourse)
	} else {
		fmt.Println("JSON was not valid")
	}

	// some cases where you just want to add data to key value

	var myOnlienData map[string]interface{}
	json.Unmarshal(jsonDataFromWeb, &myOnlienData)
	fmt.Printf("%#v\n", myOnlienData)

	for k, v := range myOnlienData {
		fmt.Printf("\nKey is %v and value is %v and Type is: %T\n", k, v, v)

	}

}
