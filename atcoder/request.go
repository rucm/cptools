package atcoder

import (
	"io/ioutil"
	"log"
	"net/http"
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

// Execute :
func (handler *RequestHandler) Execute(param usecase.Parameter) *usecase.Response {

	req := createRequest(
		handler.Method,
		handler.URL,
		param.Encode(),
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

func createRequest(method string, URL string, param string) *http.Request {

	req, err := http.NewRequest(
		method,
		URL,
		strings.NewReader(param),
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
		Body:    string(body),
		Session: session,
	}

	return response
}
