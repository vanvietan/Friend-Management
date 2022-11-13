package user

import (
	"encoding/json"
	"fm/api/internal/models"
	"net/http"
)

type AUserInput struct {
	Email string `json:"email"`
}
type AddUserResponse struct {
	ID    int64  `json:"ID"`
	Email string `json:"email"`
}

// DecodeUser from body
func DecodeUser(r *http.Request) (models.User, error) {
	var input AUserInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		return models.User{}, err
	}

	return models.User{
		Email: input.Email,
	}, nil
}

func toAddUserResponse(user models.User) AddUserResponse {
	return AddUserResponse{
		ID:    user.ID,
		Email: user.Email,
	}
}
