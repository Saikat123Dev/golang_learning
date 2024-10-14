package main

import "fmt"

func main() {
	x := 7
	if x > 10 {
		fmt.Printf("x is gt than 10")
	} else if x < 10 && x > 5 {
		fmt.Printf("x is ls than 10 but gt than 5")
	} else {
		fmt.Printf("something else")
	}
}
