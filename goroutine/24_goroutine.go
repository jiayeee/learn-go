package main

import (
	"fmt"
	"time"
)

/**
 * goroutine 可能切换的点：
 * I/O, select
 * channel
 * 等待锁
 * 函数调用(有时)
 * runtime.Gosched()
 */
func main() {

	for i := 0; i < 10; i++ {
		go func(i int) {
			for {
				fmt.Printf("goroutine %d\n", i)
			}
		}(i)
	}

	time.Sleep(time.Second)
}
