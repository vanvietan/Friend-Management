package relationship

import (
	"encoding/json"
	"errors"
	"net/http"
)

// SubscribeRequest request for subscribe
type SubscribeRequest struct {
	Requester string `json:"requester"`
	Target    string `json:"target"`
}

// Decode2MailsV2 decode with requester and target
func Decode2MailsV2(r *http.Request) (string, string, error) {
	var input SubscribeRequest
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		return "", "", err
	}
	if input.Requester == "" || input.Target == "" {
		return "", "", errors.New("invalid email")
	}
	return input.Requester, input.Target, nil
}
