package main

import (
	"productnotify/crawler/engine"
	"productnotify/crawler/model"
	"productnotify/crawler/parser"
)

func main() {
	requests := makeSeeds()

	e := engine.Engine{
		WorkCount:         5,
		Scheduler:         engine.Scheduler{},
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
