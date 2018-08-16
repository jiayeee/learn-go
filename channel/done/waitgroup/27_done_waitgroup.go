package main

import (
	"fmt"
	"sync"
)

type Worker2 struct {
	in   chan int
	done func()
}

func doWork(i int, worker Worker2) {
	for n := range worker.in {
		fmt.Printf("createWorker %d received %c\n", i, n)

		worker.done()
	}
}

func createWorker(i int, wg *sync.WaitGroup) Worker2 {
	worker := Worker2{
		in: make(chan int),
		done: func() {
			wg.Done()
		},
	}
	go doWork(i, worker)
	return worker
}

func channelDemo() {
	var wg sync.WaitGroup
	var workers [10]Worker2

	for i := 0; i < 10; i++ {
		workers[i] = createWorker(i, &wg)
	}

	wg.Add(20)
	for i, worker := range workers {
		worker.in <- 'a' + i
	}

	for i, worker := range workers {
		worker.in <- 'A' + i
	}

	wg.Wait()
}

func main() {
	channelDemo()
}
