package relationship

import "gorm.io/gorm"

// Repository contains repository of relationship
type Repository interface {
}
type impl struct {
	gormDB *gorm.DB
}

func New(gormDB *gorm.DB) Repository {
	return impl{
		gormDB: gormDB,
	}
}
