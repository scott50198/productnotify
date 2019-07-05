package main

import (
	"productnotify/crawler/engine"
	"productnotify/crawler/model"
	"productnotify/crawler/parser"
	"productnotify/crawler/scheduler"
)

func main() {
	fmt.Println("hello world")
	requests := makeSeeds()

	e := engine.ConcurrentEngine{
		WorkCount:         5,
		Scheduler:         scheduler.SimpleScheduler{},
		RestartTimeSecond: 60,
	}

	e.Run(requests...)

}

func makeSeeds() []model.Request {
	requests := make([]model.Request, 0)
	pttUrls := []string{"https://www.ptt.cc/bbs/MacShop/index.html", "https://www.ptt.cc/bbs/mobilesales/index.html", "https://www.ptt.cc/bbs/nb-shopping/index.html"}

	for i := 0; i < len(pttUrls); i++ {
		requests = append(requests, model.Request{
			Url:       pttUrls[i],
			Method:    "GET",
			ParseFunc: parser.ParsePtt,
		})
	}

	return requests
}
