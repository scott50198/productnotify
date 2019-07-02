package scheduler

import (
	"productnotify/crawler/model"
)

type SimpleScheduler struct {
	requestChan chan model.Request
	resultChan  chan model.ParseResult
}

func (s *SimpleScheduler) GetResultChan() chan model.ParseResult {
	return s.resultChan
}

func (s *SimpleScheduler) Build() {
	s.requestChan = make(chan model.Request)
	s.resultChan = make(chan model.ParseResult)
}

func (s *SimpleScheduler) Submit(request model.Request) {
	go func() {
		s.requestChan <- request
	}()
}

func (s *SimpleScheduler) SendResult(result model.ParseResult) {
	s.resultChan <- result
}

func (s *SimpleScheduler) GetResult() model.ParseResult {
	return <-s.resultChan
}

func (s *SimpleScheduler) GetResuest() model.Request {
	return <-s.requestChan
}
