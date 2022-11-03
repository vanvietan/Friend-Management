package user

import (
	"context"
	"fm/api/internal/models"
	"gorm.io/gorm"
)

// Repository contains all user repository functions
type Repository interface {
	//FindUserByEmail find user by email
	FindUserByEmail(ctx context.Context, input string) (models.User, error)

	//FindUserByID find user by its id
	FindUserByID(ctx context.Context, id int64) (models.User, error)
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
