package relationship

import (
	"fm/api/internal/pkg/common"
	"net/http"
)

// Block to block target by requester
// pass requester email and target email
func (h Handler) Block(w http.ResponseWriter, r *http.Request) {
	requester, addressee, err := Decode2MailsV2(r)
	if err != nil {
		common.ResponseJSON(w, http.StatusBadRequest, common.CommonErrorResponse{
			Code:        "invalid_request",
			Description: err.Error(),
		})
		return
	}
	errS := h.RelationshipSvc.Block(r.Context(), requester, addressee)
	if errS != nil {
		common.ResponseJSON(w, http.StatusInternalServerError, common.CommonErrorResponse{
			Code:        "internal_server_error",
			Description: errS.Error(),
		})
		return
	}
	common.ResponseJSON(w, http.StatusOK, toConnectionResponse())
}
