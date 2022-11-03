package relationship

import (
	"encoding/json"
	"errors"
	"net/http"
)

// FriendConnectionRequest struct used when user request the service to get a list of friend emails
type FriendConnectionRequest struct {
	Friends []string `json:"friends"`
}

// FriendConnectionResponse struct used when the service return a list of friend emails
type FriendConnectionResponse struct {
	Success string `json:"success"`
}

func toFriendConnectionResponse() FriendConnectionResponse {
	return FriendConnectionResponse{Success: "true"}
}

func Decode2Mails(r *http.Request) (string, string, error) {
	var input FriendConnectionRequest
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		return "", "", err
	}
	if len(input.Friends) != 2 {
		return "", "", errors.New("need 2 emails")
	}
	return input.Friends[0], input.Friends[1], nil
}
