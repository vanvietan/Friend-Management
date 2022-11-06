package relationship

import (
	"context"
	"fm/api/internal/models"
)

// FindRelationshipWithTwoEmail find a relationship between 2 emails
func (i impl) FindRelationshipWithTwoEmail(ctx context.Context, requesterID int64, addresseeID int64) (models.Relationship, error) {
	var relationship models.Relationship
	tx := i.gormDB.Select("relationships.*").Where("requester_id = ?", requesterID).Where("addressee_id = ?", addresseeID).First(&relationship)
	if tx.Error != nil {
		return models.Relationship{}, tx.Error
	}
	return relationship, nil
}
