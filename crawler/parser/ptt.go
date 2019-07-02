package parser

import (
	"productnotify/crawler/model"
	"regexp"
	"time"
)

var pttRe = regexp.MustCompile(`<a href="(/bbs/[^\/]+/[A-Z0-9\.]+html)">([^<]+)</a>`)
var timeRe = regexp.MustCompile(`<div class="date">\s*([^<]+)</div>`)
var pttPrefix = "https://www.ptt.cc"

func ParsePtt(contents []byte) model.ParseResult {
	result := model.ParseResult{}

	matches := pttRe.FindAllSubmatch(contents, -1)
	dates := timeRe.FindAllSubmatch(contents, -1)
	date := getTodayFormat()

	for i, v := range matches {
		if string(dates[i][1]) != date {
			continue
		}

		result.Items = append(result.Items,
			model.Item{
				Url:   pttPrefix + string(v[1]),
				Title: string(v[2]),
			})
	}

	// if len(result.Items) > 0 {
	// 	nextRe := regexp.MustCompile(`<a class="btn wide" href="(/bbs/[a-zA-z\d]+/index[\d]+\.html)">&lsaquo; 上頁</a>`)
	// 	match := nextRe.FindAllSubmatch(contents, -1)

	// 	for _, v := range match {
	// 		result.Requests = append(result.Requests,
	// 			model.Request{
	// 				Url:       pttPrefix + string(v[1]),
	// 				Method:    "GET",
	// 				ParseFunc: ParsePtt,
	// 			})
	// 	}
	// }

	return result
}

func getTodayFormat() string {
	dt := time.Now()
	return dt.Format("1/02")
}
