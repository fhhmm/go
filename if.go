package main

import "fmt"

func if1(flag bool) bool {
	if flag {
		fmt.Println("flag is true")
	} else {
		fmt.Println("flag is false")
	}

	return flag
}

func main() {
	if1(true)

	if a := false; a {
		fmt.Println("a is true")
	} else {
		fmt.Println("a is false")
	}
}
