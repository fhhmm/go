package main

import "fmt"

func main() {
	a := 1
	switch a {
	case 1, 2:
		fmt.Println("a is 1 or 2")
	case 3:
		fmt.Println("a is 3")
	default:
		fmt.Println("default")
	}

	switch {
	case a == 1 || a == 2:
		fmt.Println("a is 1 or 2")
	case a == 3:
		fmt.Println("a is 3")
	default:
		fmt.Println("default")
	}
}
