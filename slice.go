package main

import (
	"fmt"
	"slices"
)

func main() {
	var x = []int{10, 20, 30, 40, 50}
	fmt.Println(x)

	// delete x[1], x[2] (delete 20, 30)
	x = append(x[:1], x[3:]...)
	fmt.Println(x)

	var y = []int{10, 20, 30, 40, 50}
	// delete x[1], x[2] (delete 20, 30)
	y = slices.Delete(y, 1, 3)
	fmt.Println(y)

	var z = []int{1, 2, 3, 4, 5}
	x = slices.Insert(x, 1, z[1:3]...)
	fmt.Println(x)
}
