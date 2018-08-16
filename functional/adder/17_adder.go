package main

import "fmt"

func adder() func(v int) int {
	sum := 0
	return func(v int) int {
		sum += v
		return sum
	}
}

func main() {

	a := adder()
	for i := 0; i < 100; i++ {
		fmt.Printf("%d, [%d]\n", i, a(i))
	}

}
