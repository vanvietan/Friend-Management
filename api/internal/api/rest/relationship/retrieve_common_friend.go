package relationship

import (
	"fm/api/internal/pkg/common"
	"net/http"
)

// RetrieveCommonFriend retrieve common friend list from 2 emails
// Pass 2 emails in body
func (h Handler) RetrieveCommonFriend(w http.ResponseWriter, r *http.Request) {
	email1, email2, err := Decode2Mails(r)
	if err != nil {
		common.ResponseJSON(w, http.StatusBadRequest, common.CommonErrorResponse{
			Code:        "invalid_request",
			Description: err.Error(),
		})
		return
	}
	emails, errS := h.RelationshipSvc.CommonFriend(r.Context(), email1, email2)
	if errS != nil {
		common.ResponseJSON(w, http.StatusInternalServerError, common.CommonErrorResponse{
			Code:        "invalid_request",
			Description: errS.Error(),
		})
		return
	}

	common.ResponseJSON(w, http.StatusOK, toRetrieveFriendListResponse(dataToResponseArray(emails)))
}
