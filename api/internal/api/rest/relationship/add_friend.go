package relationship

import (
	"fm/api/internal/pkg/common"
	"net/http"
)

// AddFriend function works as a controller for creating friend connection between 2 user emails
// pass 2 emails
func (h Handler) AddFriend(w http.ResponseWriter, r *http.Request) {
	requestEmail, addressEmail, err := Decode2Mails(r)
	if err != nil {
		common.ResponseJSON(w, http.StatusBadRequest, common.CommonErrorResponse{
			Code:        "invalid_request",
			Description: err.Error(),
		})
		return
	}
	errS := h.RelationshipSvc.AddFriend(r.Context(), requestEmail, addressEmail)
	if errS != nil {
		common.ResponseJSON(w, http.StatusInternalServerError, common.CommonErrorResponse{
			Code:        "internal_server_error",
			Description: errS.Error(),
		})
		return
	}

	common.ResponseJSON(w, http.StatusOK, toConnectionResponse())
}
