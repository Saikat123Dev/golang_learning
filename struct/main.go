package main

import "fmt"

type Person struct {
	FirstName string
	LastName  string
	Age       int64
}

func main() {
	var p1 = Person{
		FirstName: "Saikat",
		LastName:  "Bera",
		Age:       21,
	}

	var p4 = Person{FirstName: "Bruce", LastName: "Wayne"}

	fmt.Println("Person 4:", p4)
	fmt.Println("Person1 :", p1)

	var a = struct {
		Name string
	}{"Goru"}

	fmt.Println("Anonymous:", a)
	// two structs are equal if all their corresponding fields are equal as well.

	ptr := &p1
	fmt.Print(ptr.Age)
}
