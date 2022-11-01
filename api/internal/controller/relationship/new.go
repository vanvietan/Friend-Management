package relationship

import (
	"context"
	"fm/api/internal/repository/relationship"
)

// Service contains relationship services
type Service interface {
	AddFriend(ctx context.Context, requesterID int64, addresseeID int64) error
}

type impl struct {
	relationshipRepo relationship.Repository
}

// New DI
func New(relationshipRepo relationship.Repository) Service {
	return impl{
		relationshipRepo: relationshipRepo,
	}
}
