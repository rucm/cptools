package atcoder

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"

	"github.com/rucm/cptools/domain"
	"github.com/rucm/cptools/usecase"
)

// RequestHandler :
type RequestHandler struct {
	URL         string
	Method      string
	ContentType string
	Session     *domain.Session
}

// Parameter :
type Parameter struct {
	Values url.Values
}

// Execute :
func (handler *RequestHandler) Execute(param usecase.Parameter) *usecase.Response {

	req := createRequest(
		handler.Method,
		handler.URL,
		param,
	)

	if handler.Session != nil {
		for key, val := range *handler.Session {
			req.AddCookie(&http.Cookie{Name: key, Value: val})
		}
	}

	client := &http.Client{}
	client.CheckRedirect = func(req *http.Request, via []*http.Request) error {
		return http.ErrUseLastResponse
	}

	req.Header.Set("Content-Type", handler.ContentType)
	res, err := client.Do(req)

	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()

	response := createResponse(res)

	return response
}

// NewParameter :
func NewParameter() *Parameter {
	param := &Parameter{}
	param.Values = url.Values{}
	return param
}

// Get :
func (param *Parameter) Get(key string) string {
	return param.Values.Get(key)
}

// Set :
func (param *Parameter) Set(key string, value string) {
	param.Values.Set(key, value)
}

// Del :
func (param *Parameter) Del(key string) {
	param.Values.Del(key)
}

// Encode :
func (param *Parameter) Encode() string {
	return param.Values.Encode()
}

func createRequest(method string, URL string, param usecase.Parameter) *http.Request {

	encode := ""

	if param != nil {
		encode = param.Encode()
	}

	req, err := http.NewRequest(
		method,
		URL,
		strings.NewReader(encode),
	)

	if err != nil {
		log.Fatal(err)
	}

	return req
}

func createResponse(res *http.Response) *usecase.Response {

	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		log.Fatal(err)
	}

	session := &domain.Session{}
	for _, cookie := range res.Cookies() {
		(*session)[cookie.Name] = cookie.Value
	}

	response := &usecase.Response{
		StatusCode: res.StatusCode,
		Body:       string(body),
		Session:    session,
	}

	return response
}
