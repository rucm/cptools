package infrastractures

import (
	"net/url"
)

// AtCoderRequestHandler :
type AtCoderRequestHandler struct {
}

// AtCoderParameter :
type AtCoderParameter struct {
	Values url.Values
}

// AtCoderResponse :
type AtCoderResponse struct {
	Values map[string]string
}
