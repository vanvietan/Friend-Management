package relationship

import (
	"context"
	"errors"
	"fm/api/internal/models"
	log "github.com/sirupsen/logrus"
)

// Subscribe subscribe a target from requester
func (i impl) Subscribe(ctx context.Context, requesterEmail string, addresseeEmail string) error {
	//can't subscribe to self
	if requesterEmail == addresseeEmail {
		return errors.New("can't subscribe to yourself")
	}
	//check email valid
	user1, err2 := CheckValidAndFindUser(ctx, requesterEmail, i)
	if err2 != nil {
		return err2
	}
	user2, err3 := CheckValidAndFindUser(ctx, addresseeEmail, i)
	if err3 != nil {
		return err3
	}

	//check relationship
	rela, _ := i.relationshipRepo.FindRelationshipWithTwoEmail(ctx, user1.ID, user2.ID)
	if rela.Type == models.TypeBlocked {
		return errors.New(requesterEmail + " is blocked")
	}
	if rela.Type == models.TypeFriend {
		return errors.New(addresseeEmail + " is already your friend")
	}
	//create relationship
	var relationship models.Relationship
	ID, errG := getNextIDFunc()
	if errG != nil {
		log.Printf("error when generate ID %v ", errG)
		return errG
	}
	relationship.ID = ID
	relationship.RequesterID = user1.ID
	relationship.AddresseeID = user2.ID
	relationship.Type = models.TypeSubscribed

	_, errR := i.relationshipRepo.CreateRelationship(ctx, relationship)
	if errR != nil {
		log.Printf("error when subscribe %v ", errR)
		return errR
	}
	return nil
}
