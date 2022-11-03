package relationship

import (
	"fm/api/internal/pkg/common"
	"net/http"
)

// RetrieveFriendList retrieve friend list of an email
// pass a string of email
func (h Handler) RetrieveFriendList(w http.ResponseWriter, r *http.Request) {
	input, errD := DecodeAEmail(r)
	if errD != nil {
		common.ResponseJSON(w, http.StatusBadRequest, common.CommonErrorResponse{
			Code:        "invalid_request",
			Description: errD.Error(),
		})
		return
	}
	emails, errS := h.RelationshipSvc.FriendList(r.Context(), input)
	if errS != nil {
		common.ResponseJSON(w, http.StatusInternalServerError, common.CommonErrorResponse{
			Code:        "invalid_request",
			Description: errS.Error(),
		})
		return
	}

	common.ResponseJSON(w, http.StatusOK, toRetrieveFriendListResponse(dataToResponseArray(emails)))
}
