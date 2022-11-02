package relationship

import (
	"context"
	"fm/api/internal/models"
)

// CreateRelationship create a relationship
func (i impl) CreateRelationship(ctx context.Context, relationship models.Relationship) (models.Relationship, error) {
	//tx := i.gormDB.Create(&relationship)
	tx := i.gormDB.Exec(`INSERT INTO public.relationship(addressee_id,requester_id,type,created_at,updated_at,id) VALUES(?,?,?,?,?,?)`, relationship.AddresseeID, relationship.RequesterID, relationship.Type, relationship.CreatedAt, relationship.UpdatedAt, relationship.ID)
	if tx.Error != nil {
		return models.Relationship{}, tx.Error
	}
	tx.Commit()
	return relationship, nil
}
