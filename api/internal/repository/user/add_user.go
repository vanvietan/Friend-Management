package user

import (
	"context"
	"fm/api/internal/models"
)

// AddUser add a user
func (i impl) AddUser(ctx context.Context, user models.User) (models.User, error) {
	tx := i.gormDB.Create(&user)
	if tx.Error != nil {
		return models.User{}, tx.Error
	}
	return user, nil
}
