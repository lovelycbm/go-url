package main

import (
	"errors"
	"fmt"
	"net/http"
)

var errRequestFailed = errors.New("Request Failed")

type requestResult struct {
	url    string
	status string
}

func main() {
	results := make(map[string]string)
	c := make(chan requestResult)
	urls := []string{
		"https://www.amazon.com",
		"https://www.reddit.com",
		"https://www.google.com",
		"https://soundcloud.com",
		"https://www.facebook.com",
		"https://www.instagram.com",
		"https://academy.nomadcoders.co",
	}

	for _, url := range urls {
		go hitUrl(url, c)
	}

	for i := 0; i < len(urls); i++ {
		//fmt.Println(<-c)
		result := <-c
		results[result.url] = result.status
	}

	for url, status := range results {
		fmt.Println(url, status)
	}
}

// (파라미터) chan 뒤에 화살표 형으로 넣으면 함수내에서 보내는것밖에 안됨
func hitUrl(url string, c chan<- requestResult) {
	resp, err := http.Get(url)
	status := "OK"
	if err != nil || resp.StatusCode >= 400 {
		status = "FAILED"
	}
	c <- requestResult{url: url, status: status}

}
