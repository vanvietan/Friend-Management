package relationship

import (
	"context"
	"fm/api/internal/models"
	"gorm.io/gorm"
)

// Repository contains repository of relationship
type Repository interface {
	//CreateRelationship create a relationship
	CreateRelationship(ctx context.Context, relationship models.Relationship) (models.Relationship, error)

	//FindRelationshipWithTwoEmail find relationship with 2 email
	FindRelationshipWithTwoEmail(ctx context.Context, requesterID int64, addresseeID int64) (models.Relationship, error)
}
type impl struct {
	gormDB *gorm.DB
}

// New DI
func New(gormDB *gorm.DB) Repository {
	return impl{
		gormDB: gormDB,
	}
}
