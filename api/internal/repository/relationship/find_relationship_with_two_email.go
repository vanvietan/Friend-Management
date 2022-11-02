package relationship

import (
	"context"
	"fm/api/internal/models"
)

func (i impl) FindRelationshipWithTwoEmail(ctx context.Context, requesterID int64, addresseeID int64) (models.Relationship, error) {
	var relationship models.Relationship
	tx := i.gormDB.Select("relationships.*").Where("requester_id = ?", requesterID).Where("addressee_id = ?", addresseeID).First(&relationship)
	//tx := i.gormDB.Exec("SELECT relationships.* FROM public.relationships WHERE requester_id = $1 AND addressee_id = $2", requesterID, addresseeID).Scan(&relationship)
	if tx.Error != nil {
		return models.Relationship{}, tx.Error
	}
	//tx.Commit()
	return relationship, nil
}
