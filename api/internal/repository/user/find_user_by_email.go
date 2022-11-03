package user

import (
	"context"
	"fm/api/internal/models"
)

// FindUserByEmail find user by email
func (i impl) FindUserByEmail(ctx context.Context, input string) (models.User, error) {
	user := models.User{}
	tx := i.gormDB.Where("email = ?", input).First(&user)
	if tx.Error != nil {
		return models.User{}, tx.Error
	}
	return user, nil
}
