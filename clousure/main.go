// This is an example of a closure. A closure is a function that "captures" variables from its surrounding environment (in this case, sum). Even though the outer function myFunction finishes execution, the returned inner function still has access to the sum variable.

package main

import "fmt"

func myFunction() func(int) int {
	sum := 0

	return func(v int) int {
		sum += v
		return sum
	}
}

func main() {
	f := myFunction()
	fmt.Println(f(10))
	fmt.Println(f(5)) // Output: 15  -> sum becomes 10 + 5 = 15
	fmt.Println(f(7)) // Output: 22  -> sum becomes 15 + 7 = 22
}
