package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	//file()
	//file2()
	// file3()
	// file4()
	file5()

}

func file5() {
	file, err := os.OpenFile("temp2.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0600)

	if err != nil {
		fmt.Println(err)

	}

	defer file.Close()

	_, err = file.WriteString("Have a great day\n")
	if err != nil {
		fmt.Println(err)
	}

}

func file4() {
	path3 := "temp.txt"

	file, err := os.Open(path3)

	if err != nil {
		fmt.Println("Error: ", err)

	}

	b := make([]byte, 4)

	for {
		n, err := file.Read(b)
		//data, err := f.Read(b)

		if err != nil {
			fmt.Println("Error: ", err)
			break

		}

		fmt.Println(string(b[0:n]))
	}

}

func file2() {

	fileinfo, err := os.Stat("32files")

	if err != nil {
		fmt.Println(err)

	}

	fmt.Println(fileinfo.Name())
	fmt.Println(fileinfo.Size())
	fmt.Println(fileinfo.Mode())
	fmt.Println(fileinfo.IsDir())

}

func file3() {
	path3 := "temp.txt"

	data, err := os.ReadFile(path3)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(data)

	fmt.Println(string(data))

}

func file() {
	path1 := filepath.Join("dir1", "dir2", "dir/../dir3", "text.txt")

	fmt.Println(path1)
	fmt.Println(filepath.Dir(path1))
	fmt.Println(filepath.Base(path1))

	fmt.Println(filepath.IsAbs(path1))       // to check absolute path
	fmt.Println(filepath.IsAbs("/dir/file")) // to check absolute path

	fmt.Println(filepath.Ext(path1)) // give extension of the file

}
