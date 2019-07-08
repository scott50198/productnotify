package engine

import (
	"fmt"
	"productnotify/crawler/fetcher"
	"productnotify/crawler/handler"
	"productnotify/crawler/model"
	"time"
)

type Engine struct {
	WorkCount         int
	Scheduler         Scheduler
	RestartTimeSecond int
	ItemHandler       handler.ItemHandler
	ErrorHandler      handler.ErrorHandler
}

func (e *Engine) Run(seeds ...model.Request) {

	e.build()

	for i := 0; i < e.WorkCount; i++ {
		createWorker(e.Scheduler, e.Scheduler.errorChan)
	}

	e.submitSeeds(seeds...)

	for {
		select {
		case result := <-e.Scheduler.GetResultChan():
			for _, item := range result.Items {
				e.ItemHandler.GetItemChan() <- item
			}

			for _, request := range result.Requests {
				e.Scheduler.GetRequestChan() <- request
			}
		case <-time.Tick(time.Duration(e.RestartTimeSecond) * time.Second):
			e.submitSeeds(seeds...)
		case err := <-e.Scheduler.GetErrorChan():
			fmt.Println(err.Error())
		}
	}
}

func (e *Engine) build() {
	e.Scheduler.Build()
	e.ItemHandler.Build()
	e.ErrorHandler.Build()
}

func (e *Engine) submitSeeds(seeds ...model.Request) {
	for _, r := range seeds {
		e.Scheduler.GetRequestChan() <- r
	}
}

func createWorker(s Scheduler, errorChan chan error) {
	go func() {
		for {
			request := <-s.GetRequestChan()
			result, err := worker(request)
			if err != nil {
				errorChan <- err
				continue
			}
			s.GetResultChan() <- result
		}
	}()
}

func worker(r model.Request) (model.ParseResult, error) {
	if r.Method == "GET" {
		body, err := fetcher.Get(r.Url)
		if err != nil {
			return model.ParseResult{}, err
		}
		return r.ParseFunc(body), nil
	}
	return model.ParseResult{}, nil
}
