package main

import "fmt"

type Worker struct {
	in   chan int
	done chan bool
}

func doWork(i int, in chan int, done chan bool) {
	for n := range in {
		fmt.Printf("createWorker %d received %c\n", i, n)

		go func() { done <- true }()
	}
}

func createWorker(i int) Worker {
	worker := Worker{
		in:   make(chan int),
		done: make(chan bool),
	}
	go doWork(i, worker.in, worker.done)
	return worker
}
func channelDemo() {

	var workers [10]Worker
	for i := 0; i < 10; i++ {
		workers[i] = createWorker(i)
	}

	for i, worker := range workers {
		worker.in <- 'a' + i
	}

	for _, worker := range workers {
		<-worker.done
	}

	for i, worker := range workers {
		worker.in <- 'A' + i
	}

	// 等待响应
	for _, worker := range workers {
		<-worker.done
	}
}

func main() {
	channelDemo()
}
