package relationship

import (
	"context"
	"fm/api/internal/models"
)

// FindNotificationList get a list of eligible receiving notifications from sender
func (i impl) FindNotificationList(ctx context.Context, id int64) ([]models.User, error) {
	var list []models.User
	tx := i.gormDB.Raw(`select  public.users.* from (select relationships.requester_id from public.relationships where addressee_id = ? and type != 'Blocked') u left join public.users on public.users.id = u.requester_id`, id).Find(&list)
	if tx.Error != nil {
		return nil, tx.Error
	}
	tx.Commit()
	return list, nil
}
