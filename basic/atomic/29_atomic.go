package main

import (
	"fmt"
	"sync"
	"time"
)

type atomicInt struct {
	value int
	lock  sync.Mutex
}

func (a *atomicInt) increment() {
	fmt.Println("safe increment")

	func() {
		a.lock.Lock()
		defer a.lock.Unlock()
		a.value++
	}()
}

func (a *atomicInt) get() int {
	a.lock.Lock()
	defer a.lock.Unlock()
	return a.value
}

/**
 * go run 29_atomic.go
 * go run -race 29_atomic.go
 */
func main() {

	var a atomicInt

	for i := 0; i < 10; i++ {
		go func() {
			a.increment()
		}()
	}

	time.Sleep(time.Millisecond)
	fmt.Println(a.get())
}
