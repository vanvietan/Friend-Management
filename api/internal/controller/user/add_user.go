package user

import (
	"context"
	"fm/api/internal/models"
	"fm/api/internal/pkg"
	log "github.com/sirupsen/logrus"
)

var getNextIDFunc = pkg.GetNextId

// AddUser add a user
func (i impl) AddUser(ctx context.Context, input models.User) (models.User, error) {
	userF, _ := i.userRepo.FindUserByEmail(ctx, input.Email)
	if (userF.Email) == input.Email {
		return userF, nil
	}

	ID, err := getNextIDFunc()
	if err != nil {
		log.Printf("error when generate ID %v ", err)
		return models.User{}, err
	}
	input.ID = ID

	user, errI := i.userRepo.AddUser(ctx, input)
	if errI != nil {
		log.Printf("error when add a card: %+v", input)
		return models.User{}, err
	}
	return user, nil
}
