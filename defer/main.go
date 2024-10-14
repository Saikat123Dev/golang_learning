package main

import "fmt"

// As we can see, defer statements are stacked and executed in a last in first out manner.

// So, Defer is incredibly useful and is commonly used for doing cleanup or error handling.

// Functions can also be used with generics but we will discuss them later in the course.

func main() {
	defer fmt.Println("I am finished")
	defer fmt.Println("Are you?")

	fmt.Println("Doing some work...")
}
