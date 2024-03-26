package main

import "fmt"

func main() {
	var nilmap map[string]int
	fmt.Println(len(nilmap))
	fmt.Println(nilmap["abc"])
	// ↓はエラーになる
	// nilmap["abc"] = 1

	totalWins := map[string]int{}
	totalWins["abc"] = 1
	fmt.Println(totalWins["abc"])

	teams := map[string][]string{
		"a": []string{"a", "aa", "aaa"},
		"b": []string{"b", "bb", "bbb"},
		"c": []string{"c", "cc", "ccc"},
	}
	fmt.Println(teams["a"][2])

	// loop
	for k, v := range teams {
		fmt.Println(k, v)
	}
	// only key
	for k := range teams {
		fmt.Println(k)
	}
	// only value
	for _, v := range teams {
		fmt.Println(v)
	}

	// comma ok idiom
	m := map[string]int{
		"hello": 5,
		"world": 0,
	}
	v, ok := m["hello"]
	fmt.Println(v, ok)

	v, ok = m["world"]
	fmt.Println(v, ok)

	v, ok = m["goodbye"]
	fmt.Println(v, ok)

	// delete
	delete(m, "hello")
	v, ok = m["hello"]
	fmt.Println(v, ok)
}
