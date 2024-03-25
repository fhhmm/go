package main

import "fmt"

func main() {
	var x = [3]int{10, 20, 30}
	var y = [12]int{1, 5: 4, 6, 10: 100, 15}
	var _x = [...]int{10, 20, 30}

	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(_x)

	fmt.Println(x == _x)
}
