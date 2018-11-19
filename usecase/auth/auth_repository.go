package usecase

import "github.com/rucm/cptools/domain"

// AuthRepository : base handler
type AuthRepository interface {
	Write(session *domain.Session)
	Read() *domain.Session
}
