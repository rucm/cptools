package usecase

// RequestHandler : handler
type RequestHandler interface {
	Execute(param Parameter) Response
}

// Parameter : request param (post request)
type Parameter interface {
	Del(key string) error
	Get(key string) (string, error)
	Set(key string, value string) error
}

// Response : response
type Response interface {
	Get(key string) string
	Exists(key string) bool
}
