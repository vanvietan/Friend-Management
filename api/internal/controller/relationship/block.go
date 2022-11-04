package relationship

import (
	"context"
	"errors"
	"fm/api/internal/models"
	"fm/api/internal/pkg"
	"fmt"
	log "github.com/sirupsen/logrus"
)

// Block to block target by requester
func (i impl) Block(ctx context.Context, requesterEmail string, addresseeEmail string) error {
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
		return nil
	}

	if rela.Type == models.TypeSubscribed {
		rela.Type = models.TypeBlocked
		_, errT := i.relationshipRepo.UpdateRelationship(ctx, rela)
		if errT != nil {
			log.Printf("error when update new relationship %v ", errT)
			return errT
		}
		return nil
	}

	if rela.Type == models.TypeFriend {
		rela.Type = models.TypeBlocked
		_, errT := i.relationshipRepo.UpdateRelationship(ctx, rela)
		//block another side of friendship,
		rela2, _ := i.relationshipRepo.FindRelationshipWithTwoEmail(ctx, rela.AddresseeID, rela.RequesterID)
		fmt.Println(rela2)
		rela2.Type = models.TypeBlocked
		_, errT = i.relationshipRepo.UpdateRelationship(ctx, rela2)
		if errT != nil {
			log.Printf("error when update new relationship %v ", errT)
			return errT
		}
		return nil
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
	relationship.Type = models.TypeBlocked

	_, errR := i.relationshipRepo.CreateRelationship(ctx, relationship)
	if errR != nil {
		log.Printf("error when create new block %v ", errR)
		return errR
	}
	return nil
}
