package user

import (
	"gorm.io/gorm"
)

// Repository contains all user repository functions
type Repository interface {
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
