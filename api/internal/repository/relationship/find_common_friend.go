package relationship

import (
	"context"
	"fm/api/internal/models"
)

func (i impl) FindCommonFriend(ctx context.Context, id1 int64, id2 int64) ([]models.Relationship, error) {
	var lists []models.Relationship
	tx := i.gormDB.Raw(`SELECT relationships.* FROM public.relationships WHERE requester_id = ? AND type = 'Friend'
							UNION SELECT relationships.* FROM public.relationships WHERE requester_id = ? AND type = 'Friend'`, id1, id2).Find(&lists)
	//tx := i.gormDB.Select("relationships.*").Where("requester_id = ?", id1).Where("requester_id = ?", id2).Where("type = ?", "Friend").Find(&lists)
	//select public.users.id, public.users.email from (SELECT relationships.* FROM public.relationships WHERE requester_id = 101 and type = 'Friend' union SELECT relationships.* FROM public.relationships WHERE requester_id = 105 and type = 'Friend') u
	//	inner join public.users on  public.users.id = u.addressee_id;

	if tx.Error != nil {
		return nil, tx.Error
	}
	tx.Commit()
	return lists, nil
}
