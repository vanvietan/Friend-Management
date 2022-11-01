package relationship

import (
	"context"
	"fm/api/internal/models"
	"fm/api/internal/pkg"
	log "github.com/sirupsen/logrus"
)

var getNextIDFunc = pkg.GetNextId

func (i impl) AddFriend(ctx context.Context, requesterEmail string, addresseeEmail string) error {

	//check email
	emailRequester, err := i.userRepo.FindUserByEmail(ctx, requesterEmail)
	if err != nil {
		log.Printf("error when find email %v ", err)
		return err
	}
	emailAddressee, err := i.userRepo.FindUserByEmail(ctx, addresseeEmail)
	if err != nil {
		log.Printf("error when find email %v ", err)
		return err
	}
	/*
		TODO: check Type isBlocked or not
	*/

	var relationship models.Relationship
	ID, errG := getNextIDFunc()
	if errG != nil {
		log.Printf("error when generate ID %v ", errG)
		return errG
	}

	relationship.ID = ID
	relationship.RequesterID = emailRequester.ID
	relationship.AddresseeID = emailAddressee.ID
	relationship.Type = models.TypeFriend

	_, errR := i.relationshipRepo.CreateRelationship(ctx, relationship)
	if errR != nil {
		log.Printf("error when add friend %v ", errR)
		return errR
	}
	return nil
}
