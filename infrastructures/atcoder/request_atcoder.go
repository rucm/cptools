package infrastractures

import (
	"log"
	"net/http"
	"net/url"
	"strings"

	"github.com/rucm/cptools/domain"
	"github.com/rucm/cptools/usecase/request"
)

// AtCoderRequestHandler :
type AtCoderRequestHandler struct {
	Request *http.Request
}

// AtCoderParameter :
type AtCoderParameter struct {
	Values url.Values
}

// AtCoderResponse :
type AtCoderResponse struct {
	Response *http.Response
}

// NewAtCoderRequestHandler : create new handler
func NewAtCoderRequestHandler(url string, param AtCoderParameter) usecase.RequestHandler {

	req, err := http.NewRequest(
		"POST",
		url,
		strings.NewReader(param.Values.Encode()),
	)

	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	handler := &AtCoderRequestHandler{
		req,
	}

	return handler
}

// Execute : execute request
func (handler *AtCoderRequestHandler) Execute() usecase.Response {

	client := &http.Client{}
	client.CheckRedirect = func(req *http.Request, via []*http.Request) error {
		return http.ErrUseLastResponse
	}

	res, err := client.Do(handler.Request)

	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()

	response := &AtCoderResponse{
		res,
	}

	return response
}

// Get : AtCoderParameter
func (param *AtCoderParameter) Get(key string) string {

	return param.Values.Get(key)
}

// Set : AtCoderParameter
func (param *AtCoderParameter) Set(key string, value string) {

	param.Values.Set(key, value)
}

// Del : AtCoderParameter
func (param *AtCoderParameter) Del(key string) {

	param.Values.Del(key)
}

// Bind : AtCoderResponse
func (res *AtCoderResponse) Bind(session *domain.Session) {

	for _, c := range res.Response.Cookies() {
		(*session)[c.Name] = c.Value
	}
}
