package relationship

import (
	"context"
	"errors"
	"fm/api/internal/models"
	log "github.com/sirupsen/logrus"
)

// CommonFriend retrieve common friend list
func (i impl) CommonFriend(ctx context.Context, requesterEmail string, addresseeEmail string) ([]models.User, error) {
	//check email valid
	user1, err2 := CheckValidAndFindUser(ctx, requesterEmail, i)
	if err2 != nil {
		return nil, err2
	}
	user2, err3 := CheckValidAndFindUser(ctx, addresseeEmail, i)
	if err3 != nil {
		return nil, err3
	}
	listOfUser1, errR := i.relationshipRepo.FindFriendList(ctx, user1.ID)
	listOfUser2, errR := i.relationshipRepo.FindFriendList(ctx, user2.ID)
	if errR != nil {
		log.Printf("error when find relationship, %v", errR)
		return nil, errors.New("can't find your common friend dude")
	}

	//check the matching record of 2 list
	if len(listOfUser1) == 0 || len(listOfUser2) == 0 {
		return nil, nil
	}
	var list []models.User
	for _, s := range listOfUser1 {
		for _, k := range listOfUser2 {
			if s.ID == k.ID {
				list = append(list, s)
			}
		}
	}
	return list, nil
}
