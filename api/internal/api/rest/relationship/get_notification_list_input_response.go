package relationship

import (
	"encoding/json"
	"errors"
	"net/http"
)

// NotiRequest notification input
type NotiRequest struct {
	Sender string `json:"sender"`
	Text   string `json:"text"`
}

// decodeNoti decode
func decodeNoti(r *http.Request) (string, string, error) {
	var input NotiRequest
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		return "", "", err
	}
	if input.Sender == "" || input.Text == "" {
		return "", "", errors.New("invalid request, both sender and text must not be empty")
	}
	return input.Sender, input.Text, nil
}

type getNotificationListResponse struct {
	Success    string           `json:"success"`
	Recipients []AEmailResponse `json:"recipients"`
}

func toGetNotificationListResponse(array []AEmailResponse) getNotificationListResponse {
	return getNotificationListResponse{
		Success:    "true",
		Recipients: array,
	}
}
