package scheduler

import "learn-go/crawler/engine"

type SimpleScheduler struct {
	workChannel chan engine.Request
}

func (s *SimpleScheduler) WorkerChannel() chan engine.Request {
	return s.workChannel
}

func (s *SimpleScheduler) WorkerReady(chan engine.Request) {
}

func (s *SimpleScheduler) Run() {
	s.workChannel = make(chan engine.Request)
}

func (s *SimpleScheduler) Submit(r engine.Request) {
	go func() {
		s.workChannel <- r
	}()
}
