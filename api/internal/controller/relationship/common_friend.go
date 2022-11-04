package relationship

import (
	"context"
	"errors"
	"fm/api/internal/models"
	"fm/api/internal/pkg"
	log "github.com/sirupsen/logrus"
)

// CommonFriend retrieve common friend list
func (i impl) CommonFriend(ctx context.Context, requesterEmail string, addresseeEmail string) ([]models.User, error) {
	//check email valid
	errC := pkg.CheckValidEmail(requesterEmail)
	errC = pkg.CheckValidEmail(addresseeEmail)
	if errC != nil {
		log.Printf("error when check valid email %v,", errC)
		return nil, errC
	}
	//check email
	emailRequester, err := i.userRepo.FindUserByEmail(ctx, requesterEmail)
	if err != nil {
		log.Printf("error when find email %v ", err)
		return nil, errors.New("can't find requester email")
	}
	emailAddressee, err := i.userRepo.FindUserByEmail(ctx, addresseeEmail)
	if err != nil {
		log.Printf("error when find email, %v ", err)
		return nil, errors.New("can't find addressee email")
	}
	listOfRequester, errR := i.relationshipRepo.FindFriendList(ctx, emailRequester.ID)
	listOfAddressee, errR := i.relationshipRepo.FindFriendList(ctx, emailAddressee.ID)
	if errR != nil {
		log.Printf("error when find relationship, %v", errR)
		return nil, errors.New("can't find your common friend dude")
	}

	//check the matching record of 2 list 
	if len(listOfRequester) == 0 || len(listOfAddressee) == 0 {
		return nil, nil
	}
	var listUser []models.User
	for _, s := range listOfRequester {
		for _, k := range listOfAddressee {
			if s.ID == k.ID {
				listUser = append(listUser, s)
			}
		}
	}
	return listUser, nil
}
