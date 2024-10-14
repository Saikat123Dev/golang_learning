package main

import "fmt"

func main() {
	var arr [4]int
	fmt.Println(arr)
	var arr_1 = [4]int{1, 2, 3, 4}
	fmt.Println(arr_1)
	arr_2 := [4]int{8, 7, 6}
	fmt.Println(arr_2)

	// for i := 0; i < len(arr_1); i++ {
	// 	fmt.Printf("Index: %d , Element: %d\n", i, arr_1[i])
	// }

	// for i, e := range arr_1 {
	// 	fmt.Printf("Index %d,Element %d\n", i, e)
	// }

	// for e := range arr_1 {
	// 	fmt.Printf("Index %d\n", e)
	// }

	for _, e := range arr_1 {
		fmt.Printf(" %d\n", e)
	}
}
