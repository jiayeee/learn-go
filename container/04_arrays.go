package main

import "fmt"

func printArray(arr [5]int) {
	for i, v := range arr {
		fmt.Println(i, v)
	}
}

/**
 * [10]int [5]int 是不同的类型
 */
func main() {

	var arr1 [5]int
	arr2 := [3]int{1, 3, 4}
	arr3 := [...]int{1, 3, 4, 5, 6}
	var grid [3][4]int

	fmt.Println(arr1, arr2, arr3)

	fmt.Println(grid)

	for i := 0; i < len(arr3); i++ {
		fmt.Println(arr3[i])
	}

	fmt.Println("----------------")

	printArray(arr1)
	//printArray(arr2)
	printArray(arr3)
}
