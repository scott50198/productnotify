package parser

import (
	"productnotify/crawler/model"
	"regexp"
)

func ParseProxyList(contents []byte) ([]model.Proxy, error) {
	result := make([]model.Proxy, 0)
	proxyRe := regexp.MustCompile(`<td>([^<]+)</td>[\W]+<td>([\d]+)</td>`)

	matches := proxyRe.FindAllSubmatch(contents, -1)
	for _, v := range matches {
		result = append(result,
			model.Proxy{
				Host: string(v[1]),
				Port: string(v[2]),
			})
	}

	return result, nil
}
