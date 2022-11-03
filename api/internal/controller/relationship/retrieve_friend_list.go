package relationship

import (
	"context"
	"errors"
	"fm/api/internal/models"
	"fm/api/internal/pkg"
	log "github.com/sirupsen/logrus"
)

// FriendList retrieve a list of friend given by an email
func (i impl) FriendList(ctx context.Context, input string) ([]models.User, error) {
	//check email
	errC := pkg.CheckValidEmail(input)
	if errC != nil {
		return nil, errC
	}
	user, errU := i.userRepo.FindUserByEmail(ctx, input)
	if errU != nil {
		log.Printf("error when find email, %v ", errU)
		return nil, errors.New("cant find requester email")
	}
	relationships, errR := i.relationshipRepo.FindFriendList(ctx, user.ID)
	if errR != nil {
		log.Printf("error when find relationship, %v", errR)
		return nil, errors.New("can't find your friend dude")
	}
	var listUser []models.User

	for _, s := range relationships {
		u, errF := i.userRepo.FindUserByID(ctx, s.AddresseeID)
		if errF != nil {
			log.Printf("error when find user, %v", errF)
			return nil, errors.New("can't find your friend dude")
		}
		listUser = append(listUser, u)
	}

	return listUser, nil
}