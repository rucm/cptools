package usecase

import "github.com/rucm/cptools/domain"

// SessionRepository : base handler
type SessionRepository interface {
	Write(session *domain.Session)
	Read() *domain.Session
}
