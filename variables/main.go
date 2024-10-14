package main

import "fmt"

func test() {
	foo := "Hello World" // Shorthand only works inside function bodies.
	fmt.Println(foo)
}
func main() {
	var foo string
	var country string = "India"
	var c1, c2, c3 int = 1, 2, 3
	fmt.Println(foo)
	fmt.Println(foo, country, c1, c2, c3)
	test()
	const name string = "Developer"
	//name = "saikat"
	fmt.Println(name)
}
