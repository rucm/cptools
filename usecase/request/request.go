package usecase

import (
	"github.com/rucm/cptools/domain"
)

// RequestHandler : handler
type RequestHandler interface {
	Execute() Response
}

// Parameter : request param (post request)
type Parameter interface {
	Get(key string) string
	Set(key string, value string)
	Del(key string) error
}

// Response : response
type Response interface {
	Bind(session *domain.Session)
}
