package relationship

import (
	"context"
	"fm/api/internal/models"
	"fm/api/internal/repository/relationship"
	"fm/api/internal/repository/user"
)

// Service contains relationship services
type Service interface {
	//AddFriend Add friends between 2 emails
	AddFriend(ctx context.Context, requesterEmail string, addresseeEmail string) error

	//FriendList retrieve friend list
	FriendList(ctx context.Context, input string) ([]models.User, error)

	//CommonFriend retrieve common friend list
	CommonFriend(ctx context.Context, requesterEmail string, addresseeEmail string) ([]models.User, error)

	//Subscribe subscribe a target from requester
	Subscribe(ctx context.Context, requester string, addressee string) error

	//Block to block a target from requester
	Block(ctx context.Context, requester string, addressee string) error

	//NotificationList get notification list
	NotificationList(ctx context.Context, sender string, text string) ([]models.User, error)
}

type impl struct {
	relationshipRepo relationship.Repository
	userRepo         user.Repository
}

// New DI
func New(relationshipRepo relationship.Repository, userRepo user.Repository) Service {
	return impl{
		relationshipRepo: relationshipRepo,
		userRepo:         userRepo,
	}
}
