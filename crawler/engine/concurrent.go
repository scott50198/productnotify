package engine

import (
	"fmt"
	"productnotify/crawler/controller"
	"productnotify/crawler/fetcher"
	"productnotify/crawler/model"
	"productnotify/crawler/scheduler"
	"time"
)

type ConcurrentEngine struct {
	WorkCount         int
	Scheduler         scheduler.SimpleScheduler
	RestartTimeSecond int
}

func (e *ConcurrentEngine) Run(seeds ...model.Request) {

	e.Scheduler.Build()

	c := controller.ItemController{}
	c.Build()

	for i := 0; i < e.WorkCount; i++ {
		createWorker(e.Scheduler)
	}

	e.submitSeeds(seeds...)

	for {
		select {
		case result := <-e.Scheduler.GetResultChan():
			for _, item := range result.Items {
				fmt.Printf("got item : %v\n", item)
				if !c.CheckItemExist(item.Url) {
					c.Items[item.Url] = item
					c.LineNotify(item)
				}
			}

			for _, request := range result.Requests {
				e.Scheduler.Submit(request)
			}
		case <-time.Tick(time.Duration(e.RestartTimeSecond) * time.Second):
			e.submitSeeds(seeds...)
		}
	}
}

func (e *ConcurrentEngine) submitSeeds(seeds ...model.Request) {
	for _, r := range seeds {
		e.Scheduler.Submit(r)
	}
}

func createWorker(s scheduler.SimpleScheduler) {
	go func() {
		for {
			request := s.GetResuest()
			result, err := worker(request)
			if err != nil {
				fmt.Println(err.Error())
				continue
			}
			s.SendResult(result)
		}
	}()
}

func worker(r model.Request) (model.ParseResult, error) {
	if r.Method == "GET" {
		body, err := fetcher.Get(r.Url)
		if err != nil {
			fmt.Println("Fetcher err : " + err.Error())
			return model.ParseResult{}, err
		}

		return r.ParseFunc(body), nil
	}
	return model.ParseResult{}, nil
}
