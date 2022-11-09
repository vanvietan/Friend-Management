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

func TestFriendList(t *testing.T) {
	type friendList struct {
		mockInput             string
		mockOutUser           models.User
		mockOutFindFriendList []models.User
		mockCheckErr          error
		mockFindFriendListErr error
	}
	type arg struct {
		friendList friendList
		givenInput string
		expRs      []models.User
		expErr     error
	}
	tcs := map[string]arg{
		"success": {
			friendList: friendList{
				mockInput: "van1@gmail.com",
				mockOutUser: models.User{
					ID:    101,
					Email: "van1@gmail.com",
				},
				mockOutFindFriendList: []models.User{
					{
						ID:    102,
						Email: "van2@gmail.com",
					},
					{
						ID:    103,
						Email: "van3@gmail.com",
					},
				},
			},
			givenInput: "van1@gmail.com",
			expRs: []models.User{
				{
					ID:    102,
					Email: "van2@gmail.com",
				},
				{
					ID:    103,
					Email: "van3@gmail.com",
				},
			},
		},
		"fail: error from repo ": {
			friendList: friendList{
				mockInput: "van1@gmail.com",
				mockOutUser: models.User{
					ID:    101,
					Email: "van1@gmail.com",
				},
				mockOutFindFriendList: []models.User{
					{
						ID:    102,
						Email: "van2@gmail.com",
					},
					{
						ID:    103,
						Email: "van3@gmail.com",
					},
				},
				mockFindFriendListErr: errors.New("something wrong"),
			},
			givenInput: "van1@gmail.com",
			expRs: []models.User{
				{
					ID:    102,
					Email: "van2@gmail.com",
				},
				{
					ID:    103,
					Email: "van3@gmail.com",
				},
			},
			expErr: errors.New("can't find your friend dude"),
		},
		"fail: invalid email address": {
			givenInput: "@gmail.com",
			expRs:      nil,
			expErr:     errors.New("invalid email address"),
		},
		"fail: error when find user": {
			friendList: friendList{
				mockInput: "van1@gmail.com",
				mockOutUser: models.User{
					ID:    101,
					Email: "van1@gmail.com",
				},
				mockOutFindFriendList: []models.User{
					{
						ID:    102,
						Email: "van2@gmail.com",
					},
					{
						ID:    103,
						Email: "van3@gmail.com",
					},
				},
				mockFindFriendListErr: errors.New("something wrong"),
			},
			givenInput: "van1@gmail.com",
			expRs:      nil,
			expErr:     errors.New("can't find your friend dude"),
		},
	}
	for s, tc := range tcs {
		t.Run(s, func(t *testing.T) {
			//GIVEN
			relaRepo := new(mocksR.Repository)
			userRepo := new(mocksU.Repository)
			userRepo.On("FindUserByEmail", mock.Anything, tc.friendList.mockInput).
				Return(tc.friendList.mockOutUser, tc.friendList.mockCheckErr)
			relaRepo.On("FindFriendList", mock.Anything, tc.friendList.mockOutUser.ID).
				Return(tc.friendList.mockOutFindFriendList, tc.friendList.mockFindFriendListErr)

			//WHEN
			svc := New(relaRepo, userRepo)
			list, err := svc.FriendList(context.Background(), tc.givenInput)

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
