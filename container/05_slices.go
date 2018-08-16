package main

import "fmt"

func updateSlice(s []int) {
	s[0] = 100
}

/**
 * slice 是对数组的一个视图
 */
func main() {
	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7}

	s1 := arr[2:6]
	fmt.Printf("s1 = %d\n", s1)

	s2 := arr[2:]
	fmt.Printf("s2 = %d\n", s2)

	s3 := arr[:6]
	fmt.Printf("s3 = %d\n", s3)

	s4 := arr[:]
	fmt.Printf("s4 = %d\n", s4)

	fmt.Println("------------------")

	updateSlice(s1)
	updateSlice(s2)
	updateSlice(s3)
	updateSlice(s4)
	fmt.Printf("s1 = %d\n", s1)
	fmt.Printf("s2 = %d\n", s2)
	fmt.Printf("s3 = %d\n", s2)
	fmt.Printf("s4 = %d\n", s4)

	fmt.Println("------------------")

	fmt.Println("reslice")
	s2 = s2[2:5]
	fmt.Printf("s2 = %d\n", s2)
	s2 = s2[2:]
	fmt.Printf("s2 = %d\n", s2)

	fmt.Println("------------------")

	fmt.Println("extending slice")
	arr[0], arr[2] = 0, 2
	s1 = arr[2:5]
	fmt.Printf("s1 = %d\n", s1)

	// s[i]不可以超越len(s),向后扩展不可以超过底层数据cap(s)
	s1 = s1[2:6]
	fmt.Printf("s1 = %d\n", s1)
	fmt.Printf("s1 = %d, len(s1) = %d, cap(s1) = %d\n", s1, len(s1), cap(s1))

	fmt.Println("------------------")

	// 添加元素时如果超越cap，系统会重新分配更大的底层数组
	// 由于值传递的关系，必须接收append的返回值
	s5 := append(s1, 10)
	s6 := append(s5, 11)
	s7 := append(s6, 12)

	fmt.Printf("s5 = %d, s6= %d, s7 = %d\n", s5, s6, s7)
}
