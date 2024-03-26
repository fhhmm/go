package main

import "fmt"

func main() {
	type person struct {
		name string
		age  int
		pet  string
	}
	julia := person{name: "julia", age: 10}
	julia.pet = "dog"
	fmt.Println(julia.name, julia.age, julia.pet)

	p := struct {
		name string
		age  int
	}{name: "Gopher", age: 10}
	fmt.Println(p.name, p.age)
}
