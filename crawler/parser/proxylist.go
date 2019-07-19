package parser

import (
	"productnotify/crawler/model"
	"regexp"
)

func ParseProxyList(contents []byte) ([]model.Proxy, error) {
	result := make([]model.Proxy, 0)
	proxyRe := regexp.MustCompile(`<td>([\d]+\.[\d]+\.[\d]+\.[\d]+)</td>`)
	portRe := regexp.MustCompile(`<td>([\d]+)</td>`)
	locationRe := regexp.MustCompile(`<td class='hm'>([A-Z][^<]+)</td>`)
	secureRe := regexp.MustCompile(`<td class='hx'>([\w]+)</td>`)

	proxys := proxyRe.FindAllSubmatch(contents, -1)
	ports := portRe.FindAllSubmatch(contents, -1)
	locations := locationRe.FindAllSubmatch(contents, -1)
	secures := secureRe.FindAllSubmatch(contents, -1)

	for i := 0; i < len(proxys); i++ {
		result = append(result,
			model.Proxy{
				Host:     string(proxys[i][1]),
				Port:     string(ports[i][1]),
				Location: string(locations[i][1]),
				Scheme:   string(secures[i][1]),
			})
	}

	return result, nil
}
