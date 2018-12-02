package infrastractures

import (
	"log"
	"net/http"
	"net/url"
	"strings"

	"github.com/rucm/cptools/domain"
	"github.com/rucm/cptools/usecase/request"
)

// AuthRequestHandler :
type AuthRequestHandler struct {
	Request *http.Request
}

// AuthParameter :
type AuthParameter struct {
	Values url.Values
}

// AuthResponse :
type AuthResponse struct {
	Response *http.Response
}

// NewAuthRequestHandler : create new handler
func NewAuthRequestHandler(url string, param AuthParameter) usecase.RequestHandler {

	req, err := http.NewRequest(
		"POST",
		url,
		strings.NewReader(param.Values.Encode()),
	)

	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	handler := &AuthRequestHandler{
		req,
	}

	return handler
}

// Execute : execute request
func (handler *AuthRequestHandler) Execute() usecase.Response {

	client := &http.Client{}
	client.CheckRedirect = func(req *http.Request, via []*http.Request) error {
		return http.ErrUseLastResponse
	}

	res, err := client.Do(handler.Request)

	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()

	response := &AuthResponse{
		res,
	}

	return response
}

// Get : AuthParameter
func (param *AuthParameter) Get(key string) string {

	return param.Values.Get(key)
}

// Set : AuthParameter
func (param *AuthParameter) Set(key string, value string) {

	param.Values.Set(key, value)
}

// Del : AuthParameter
func (param *AuthParameter) Del(key string) {

	param.Values.Del(key)
}

// Bind : AuthResponse
func (res *AuthResponse) Bind(session *domain.Session) {

	for _, c := range res.Response.Cookies() {
		(*session)[c.Name] = c.Value
	}
}
