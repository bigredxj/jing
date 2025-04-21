package http

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func Get(url string, headers map[string]string) string {
	client := &http.Client{
		Timeout: 30 * time.Second,
	}

	// 创建HTTP请求
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	if len(headers) > 0 {
		for k, v := range headers {
			req.Header.Add(k, v)
		}
	}

	// 发送HTTP请求
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return ""
	}

	// 处理HTTP响应
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	return string(body)
}
