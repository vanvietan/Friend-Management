package relationship

import (
	"fm/api/internal/pkg/common"
	"net/http"
)

func (h Handler) AddFriend(w http.ResponseWriter, r *http.Request) {
	requestEmail, addressEmail, err := decode(r)
	if err != nil {
		common.ResponseJSON(w, http.StatusBadRequest, common.CommonErrorResponse{
			Code:        "invalid_request",
			Description: err.Error(),
		})
		return
	}
	errS := h.RelationshipSvc.AddFriend(r.Context(), requestEmail, addressEmail)
	if errS != nil {
		common.ResponseJSON(w, http.StatusInternalServerError, common.InternalCommonErrorResponse)
		return
	}

	common.ResponseJSON(w, http.StatusOK, toFriendConnectionResponse)
}
