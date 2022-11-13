package user

import (
	"fm/api/internal/pkg/common"
	"net/http"
)

// AddUser add a user
// pass a body
func (h Handler) AddUser(w http.ResponseWriter, r *http.Request) {
	input, err := DecodeUser(r)
	if err != nil {
		common.ResponseJSON(w, http.StatusBadRequest, common.CommonErrorResponse{
			Code:        "invalid_request",
			Description: err.Error(),
		})
		return
	}
	user, errS := h.UserSvc.AddUser(r.Context(), input)
	if errS != nil {
		common.ResponseJSON(w, http.StatusInternalServerError, common.CommonErrorResponse{
			Code:        "internal_server_error",
			Description: errS.Error(),
		})
		return
	}
	common.ResponseJSON(w, http.StatusOK, toAddUserResponse(user))
}
