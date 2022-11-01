package relationship

import "fm/api/internal/controller/relationship"

type Handler struct {
	RelationshipSvc relationship.Service
}

func New(relationshipSvc relationship.Service) Handler {
	return Handler{
		RelationshipSvc: relationshipSvc,
	}
}
