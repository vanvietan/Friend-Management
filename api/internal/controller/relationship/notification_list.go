package relationship

import (
	"context"
	"errors"
	"fm/api/internal/models"
	"fm/api/internal/pkg"
	log "github.com/sirupsen/logrus"
	"strings"
)

// NotificationList get notification list
func (i impl) NotificationList(ctx context.Context, sender string, text string) ([]models.User, error) {
	user, err := CheckValidAndFindUser(ctx, sender, i)
	if err != nil {
		return nil, err
	}

	list, errR := i.relationshipRepo.FindNotificationList(ctx, user.ID)
	if errR != nil {
		log.Printf("error when find relationship, %v", errR)
		return nil, errors.New("can't find your list dude")
	}
	// if being mentioned in the update
	var mentionList []models.User
	textArr := strings.Fields(text)
	for _, t := range textArr {
		if pkg.CheckValidEmail(t) == nil {
			u, err2 := i.userRepo.FindUserByEmail(ctx, t)
			if err2 != nil {
				log.Printf("error when find email %v ", err)
				return nil, errors.New("can't find " + t + " email")
			}
			rela, _ := i.relationshipRepo.FindRelationshipWithTwoEmail(ctx, user.ID, u.ID)
			if rela.Type != models.TypeBlocked && u.ID != user.ID {
				mentionList = append(mentionList, u)
			}
		}
	}
	notiList := appendUniqueSlice(list, mentionList)

	return notiList, nil
}
func appendUniqueSlice(slice []models.User, mention []models.User) []models.User {
	for _, ele := range mention {
		for _, m := range slice {
			if ele.ID == m.ID {
				return slice
			}
		}
		slice = append(slice, ele)
	}
	return slice
}
