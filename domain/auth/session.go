package auth

import (
	"encoding/gob"
	"io"
	"net/http"
)

// Session data
type Session map[string]string

// NewSession create session from cookies
func NewSession(cookies []*http.Cookie) *Session {
	session := Session{}
	for _, c := range cookies {
		session[c.Name] = c.Value
	}

	return &session
}

func (session *Session) Write(w io.Writer) {
	encoder := gob.NewEncoder(w)
	encoder.Encode(session)
}

func (session *Session) Read(r io.Reader) {
	decoder := gob.NewDecoder(r)
	decoder.Decode(session)
}
