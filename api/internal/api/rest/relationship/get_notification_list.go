package relationship

import (
	"fm/api/internal/pkg/common"
	"net/http"
)

// GetNotificationList get a list of eligible for receive notification from sender
// pass a sender email and whoever email got mentioned in text
func (h Handler) GetNotificationList(w http.ResponseWriter, r *http.Request) {
	sender, text, err := decodeNoti(r)
	if err != nil {
		common.ResponseJSON(w, http.StatusBadRequest, common.CommonErrorResponse{
			Code:        "invalid_request",
			Description: err.Error(),
		})
		return
	}
	emails, errS := h.RelationshipSvc.NotificationList(r.Context(), sender, text)
	if errS != nil {
		common.ResponseJSON(w, http.StatusInternalServerError, common.CommonErrorResponse{
			Code:        "invalid_request",
			Description: errS.Error(),
		})
		return
	}
	common.ResponseJSON(w, http.StatusOK, toGetNotificationListResponse(dataToResponseArray(emails)))
}
