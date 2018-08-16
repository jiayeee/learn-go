package engine

import "fmt"

type ConcurrentEngine struct {
	Scheduler   Scheduler
	WorkerCount int
}

type Scheduler interface {
	ReadyNotifier
	Submit(request Request)
	WorkerChannel() chan Request
	Run()
}

type ReadyNotifier interface {
	WorkerReady(chan Request)
}

func createWorker(in chan Request, out chan ParserResult, ready ReadyNotifier) {
	go func() {
		for {
			ready.WorkerReady(in)
			r := <-in
			result, e := worker(r)
			if e != nil {
				continue
			}
			out <- result
		}
	}()
}

func (ce *ConcurrentEngine) Run(seeds ...Request) {

	out := make(chan ParserResult)
	ce.Scheduler.Run()

	for i := 0; i < ce.WorkerCount; i++ {
		createWorker(ce.Scheduler.WorkerChannel(), out, ce.Scheduler)
	}

	for _, r := range seeds {
		ce.Scheduler.Submit(r)
	}

	var count int
	for {
		parserResult := <-out
		for _, item := range parserResult.Items {
			count++
			fmt.Printf("got item : %v, count = #%d#\n", item, count)
		}

		for _, request := range parserResult.Requests {
			ce.Scheduler.Submit(request)
		}
	}
}
