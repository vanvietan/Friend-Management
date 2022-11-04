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
	user1, err3 := CheckValidAndFindUser(ctx, requesterEmail, i)
	if err3 != nil {
		return err3
	}
	user2, err4 := CheckValidAndFindUser(ctx, addresseeEmail, i)
	if err4 != nil {
		return err4
	}

	//check relationship
	rela, _ := i.relationshipRepo.FindRelationshipWithTwoEmail(ctx, user1.ID, user2.ID)
	if rela.Type == models.TypeBlocked {
		return nil
	}
	//block if type is subscribe
	if rela.Type == models.TypeSubscribed {
		rela.Type = models.TypeBlocked
		_, errT := i.relationshipRepo.UpdateRelationship(ctx, rela)
		if errT != nil {
			log.Printf("error when update new relationship %v ", errT)
			return errT
		}
		err5 := createBlockRelationship(ctx, user2, user1, i)
		if err5 != nil {
			return err5
		}
		return nil
	}

	//block if type is friend
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

	//create relationship1
	err2 := createBlockRelationship(ctx, user1, user2, i)
	err2 = createBlockRelationship(ctx, user2, user1, i)
	if err2 != nil {
		return err2
	}
	return nil
}

func CheckValidAndFindUser(ctx context.Context, email string, i impl) (models.User, error) {
	errC := pkg.CheckValidEmail(email)
	if errC != nil {
		log.Printf("error when check valid email %v,", errC)
		return models.User{}, errC
	}
	//check email
	user, err := i.userRepo.FindUserByEmail(ctx, email)
	if err != nil {
		log.Printf("error when find email %v ", err)
		return models.User{}, errors.New("can't find " + email + " email")
	}
	return user, nil
}

func createBlockRelationship(ctx context.Context, user1 models.User, user2 models.User, i impl) error {
	var relationship models.Relationship
	ID, errG := getNextIDFunc()
	if errG != nil {
		log.Printf("error when generate ID %v ", errG)
		return errG
	}
	relationship.ID = ID
	relationship.RequesterID = user1.ID
	relationship.AddresseeID = user2.ID
	relationship.Type = models.TypeBlocked

	_, errR := i.relationshipRepo.CreateRelationship(ctx, relationship)
	if errR != nil {
		log.Printf("error when create new block %v ", errR)
		return errR
	}
	return nil
}
