package relationship

import (
	"context"
	"fm/api/internal/models"
)

// UpdateRelationship update a relationship
func (i impl) UpdateRelationship(ctx context.Context, relationship models.Relationship) (models.Relationship, error) {
	tx := i.gormDB.Save(&relationship)
	if tx.Error != nil {
		return models.Relationship{}, tx.Error
	}
	return relationship, nil
}
