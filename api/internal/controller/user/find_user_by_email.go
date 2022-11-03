package user

import (
	"context"
	"fm/api/internal/models"
	"fm/api/internal/pkg"
	log "github.com/sirupsen/logrus"
)

// FindUserByEmail find user by email
func (i impl) FindUserByEmail(ctx context.Context, input string) (models.User, error) {
	//check input valid
	errC := pkg.CheckValidEmail(input)
	if errC != nil {
		log.Printf("error when checking validation, %v", errC)
		return models.User{}, errC
	}

	email, err := i.userRepo.FindUserByEmail(ctx, input)
	if err != nil {
		log.Printf("error when find user %v", err)
		return models.User{}, err
	}

	return email, nil
}
