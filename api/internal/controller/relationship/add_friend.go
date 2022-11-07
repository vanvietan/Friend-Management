package relationship

import (
	"context"
	"errors"
	"fm/api/internal/models"
	"fm/api/internal/pkg"
	log "github.com/sirupsen/logrus"
)

var getNextIDFunc = pkg.GetNextId

// AddFriend add friend controller
func (i impl) AddFriend(ctx context.Context, requesterEmail string, addresseeEmail string) error {
	//can't add friend to self
	if requesterEmail == addresseeEmail {
		return errors.New("can't add friend to yourself")
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
		return errors.New("requester is blocked")
	}
	if rela.Type == models.TypeFriend {
		return nil
	}

	err4 := createFriendRelationship(ctx, user1, user2, i)
	if err4 != nil {
		return err4
	}
	err5 := createFriendRelationship(ctx, user2, user1, i)
	if err5 != nil {
		return err5
	}
	return nil
}

func createFriendRelationship(ctx context.Context, user1 models.User, user2 models.User, i impl) error {
	var relationship models.Relationship
	ID, errG := getNextIDFunc()
	if errG != nil {
		log.Printf("error when generate ID %v ", errG)
		return errG
	}
	relationship.ID = ID
	relationship.RequesterID = user1.ID
	relationship.AddresseeID = user2.ID
	relationship.Type = models.TypeFriend

	_, errR := i.relationshipRepo.CreateRelationship(ctx, relationship)
	if errR != nil {
		log.Printf("error when add friend %v ", errR)
		return errR
	}
	return nil
}
