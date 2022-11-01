package relationship

import (
	"context"
	"fm/api/internal/models"
	"fm/api/internal/pkg"
	log "github.com/sirupsen/logrus"
)

var getNextIDFunc = pkg.GetNextId

func (i impl) AddFriend(ctx context.Context, requesterID int64, addresseeID int64) error {
	var relationship models.Relationship
	ID, err := getNextIDFunc()
	if err != nil {
		log.Printf("error when generate ID %v ", err)
		return err
	}
	relationship.ID = ID
	relationship.RequesterID = requesterID
	relationship.AddresseeID = addresseeID
	relationship.Type = models.TypeFriend

	_, errR := i.relationshipRepo.CreateRelationship(ctx, relationship)
	if errR != nil {
		log.Printf("error when add friend %v ", err)
		return errR
	}
	return nil
}
