package user

import (
	"fm/api/internal/repository/user"
)

// Service contains user services
type Service interface {
}

type impl struct {
	userRepo user.Repository
}

// New
func New(userRepo user.Repository) Service {
	return impl{userRepo: userRepo}
}
