package main

import (
	"fmt"
	"math/rand"
	"time"
)

func createWorker(i int) chan int {
	c := make(chan int)
	go func() {
		for n := range c {
			time.Sleep(time.Second)
			fmt.Printf("createWorker %d received %d\n", i, n)
		}
	}()
	return c
}

func generate() chan int {
	out := make(chan int)
	go func() {
		i := 0
		for {
			time.Sleep(time.Duration(rand.Intn(1500)) * time.Millisecond)

			out <- i
			i++
		}
	}()

	return out
}

// selece 可以达到 非阻塞式的接收数据
func main() {

	var c1, c2 = generate(), generate()
	w := createWorker(0)

	// 定时器
	tm := time.After(10 * time.Second)
	tick := time.Tick(time.Second)

	var values []int
	var activeWorker chan int
	var activeValue int
	for {
		if len(values) > 0 {
			activeWorker = w
			activeValue = values[0]
		}

		select {
		case n := <-c1:
			//fmt.Println("received from c1 : ", n)
			values = append(values, n)
		case n := <-c2:
			//fmt.Println("received from c2 : ", n)
			values = append(values, n)

			// 每隔1秒钟看一下 values 中的值
		case <-tick:
			fmt.Printf("len(values) = %d, values = %v\n", len(values), values)

			// 800 毫秒之内没有接到任何数据就报超时
		case <-time.After(800 * time.Millisecond):
			fmt.Println("time out")

			// nil channel 在 select 里面是能正确运行的，但是不会被 select 到
		case activeWorker <- activeValue:
			values = values[1:]

			// 定时器到时会送一个时间到 channel 里面
		case <-tm:
			fmt.Println("bye")
			return
		}
	}
}
