package handler

import (
	"fmt"
	"productnotify/crawler/fetcher"
	"productnotify/crawler/model"
	"productnotify/crawler/parser"
)

type ProxyHandler struct {
	errorChan chan error
	proxyList []model.Proxy
}

func (this *ProxyHandler) Build() {
	this.errorChan = make(chan error)
	go func() {
		this.fetchProxyList()
		for {
			err := <-this.errorChan
			this.handleError(err)
		}
	}()
}

func (this *ProxyHandler) fetchProxyList() {
	contents, err := fetcher.Get("https://free-proxy-list.net/")
	if err != nil {
		fmt.Println(err.Error())
	}
	list, err := parser.ParseProxyList(contents)
	if err != nil {
		fmt.Println(err.Error())
	}
	this.proxyList = list
}

func (this *ProxyHandler) GetErrorChan() chan error {
	return this.errorChan
}

func (this *ProxyHandler) handleError(err error) {
	fmt.Println(err.Error())
	// proxyList, err := worker.GetProxyList("https://free-proxy-list.net/")
	// if err != nil {
	// 	fmt.Println(err.Error())
	// }
	// this.proxyList = proxyList
}

func (this *ProxyHandler) GetProxyList() []model.Proxy {
	return this.proxyList
}

func GetProxyList(url string) ([]model.Proxy, error) {
	result := make([]model.Proxy, 0)
	body, err := fetcher.Get(url)
	if err != nil {
		return result, err
	}
	return parser.ParseProxyList(body)
}
