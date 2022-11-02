package relationship

import (
	"context"
	"fm/api/internal/models"
)

// CreateRelationship create a relationship
func (i impl) CreateRelationship(ctx context.Context, relationship models.Relationship) (models.Relationship, error) {
	//tx := i.gormDB.Create(&relationship)
	tx := i.gormDB.Exec(`INSERT INTO public.relationships(addressee_id,requester_id,type,id) VALUES(?,?,?,?)`, relationship.AddresseeID, relationship.RequesterID, relationship.Type, relationship.ID)
	if tx.Error != nil {
		return models.Relationship{}, tx.Error
	}
	tx.Commit()
	return relationship, nil
}
