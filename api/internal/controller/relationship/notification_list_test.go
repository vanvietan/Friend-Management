package relationship

import (
	"context"
	"errors"
	mocksR "fm/api/internal/mocks/repository/relationship"
	mocksU "fm/api/internal/mocks/repository/user"
	"fm/api/internal/models"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNotificationList(t *testing.T) {
	type notificationList struct {
		mockInSender                string
		mockInText                  string
		mockOutSender               models.User
		mockOutUserFromText         models.User
		mockOutFindNotificationList []models.User
		mockRelationship            models.Relationship
		mockCheckErr                error
		mockFindNotificationListErr error
		mockRelationshipErr         error
	}
	type arg struct {
		notificationList notificationList
		givenSender      string
		givenText        string
		expRs            []models.User
		expErr           error
	}
	tcs := map[string]arg{
		"success": {
			notificationList: notificationList{
				mockInSender: "van1@gmail.com",
				mockInText:   "van5@gmail.com",
				mockOutSender: models.User{
					ID:    101,
					Email: "van1@gmail.com",
				},
				mockOutUserFromText: models.User{
					ID:    105,
					Email: "van5@gmail.com",
				},
				mockOutFindNotificationList: []models.User{
					{
						ID:    102,
						Email: "van2@gmail.com",
					},
					{
						ID:    103,
						Email: "van3@gmail.com",
					},
				},
				mockRelationship: models.Relationship{},
			},
			givenSender: "van1@gmail.com",
			givenText:   "Hello World! van5@gmail.com",
			expRs: []models.User{
				{
					ID:    102,
					Email: "van2@gmail.com",
				},
				{
					ID:    103,
					Email: "van3@gmail.com",
				},
				{
					ID:    105,
					Email: "van5@gmail.com",
				},
			},
		},
		"success: empty result": {
			notificationList: notificationList{
				mockInSender: "van1@gmail.com",
				mockInText:   "van5@gmail.com",
				mockOutSender: models.User{
					ID:    101,
					Email: "van1@gmail.com",
				},
				mockOutUserFromText: models.User{
					ID:    105,
					Email: "van5@gmail.com",
				},
				mockOutFindNotificationList: nil,
				mockRelationship: models.Relationship{
					ID:          1,
					AddresseeID: 101,
					RequesterID: 105,
					Type:        "Blocked",
				},
			},
			givenSender: "van1@gmail.com",
			givenText:   "Hello World! van5@gmail.com",
			expRs:       nil,
		},
		"fail: invalid email address": {
			notificationList: notificationList{},
			givenSender:      "@gmail.com",
			givenText:        "HelloWorld",
			expRs:            nil,
			expErr:           errors.New("invalid email address"),
		},
		"fail: error from repo": {
			notificationList: notificationList{
				mockInSender: "van1@gmail.com",
				mockInText:   "van5@gmail.com",
				mockOutSender: models.User{
					ID:    101,
					Email: "van1@gmail.com",
				},
				mockOutUserFromText: models.User{
					ID:    105,
					Email: "van5@gmail.com",
				},
				mockOutFindNotificationList: nil,
				mockFindNotificationListErr: errors.New("something wrong"),
				mockRelationship:            models.Relationship{},
			},
			givenSender: "van1@gmail.com",
			givenText:   "Hello World! van5@gmail.com",
			expRs: []models.User{
				{
					ID:    102,
					Email: "van2@gmail.com",
				},
				{
					ID:    103,
					Email: "van3@gmail.com",
				},
				{
					ID:    105,
					Email: "van5@gmail.com",
				},
			},
			expErr: errors.New("can't find your list dude"),
		},
	}
	for s, tc := range tcs {
		t.Run(s, func(t *testing.T) {
			//GIVEN
			relaRepo := new(mocksR.Repository)
			userRepo := new(mocksU.Repository)
			userRepo.On("FindUserByEmail", mock.Anything, tc.notificationList.mockInSender).
				Return(tc.notificationList.mockOutSender, tc.notificationList.mockCheckErr)
			relaRepo.On("FindNotificationList", mock.Anything, tc.notificationList.mockOutSender.ID).
				Return(tc.notificationList.mockOutFindNotificationList, tc.notificationList.mockFindNotificationListErr)
			userRepo.On("FindUserByEmail", mock.Anything, tc.notificationList.mockInText).
				Return(tc.notificationList.mockOutUserFromText, tc.notificationList.mockCheckErr)
			relaRepo.On("FindRelationshipWithTwoEmail", mock.Anything, tc.notificationList.mockOutSender.ID, tc.notificationList.mockOutUserFromText.ID).
				Return(tc.notificationList.mockRelationship, tc.notificationList.mockRelationshipErr)
			//WHEN
			svc := New(relaRepo, userRepo)
			list, err := svc.NotificationList(context.Background(), tc.givenSender, tc.givenText)

			//THEN
			if tc.expErr != nil {
				require.EqualError(t, err, tc.expErr.Error())
			} else {
				require.NoError(t, err)
				require.Equal(t, tc.expRs, list)
			}
		})
	}
}
