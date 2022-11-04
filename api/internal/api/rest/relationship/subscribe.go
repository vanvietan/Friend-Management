package relationship

import (
	"fm/api/internal/pkg/common"
	"net/http"
)

// Subscribe subscribe an addressee email
// Pass a requester email and addressee email
func (h Handler) Subscribe(w http.ResponseWriter, r *http.Request) {
	requester, addressee, err := Decode2MailsV2(r)
	if err != nil {
		common.ResponseJSON(w, http.StatusBadRequest, common.CommonErrorResponse{
			Code:        "invalid_request",
			Description: err.Error(),
		})
		return
	}
	errS := h.RelationshipSvc.Subscribe(r.Context(), requester, addressee)
	if errS != nil {
		common.ResponseJSON(w, http.StatusInternalServerError, common.CommonErrorResponse{
			Code:        "internal_server_error",
			Description: errS.Error(),
		})
		return
	}
	common.ResponseJSON(w, http.StatusOK, toConnectionResponse())

}
