package main

import "fmt"

func main() {
	var m = make(map[string]int)
	m["ab"] = 12
	m["bc"] = 13
	fmt.Println(m)

	var m_1 = map[string]int{
		"a": 0,
		"b": 1,
		"c": 3,
		"d": 4,
	}
	fmt.Println(m_1)
}
