package atcoder

import (
	"log"
	"net/http"
	"net/url"
	"strings"

	"github.com/rucm/cptools/domain/auth"
)

const (
	loginURL    = "https://practice.contest.atcoder.jp/login"
	contentType = "application/x-www-form-urlencoded"
)

// Auth structure
type Auth struct {
	session *auth.Session
}

// Login to AtCoder
func (a *Auth) Login(user string, password string) (*auth.Session, error) {
	return nil, nil
}

// Logout to AtCoder
func (a *Auth) Logout(session *auth.Session) error {
	return nil
}

func createPostParam(user string, password string) url.Values {
	values := url.Values{}
	values.Add("name", user)
	values.Add("password", password)
	return values
}

func createRequest(values url.Values) *http.Request {
	req, err := http.NewRequest(
		"POST",
		loginURL,
		strings.NewReader(values.Encode()),
	)

	if err != nil {
		log.Fatal(err)
	}

	return req
}

func getResponse(req *http.Request) *http.Response {
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
