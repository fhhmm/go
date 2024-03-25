package main

import "fmt"

const x int64 = 10

const (
	idKey   = "id"
	nameKey = "name"
)

const z = 20 * 10

func main() {
	const y = "hello"

	fmt.Println(x)
	fmt.Println(y)

	// error
	x = x + 1
	// error
	y = "bye"

	fmt.Println(x)
	fmt.Println(y)
}
