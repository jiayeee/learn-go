package main

import "fmt"

func printSlice(s []int) {
	fmt.Printf("v=%v, len=%d, cap=%d\n", s, len(s), cap(s))
}

func main() {

	fmt.Println("create slice")

	var s []int

	for i := 0; i < 100; i++ {
		printSlice(s)
		s = append(s, i*2+1)
	}

	fmt.Printf("%d\n", s)

	// 自建指定元素的slice
	s1 := []int{1, 2, 3, 4, 9}
	printSlice(s1)

	// 自建有16个元素的slice
	s2 := make([]int, 13)
	printSlice(s2)

	// 指定capacity
	s3 := make([]int, 13, 30)
	printSlice(s3)

	fmt.Println("------------------------------------")
	fmt.Println("copy slice")

	copy(s2, s1)
	printSlice(s2)

	fmt.Println("------------------------------------")
	fmt.Println("delete element from slice")
	s2 = append(s2[:2], s2[3:]...) // 特殊语法
	printSlice(s2)

	fmt.Println("------------------------------------")
	fmt.Println("pop from front")
	front := s2[0]
	s2 = s2[1:]
	fmt.Printf("front = %v, s2 = %v\n", front, s2)

	fmt.Println("------------------------------------")
	fmt.Println("pop from back")
	s2[len(s2)-1] = 11
	tail := s2[len(s2)-1]
	s2 = s2[:len(s2)-1]
	fmt.Printf("tail = %v, s2 = %v\n", tail, s2)

}
