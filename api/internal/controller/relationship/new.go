package relationship

import (
	"context"
	"fm/api/internal/repository/relationship"
	"fm/api/internal/repository/user"
)

// Service contains relationship services
type Service interface {
	//AddFriend Add friends between 2 emails
	AddFriend(ctx context.Context, requesterEmail string, addresseeEmail string) error
}

type impl struct {
	relationshipRepo relationship.Repository
	userRepo         user.Repository
}

// New DI
func New(relationshipRepo relationship.Repository, userRepo user.Repository) Service {
	return impl{
		relationshipRepo: relationshipRepo,
		userRepo:         userRepo,
	}
}
