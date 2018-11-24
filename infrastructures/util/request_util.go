package infrastractures

import (
	"log"
	"net/http"
	"net/url"
	"strings"
)

// CreateRequest : create http request
func CreateRequest(url string, method string, values url.Values) *http.Request {

	req, err := http.NewRequest(
		method,
		url,
		strings.NewReader(values.Encode()),
	)

	if err != nil {
		log.Fatal(err)
	}

	return req
}

// GetResponse : execute request
func GetResponse(req *http.Request, contentType string) *http.Response {

	client := &http.Client{}
	client.CheckRedirect = func(req *http.Request, via []*http.Request) error {
		return http.ErrUseLastResponse
	}

	req.Header.Set("Content-Type", contentType)
	res, err := client.Do(req)

	if err != nil {
		log.Fatal(err)
	}

	return res
}
