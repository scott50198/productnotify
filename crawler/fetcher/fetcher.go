package fetcher

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"productnotify/crawler/model"
	"time"
)

func Get(url string) ([]byte, error) {
	req, _ := http.NewRequest(http.MethodGet, url, nil)
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/65.0.3325.181 Safari/537.36")

	httpClient := &http.Client{}

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("Fetch error : status code %d", resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
	return body, err
}

func ProxyGet(site string, proxy model.Proxy) ([]byte, error) {
	proxyUrl, _ := url.Parse(proxy.Scheme + "://" + proxy.Host + ":" + proxy.Port)
	myClient := &http.Client{
		Transport: &http.Transport{Proxy: http.ProxyURL(proxyUrl)},
		Timeout:   time.Duration(5 * time.Second),
	}

	resp, err := myClient.Get(site)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("Fetch error ,Status code : ", resp.StatusCode)
	}

	return ioutil.ReadAll(resp.Body)
}
