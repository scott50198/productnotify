package worker

import (
	"productnotify/crawler/fetcher"
	"productnotify/crawler/model"
)

type WorkerDispatcher struct {
	WorkerCount int
	proxyHelper ProxyHelper
}

type ProxyHelper interface {
	ChangeProxy() model.Proxy
}

func (this *WorkerDispatcher) CreateWorker(in chan model.Request, out chan model.ParseResult) {
	for i := 0; i < this.WorkerCount; i++ {
		go func() {
			for {
				request := <-in
				result, err := GetParseResult(request)
				if err != nil {
					in <- request
					continue
				}
				out <- result
			}
		}()
	}
}

func GetParseResult(r model.Request) (model.ParseResult, error) {
	if r.Method == "GET" {
		body, err := fetcher.Get(r.Url)
		if err != nil {
			return model.ParseResult{}, err
		}
		return r.ParseFunc(body), nil
	}
	return model.ParseResult{}, nil
}
