package main

import (
	"fmt"
	"time"
)

func worker(i int, c chan int) {

	// 判断 channel 是否关闭的两种方法
	for n := range c {
		//n, ok := <-c
		//if !ok {
		//	break
		//}

		fmt.Printf("createWorker %d received %c\n", i, n)
	}
}

func createWorker(i int) chan<- int {
	c := make(chan int)
	go worker(i, c)
	return c
}

/**
 * chan<- 表示只能往 channel 发数据
 */
func channelDemo() {

	var channels [10]chan<- int
	for i := 0; i < 10; i++ {
		channels[i] = createWorker(i)
	}

	for i := 0; i < 10; i++ {
		channels[i] <- 'a' + i
	}

	for i := 0; i < 10; i++ {
		channels[i] <- 'A' + i
	}
	//var c chan  int
	//c := make(chan int)
	//go createWorker(0, c)

	//c <- 1
	//c <- 2

	time.Sleep(time.Second)
}

func bufferDemo() {

	c := make(chan int, 3)
	go worker(1, c)

	c <- 'a'
	c <- 'b'
	c <- 'c'
	c <- 'd'

	time.Sleep(time.Second)
}

func channelClose() {

	c := make(chan int, 3)
	go worker(1, c)

	c <- 'a'
	c <- 'b'
	c <- 'c'
	c <- 'd'

	close(c)

	time.Sleep(time.Second)
}

func main() {

	//channelDemo()

	bufferDemo()

	channelClose()
}
