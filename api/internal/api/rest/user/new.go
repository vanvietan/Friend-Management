package user

import "fm/api/internal/controller/user"

// Handler handle user calls
type Handler struct {
	UserSvc user.Service
}

// New DI
func New(userSvc user.Service) Handler {
	return Handler{UserSvc: userSvc}
}
