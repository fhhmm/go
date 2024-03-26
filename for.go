package main

import "fmt"

func main() {
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	j := 0
	for j < 5 {
		fmt.Println(j)
		j++
	}

	k := 0
	for {
		if k > 5 {
			break
		}
		fmt.Println(k)
		k++
	}

	slice := []int{1, 10, 100}
	for _, v := range slice {
		fmt.Println(v)
	}

	for _, v := range []int{2, 20, 200} {
		fmt.Println(v)
	}
}
