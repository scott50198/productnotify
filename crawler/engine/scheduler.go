package engine

import (
	"productnotify/crawler/model"
)

type Scheduler struct {
	requestChan chan model.Request
	resultChan  chan model.ParseResult
}

func (s *Scheduler) Build() {
	s.requestChan = make(chan model.Request)
	s.resultChan = make(chan model.ParseResult)
}

func (s *Scheduler) GetResultChan() chan model.ParseResult {
	return s.resultChan
}

func (s *Scheduler) GetRequestChan() chan model.Request {
	return s.requestChan
}
