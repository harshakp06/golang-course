package main

import (
	"fmt"
	"sort"
)

func main() {
	vars1()
}

func vars1() {

	vars := []int{5, 2, 9, 799, 7, 5, 8, 12}
	sort.Ints(vars)
	fmt.Println(vars)

	vars1 := []string{"Learning", "golang", "on", "Kodekloud", "hello"}
	sort.Strings(vars1)
	fmt.Println(vars1)

}
