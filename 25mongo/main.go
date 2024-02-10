package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/harshakp06/mongoapi/router"
)

func main() {
	fmt.Println("Mongo API")
	r := router.Router()
	fmt.Println("Server is grtting started ...")
	log.Fatal(http.ListenAndServe(":4000", r))
	fmt.Println("listening at port 4000 ...")

}
