package scheduler

import (
	"learn-go/crawler/engine"
)

type QueuedScheduler struct {
	requestChannel chan engine.Request
	workerChannel  chan chan engine.Request
}

func (qs *QueuedScheduler) WorkerChannel() chan engine.Request {
	return make(chan engine.Request)
}

func (qs *QueuedScheduler) WorkerReady(w chan engine.Request) {
	qs.workerChannel <- w
}

func (qs *QueuedScheduler) Submit(r engine.Request) {
	qs.requestChannel <- r
}

func (qs *QueuedScheduler) Run() {
	qs.requestChannel = make(chan engine.Request)
	qs.workerChannel = make(chan chan engine.Request)

	go func() {
		var requestQ []engine.Request
		var workerQ []chan engine.Request

		for {
			var activeRequest engine.Request
			var acviveWorker chan engine.Request

			if len(requestQ) > 0 && len(workerQ) > 0 {
				activeRequest = requestQ[0]
				acviveWorker = workerQ[0]
			}

			select {
			case r := <-qs.requestChannel:
				requestQ = append(requestQ, r)
			case w := <-qs.workerChannel:
				workerQ = append(workerQ, w)
			case acviveWorker <- activeRequest:
				requestQ = requestQ[1:]
				workerQ = workerQ[1:]
			}
		}
	}()
}
