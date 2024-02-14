package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {

	// Reader()

	Reader1()
}

func Reader() {

	r := strings.NewReader("Learning is fun")

	buf := make([]byte, 4)

	// n, err := r.Read(buf)

	// fmt.Println(buf[:n], err, n)

	for {
		n, err := r.Read(buf)

		fmt.Println(buf[:n], err)
		fmt.Println(string(buf[:n]), err)

		if err != nil {
			fmt.Println("breaking out")
			break
		}

	}

}

func Reader1() {

	r := strings.NewReader("Some io.Reader stream to be read\n")

	if _, err := io.Copy(os.Stdout, r); err != nil {
		log.Fatal(err)
	}
}
