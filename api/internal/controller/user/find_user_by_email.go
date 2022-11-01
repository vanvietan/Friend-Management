package user

import (
	"context"
	"fm/api/internal/models"
	log "github.com/sirupsen/logrus"
)

func (i impl) FindUserByEmail(ctx context.Context, input string) (models.User, error) {
	//check input valid

	email, err := i.userRepo.FindUserByEmail(ctx, input)
	if err != nil {
		log.Printf("error when find user %v", err)
		return models.User{}, err
	}

	return email, nil
}
