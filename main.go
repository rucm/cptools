package main

import (
	"encoding/gob"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
)

func writeGob(filePath string, object interface{}) error {
	file, err := os.Create(filePath)
	if err == nil {
		encoder := gob.NewEncoder(file)
		encoder.Encode(object)
	}
	file.Close()
	return err
}

func readGob(filePath string, object interface{}) error {
	file, err := os.Create(filePath)
	if err == nil {
		decoder := gob.NewDecoder(file)
		err = decoder.Decode(object)
	}
	file.Close()
	return err
}

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
		log.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}
	client.CheckRedirect = func(req *http.Request, via []*http.Request) error {
		return http.ErrUseLastResponse
	}

	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	cookies := res.Cookies()
	for _, c := range cookies {
		log.Println(c.String())
	}

	err = writeGob("./save.txt", cookies)
	if err != nil {
		log.Fatal(err)
	}

	cookies2 := make([]*http.Cookie, 0)

	err = readGob("./save.txt", cookies2)
	if err != nil {
		log.Fatal(err)
	}

	for _, c := range cookies2 {
		log.Println(c.String())
	}
}
