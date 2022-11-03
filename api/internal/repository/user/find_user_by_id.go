package user

import (
	"context"
	"fm/api/internal/models"
)

// FindUserByID find user by ID
func (i impl) FindUserByID(ctx context.Context, id int64) (models.User, error) {
	user := models.User{}
	tx := i.gormDB.First(&user, id)
	if tx.Error != nil {
		return models.User{}, nil
	}
	return user, nil
}
