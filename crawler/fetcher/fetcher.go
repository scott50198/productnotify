package fetcher

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func Get(url string) ([]byte, error) {
	req, _ := http.NewRequest(http.MethodGet, url, nil)
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/65.0.3325.181 Safari/537.36")
	httpClient := http.Client{}
	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("Fetch error : status code %d", resp.StatusCode)
	}
	// bufBody := bufio.NewReader(resp.Body)
	// utf8Reader := transform.NewReader(bufBody, determineEncoding(bufBody).NewDecoder())
	body, err := ioutil.ReadAll(resp.Body)
	return body, err
}

func Post(url string, body map[string]string) ([]byte, error) {

	return nil, nil
}

// // 根据html的meta头，试图自动转换编码到utf8
// func determineEncoding(r *bufio.Reader) encoding.Encoding {
// 	data, err := r.Peek(1024)
// 	if err != nil {
// 		return unicode.UTF8
// 	}
// 	e, _, _ := charset.DetermineEncoding(data, "")
// 	return e
// }
