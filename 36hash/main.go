package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

func main() {

	var password string
	fmt.Scanln(&password)
	fmt.Println("password encrypted to :", encod(password))
}

func encod(str string) string {
	var hashCode = md5.Sum([]byte(str))
	return hex.EncodeToString(hashCode[:])
}
