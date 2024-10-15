package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6}
	numbers = append(numbers, 10, 11, 12, 123, 1234, 15)
	fmt.Println("Numbers : ", numbers)
	fmt.Printf("Number has data type %T\n", numbers)
	fmt.Println("Length : ", len(numbers))
}
