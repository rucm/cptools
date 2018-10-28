package main

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	values := url.Values{}
	values.Add("name", "")
	values.Add("password", "")

	req, err := http.NewRequest(
		"POST",
		"https://practice.contest.atcoder.jp/login",
		strings.NewReader(values.Encode()),
	)

	if err != nil {
		panic(err)
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}
	client.CheckRedirect = func(req *http.Request, via []*http.Request) error {
		return http.ErrUseLastResponse
	}

	res, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	cookies := res.Cookies()
	for _, c := range cookies {
		fmt.Println(c.String())
	}
}
