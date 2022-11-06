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

	//FindFriendList find a list of relationship friend
	FindFriendList(ctx context.Context, id int64) ([]models.User, error)

	//UpdateRelationship update a relationship
	UpdateRelationship(ctx context.Context, relationship models.Relationship) (models.Relationship, error)

	//FindNotificationList get a list of eligible receiving notifications from sender
	FindNotificationList(ctx context.Context, id int64) ([]models.User, error)
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
