package main

import (
	"fmt"
	"log"
	"os"

	"github.com/sirupsen/logrus"
)

func main() {
	// log1()
	log2()

}

func log2() {
	// file, err := os.OpenFile("app.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0777)

	// if err != nil {
	// 	fmt.Println(err)

	// 	return
	// }
	// logrus.SetOutput(file)
	logrus.SetLevel(logrus.DebugLevel)
	logrus.Debug("Hello world")
	logrus.Println("Hello World")

	logrus.SetLevel(logrus.TraceLevel)
	logrus.Trace("Hellojkdzxcn")
	// file.Close()

}

func log1() {
	file, err := os.OpenFile("app.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0777)

	if err != nil {
		fmt.Println(err)

		return
	}
	log.SetOutput(file)
	log.Println("Hello World")

	file.Close()

}
