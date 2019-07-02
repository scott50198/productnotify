package controller

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"productnotify/crawler/config"
	"productnotify/crawler/model"
	"strings"
)

type ItemController struct {
	Items map[string]model.Item
}

func (c *ItemController) Build() {
	c.Items = make(map[string]model.Item)
}

func (c *ItemController) CheckItemExist(url string) bool {
	_, ok := c.Items[url]
	return ok
}

func (c *ItemController) LineNotify(item model.Item) {
	client := http.Client{}
	form := url.Values{}
	form.Add("message", item.Title+" : "+item.Url)
	req, _ := http.NewRequest("POST", "https://notify-api.line.me/api/notify", strings.NewReader(form.Encode()))
	req.Header.Add("Authorization", "Bearer "+config.LineNotifyAuthToken)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err.Error())
	}
	defer resp.Body.Close()

	contents, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(string(contents))
}
