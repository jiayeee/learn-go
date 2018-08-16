package main

import (
	"fmt"
	"learn-go/queue"
)

func main() {

	var q = queue.Queue{12}

	fmt.Println(q.IsEmpty())

	q.Push(10)
	q.Push(13)

	fmt.Printf("%v\n", q)

	fmt.Println(q.Pop())
	fmt.Println(q.Pop())
	fmt.Println(q.Pop())

	fmt.Println(q.IsEmpty())

	q.Push("abc")
	fmt.Println(q.Pop())
}
