package main

import "fmt"

func main() {
	var name string = "Saikat" // a string is a sequence of bytes
	var name_1 string = `Saikat 
	Bera`
	//---> var uiptr uintptr // Integer representation of a memory address
	fmt.Println(name, name_1)
	var i int
	var f float64
	var b bool
	var s string

	fmt.Printf("%v %v %v %q\n", i, f, b, s)
	percentage := (7.0 / 9) * 100
	fmt.Printf("%f %%", percentage)
}
