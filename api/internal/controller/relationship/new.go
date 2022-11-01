package relationship

import "fm/api/internal/repository/relationship"

// Service contains relationship services
type Service interface {
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
