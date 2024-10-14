package main

import "fmt"

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

type Hero struct {
	Person
	Alias string
}

type SuperHero struct {
	Person Person
	Alias  string
}

func main() {

	//Embeddings

	// h1 := Hero{
	// 	Person: Person{
	// 		FirstName: "Bruce",
	// 		LastName:  "Wayne",
	// 		Age:       45,
	// 	},
	// 	Alias: "Batman",
	// }

	// fmt.Println(h1)

	s := Hero{}

	s.FirstName = "Bruce"
	s.LastName = "Wayne"
	s.Age = 40
	s.Alias = "batman"
	fmt.Print(s)

	//Composition
	p := Person{"Bruce", "Wane", 45}
	s2 := SuperHero{p, "batman"}
	fmt.Print(s2)
}
