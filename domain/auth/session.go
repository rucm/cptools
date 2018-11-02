package auth

import (
	"encoding/gob"
	"io"
)

// Session data
type Session struct {
	sessions map[string]string
}

func (s *Session) Write(w io.Writer) {
	encoder := gob.NewEncoder(w)
	encoder.Encode(&s.sessions)
}

func (s *Session) Read(r io.Reader) {
	decoder := gob.NewDecoder(r)
	decoder.Decode(&s.sessions)
}
