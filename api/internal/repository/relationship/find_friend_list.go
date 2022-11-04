package relationship

import (
	"context"
	"fm/api/internal/models"
)

// FindFriendList find a list of friends
func (i impl) FindFriendList(ctx context.Context, id int64) ([]models.User, error) {
	var lists []models.User
	tx := i.gormDB.Raw(`select public.users.id, public.users.email from (SELECT relationships.* FROM public.relationships WHERE requester_id = ? and type = 'Friend') u inner join public.users on  public.users.id = u.addressee_id`, id).Find(&lists)
	//tx := i.gormDB.Select("relationships.*").Where("requester_id = ?", id).Where("type = ?", "Friend").Find(&lists)
	if tx.Error != nil {
		return nil, tx.Error
	}
	tx.Commit()
	return lists, nil
}
