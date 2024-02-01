package main

import (
	"fmt"
)

func main() {

	fmt.Println("Welcome to Golang Maps")

	languages := make(map[string]string)

	languages["JS"] = "Javascript"
	languages["PY"] = "Python"
	languages["MD"] = "Markdown"
	languages["TF"] = "Terraform"
	languages["55"] = "lang"

	fmt.Println("List of all values: ", languages)
	fmt.Println("JS shorts for: ", languages["JS"])

	delete(languages, "55")
	fmt.Println("List of all languages: ", languages)

	// golang loops

	for key, value := range languages {
		fmt.Printf("For key %v, value is %v\n", key, value)
	}

	// without key value
	for _, value := range languages {
		fmt.Printf("For key v, value is %v \n", value)
	}
}
