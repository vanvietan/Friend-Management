package rest

import (
	"fm/api/internal/api/rest/relationship"
	"fm/api/internal/api/rest/user"
	relationshipSvc "fm/api/internal/controller/relationship"
	userSvc "fm/api/internal/controller/user"
)

// Handler master Handler
type Handler struct {
	UserHandler         user.Handler
	RelationshipHandler relationship.Handler
}

// New DI
func New(userSvc userSvc.Service, relationshipSvc relationshipSvc.Service) Handler {
	return Handler{
		UserHandler:         user.Handler{UserSvc: userSvc},
		RelationshipHandler: relationship.Handler{RelationshipSvc: relationshipSvc},
	}
}
