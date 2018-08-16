package persist

import "fmt"

func ItemSaver() chan interface{} {
	c := make(chan interface{})

	go func() {
		var count int
		for {
			item := <-c

			count++
			fmt.Printf("got item : %v, count = #%d#\n", item, count)
		}
	}()

	return c
}
