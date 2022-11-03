package relationship

import (
	"context"
	"fm/api/internal/models"
)

// FindFriendList find a list of friends
func (i impl) FindFriendList(ctx context.Context, id int64) ([]models.Relationship, error) {
	var lists []models.Relationship
	tx := i.gormDB.Select("relationships.*").Where("requester_id = ?", id).Where("type = ?", "Friend").Find(&lists)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return lists, nil
}
