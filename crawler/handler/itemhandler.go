package handler

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"productnotify/crawler/config"
	"productnotify/crawler/model"
	"strings"
)

type ItemHandler struct {
	items    map[string]bool
	itemChan chan model.Item
}

func (this *ItemHandler) Build() {
	this.items = make(map[string]bool)
	this.itemChan = make(chan model.Item)
	go func() {
		for {
			item := <-this.itemChan
			this.handleItem(item)
		}
	}()
}

func (this *ItemHandler) GetItemChan() chan model.Item {
	return this.itemChan
}

func (this *ItemHandler) handleItem(item model.Item) {
	fmt.Printf("got item : %v\n", item)
	if this.checkItemExist(item) {
		this.lineNotify(item)
	}
}

func (this *ItemHandler) checkItemExist(item model.Item) bool {
	_, ok := this.items[item.Url]
	return ok
}

func (this *ItemHandler) lineNotify(item model.Item) {
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
