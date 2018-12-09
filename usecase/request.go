package usecase

import (
	"github.com/rucm/cptools/domain"
)

// RequestHandler : handler
type RequestHandler interface {
	Execute(param Parameter) *Response
}

// Parameter : request param (post request)
type Parameter interface {
	Get(key string) string
	Set(key string, value string)
	Del(key string)
	Encode() string
}

// Response : response
type Response struct {
	Body    string
	Session *domain.Session
}
