package user

import (
	"context"
	"fm/api/internal/models"
	"fm/api/internal/repository/user"
)

// Service contains user services
type Service interface {
	//FindUserByEmail find user by email
	FindUserByEmail(ctx context.Context, input string) (models.User, error)
}

type impl struct {
	userRepo user.Repository
}

// New
func New(userRepo user.Repository) Service {
	return impl{userRepo: userRepo}
}
