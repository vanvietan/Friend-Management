package relationship

import (
	"context"
	"fm/api/internal/models"
	"gorm.io/gorm"
)

// Repository contains repository of relationship
type Repository interface {
	CreateRelationship(ctx context.Context, relationship models.Relationship) (models.Relationship, error)
}
type impl struct {
	gormDB *gorm.DB
}

func New(gormDB *gorm.DB) Repository {
	return impl{
		gormDB: gormDB,
	}
}
