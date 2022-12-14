package router

import (
	"fm/api/internal/api/rest"
	relationshipSvc "fm/api/internal/controller/relationship"
	userSvc "fm/api/internal/controller/user"
	"github.com/go-chi/chi/v5"
)

// MasterRoute masterRoute
type MasterRoute struct {
	Router              *chi.Mux
	Handler             rest.Handler
	UserService         userSvc.Service
	RelationshipService relationshipSvc.Service
}

// New DI
func New(r *chi.Mux, userSvc userSvc.Service, relationshipSvc relationshipSvc.Service) {
	newHandler := rest.New(userSvc, relationshipSvc)
	mr := MasterRoute{
		Router:  r,
		Handler: newHandler,
	}
	mr.initRoutes()
}

func (mr MasterRoute) initRoutes() {
	mr.initUserRoutes()
	mr.initRelationshipRoutes()
}

func (mr MasterRoute) initUserRoutes() {
	mr.Router.Group(func(r chi.Router) {
		r.Post("/add-user", mr.Handler.UserHandler.AddUser)
	})
}

func (mr MasterRoute) initRelationshipRoutes() {
	mr.Router.Group(func(r chi.Router) {
		r.Post("/add-friend", mr.Handler.RelationshipHandler.AddFriend)
		r.Post("/friend-list", mr.Handler.RelationshipHandler.RetrieveFriendList)
		r.Post("/common-friend", mr.Handler.RelationshipHandler.RetrieveCommonFriend)
		r.Post("/subscribe", mr.Handler.RelationshipHandler.Subscribe)
		r.Post("/block", mr.Handler.RelationshipHandler.Block)
		r.Post("/notification-list", mr.Handler.RelationshipHandler.GetNotificationList)
	})
}
