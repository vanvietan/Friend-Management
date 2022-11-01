package relationship

import (
	"context"
	"fm/api/internal/models"
)

// CreateRelationship create a relationship
func (i impl) CreateRelationship(ctx context.Context, relationship models.Relationship) (models.Relationship, error) {
	tx := i.gormDB.Create(&relationship)
	if tx.Error != nil {
		return models.Relationship{}, tx.Error
	}
	return relationship, nil
}
