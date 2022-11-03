package relationship

import "fm/api/internal/controller/relationship"

// Handler relationship Handler
type Handler struct {
	RelationshipSvc relationship.Service
}

// New DI
func New(relationshipSvc relationship.Service) Handler {
	return Handler{
		RelationshipSvc: relationshipSvc,
	}
}
