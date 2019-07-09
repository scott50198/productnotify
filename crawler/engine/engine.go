package engine

import (
	"productnotify/crawler/handler"
	"productnotify/crawler/model"
	"productnotify/crawler/worker"
	"time"
)

type Engine struct {
	WorkCount         int
	Scheduler         Scheduler
	RestartTimeSecond int
	ItemHandler       handler.ItemHandler
}

func (e *Engine) Run(seeds ...model.Request) {

	e.build()

	dispatcher := worker.WorkerDispatcher{WorkerCount: 10}
	dispatcher.CreateWorker(e.Scheduler.GetRequestChan(), e.Scheduler.GetResultChan())
	e.submitSeeds(seeds...)

	for {
		select {
		case result := <-e.Scheduler.GetResultChan():
			for _, item := range result.Items {
				e.ItemHandler.Submit(item)
			}

			for _, request := range result.Requests {
				e.Scheduler.GetRequestChan() <- request
			}
		case <-time.Tick(time.Duration(e.RestartTimeSecond) * time.Second):
			e.submitSeeds(seeds...)
		}
	}
}

func (e *Engine) build() {
	e.Scheduler.Build()
	e.ItemHandler.Build()
}

func (e *Engine) submitSeeds(seeds ...model.Request) {
	for _, r := range seeds {
		e.Scheduler.GetRequestChan() <- r
	}
}
