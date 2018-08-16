package main

import (
	"fmt"
	"math"
	"reflect"
	"runtime"
)

/**
 *
func eval(a, b int, op string) int {
	switch op {
	case "+":
		return a + b
	case "-":
		return a - b
	case "*":
		return a * b
	case "/":
		return a / b
	default:
		panic(fmt.Sprintf("unsupported operation [%s]", op))
	}
}
*/
func eval(a, b int, op string) (int, error) {
	switch op {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		q, _ := div(a, b)
		return q, nil
	default:
		return 0, fmt.Errorf("unsupported operation [%s]", op)
	}
}

/**
 * 除法求商及余数。返回多个值
 *
func div(a, b int) (int, int) {
	return a / b, a % b
}
*/
func div(a, b int) (q, r int) {
	q = a / b
	r = a % b
	return
}

/**
 * 函数式编程
 */
func apply(op func(int, int) int, a, b int) int {
	p := reflect.ValueOf(op).Pointer()
	opName := runtime.FuncForPC(p).Name()
	fmt.Printf("calling function %s with args (%d, %d)\n", opName, a, b)
	return op(a, b)
}

func pow(a, b int) int {
	return int(math.Pow(float64(a), float64(b)))
}

/**
 * 可变参数列表
 */
func sum(numbers ...int) int {

	sum := 0
	for i := range numbers {
		sum += numbers[i]
	}
	return sum
}

/**
 * 交换两个数字
 */
func swap(a, b *int) {
	*a, *b = *b, *a
}

func main() {

	if result, err := eval(1, 3, "x"); err != nil {
		fmt.Println("Error : ", err)
	} else {
		fmt.Println(result)
	}

	fmt.Println(eval(1, 3, "+"))
	fmt.Println(eval(2, 3, "-"))
	fmt.Println(eval(3, 3, "*"))
	fmt.Println(eval(9, 3, "/"))
	fmt.Println(eval(9, 3, "="))

	q, r := div(13, 3)
	fmt.Println(q, r)

	fmt.Println("-----------------------------------")
	fmt.Println(apply(pow, 3, 4))

	fmt.Println(apply(func(a int, b int) int {
		return int(math.Pow(float64(a), float64(b)))
	}, 3, 4))

	fmt.Println("-----------------------------------")
	fmt.Println(sum(1, 2, 3, 4, 5))

	fmt.Println("-----------------------------------")

	a, b := 3, 4
	swap(&a, &b)
	fmt.Println(a, b)
}
