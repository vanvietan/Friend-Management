package relationship

import (
	"context"
	"errors"
	"fm/api/internal/models"
	"fm/api/internal/pkg"
	log "github.com/sirupsen/logrus"
)

// Subscribe subscribe a target from requester
func (i impl) Subscribe(ctx context.Context, requesterEmail string, addresseeEmail string) error {
	//check email valid
	errC := pkg.CheckValidEmail(requesterEmail)
	errC = pkg.CheckValidEmail(addresseeEmail)
	if errC != nil {
		log.Printf("error when check valid email %v,", errC)
		return errC
	}
	//check email
	emailRequester, err := i.userRepo.FindUserByEmail(ctx, requesterEmail)
	if err != nil {
		log.Printf("error when find email %v ", err)
		return errors.New("can't find requester email")
	}
	emailAddressee, err := i.userRepo.FindUserByEmail(ctx, addresseeEmail)
	if err != nil {
		log.Printf("error when find email, %v ", err)
		return errors.New("can't find addressee email")
	}
	//check relationship
	rela, _ := i.relationshipRepo.FindRelationshipWithTwoEmail(ctx, emailRequester.ID, emailAddressee.ID)
	if rela.Type == models.TypeBlocked {
		return errors.New("requester is blocked")
	}
	if rela.Type == models.TypeFriend {
		return errors.New("addressee is your friend")
	}
	//create relationship
	var relationship models.Relationship
	ID, errG := getNextIDFunc()
	if errG != nil {
		log.Printf("error when generate ID %v ", errG)
		return errG
	}
	relationship.ID = ID
	relationship.RequesterID = emailRequester.ID
	relationship.AddresseeID = emailAddressee.ID
	relationship.Type = models.TypeSubscribed

	_, errR := i.relationshipRepo.CreateRelationship(ctx, relationship)
	if errR != nil {
		log.Printf("error when subscribe %v ", errR)
		return errR
	}
	return nil
}
