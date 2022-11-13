package relationship

import (
	"encoding/json"
	"errors"
	"fm/api/internal/models"
	"net/http"
)

// AEmailRequest a single email request
type AEmailRequest struct {
	Email string `json:"email"`
}

// AEmailResponse a single email response
type AEmailResponse struct {
	Email string `json:"email"`
}

type FriendListResponse struct {
	Success string           `json:"success"`
	Friends []AEmailResponse `json:"friends"`
	Count   int              `json:"count"`
}

func toRetrieveFriendListResponse(emails []AEmailResponse) FriendListResponse {
	return FriendListResponse{
		Success: "true",
		Friends: emails,
		Count:   len(emails),
	}
}

// DecodeAEmail decode 1 email from Body
func DecodeAEmail(r *http.Request) (string, error) {
	var input AEmailRequest
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		return "", errors.New("invalid email request")
	}
	if input.Email == "" {
		return "", errors.New("invalid email request")
	}
	return input.Email, nil
}

func dataToResponseArray(e []models.User) []AEmailResponse {
	if len(e) == 0 {
		return nil
	}
	resp := make([]AEmailResponse, len(e))
	for i, s := range e {
		resp[i].Email = s.Email
	}
	return resp
}
