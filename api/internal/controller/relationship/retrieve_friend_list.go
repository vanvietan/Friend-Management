package relationship

import (
	"context"
	"errors"
	"fm/api/internal/models"
	log "github.com/sirupsen/logrus"
)

// FriendList retrieve a list of friend given by an email
func (i impl) FriendList(ctx context.Context, input string) ([]models.User, error) {
	//check email valid
	user, err2 := CheckValidAndFindUser(ctx, input, i)
	if err2 != nil {
		return nil, err2
	}
	listUser, errR := i.relationshipRepo.FindFriendList(ctx, user.ID)
	if errR != nil {
		log.Printf("error when find relationship, %v", errR)
		return nil, errors.New("can't find your friend dude")
	}
	return listUser, nil
}
