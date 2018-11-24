package usecase

// RequestHandler : handler
type RequestHandler interface {
	Execute(url string, param Parameter) Response
}

// Parameter : request param (post request)
type Parameter interface {
	Del(key string) error
	Get(key string) (string, error)
	Set(key string, value string) error
}

// Response : response
type Response interface {
	Bind(obj interface{})
}
