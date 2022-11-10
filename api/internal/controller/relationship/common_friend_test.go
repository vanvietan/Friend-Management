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

func TestCommonFriend(t *testing.T) {
	type commonFriend struct {
		mockInEmail1     string
		mockInEmail2     string
		mockOutUser1     models.User
		mockOutUser2     models.User
		mockOutListUser1 []models.User
		mockOutListUser2 []models.User
		mockCheckErr     error
		mockFindErr      error
	}
	type arg struct {
		commonFriend        commonFriend
		givenRequesterEmail string
		givenAddresseeEmail string
		expRs               []models.User
		expErr              error
	}
	tcs := map[string]arg{
		"success": {
			commonFriend: commonFriend{
				mockInEmail1: "van1@gmail.com",
				mockInEmail2: "van2@gmail.com",
				mockOutUser1: models.User{
					ID:    101,
					Email: "van1@gmail.com",
				},
				mockOutUser2: models.User{
					ID:    102,
					Email: "van2@gmail.com",
				},
				mockOutListUser1: []models.User{
					{
						ID:    103,
						Email: "van3@gmail.com",
					},
				},
				mockOutListUser2: []models.User{
					{
						ID:    103,
						Email: "van3@gmail.com",
					},
				},
			},
			givenRequesterEmail: "van1@gmail.com",
			givenAddresseeEmail: "van2@gmail.com",
			expRs: []models.User{
				{
					ID:    103,
					Email: "van3@gmail.com",
				},
			},
		},
		"success: empty result": {
			commonFriend: commonFriend{
				mockInEmail1: "van1@gmail.com",
				mockInEmail2: "van2@gmail.com",
				mockOutUser1: models.User{
					ID:    101,
					Email: "van1@gmail.com",
				},
				mockOutUser2: models.User{
					ID:    102,
					Email: "van2@gmail.com",
				},
				mockOutListUser1: []models.User{
					{
						ID:    103,
						Email: "van3@gmail.com",
					},
				},
				mockOutListUser2: []models.User{
					{
						ID:    104,
						Email: "van4@gmail.com",
					},
				},
			},
			givenRequesterEmail: "van1@gmail.com",
			givenAddresseeEmail: "van2@gmail.com",
			expRs:               nil,
		},
		"fail: invalid email address": {
			givenRequesterEmail: "@gmail.com",
			givenAddresseeEmail: "van2@gmail.com",
			expRs:               nil,
			expErr:              errors.New("invalid email address"),
		},
		"fail: error from repo": {
			commonFriend: commonFriend{
				mockInEmail1: "van1@gmail.com",
				mockInEmail2: "van2@gmail.com",
				mockOutUser1: models.User{
					ID:    101,
					Email: "van1@gmail.com",
				},
				mockOutUser2: models.User{
					ID:    102,
					Email: "van2@gmail.com",
				},
				mockOutListUser1: []models.User{
					{
						ID:    103,
						Email: "van3@gmail.com",
					},
				},
				mockOutListUser2: nil,
				mockFindErr:      errors.New("something wrong"),
			},
			givenRequesterEmail: "van1@gmail.com",
			givenAddresseeEmail: "van2@gmail.com",
			expRs:               nil,
			expErr:              errors.New("can't find your common friend dude"),
		},
	}
	for s, tc := range tcs {
		t.Run(s, func(t *testing.T) {
			//GIVEN
			relaRepo := new(mocksR.Repository)
			userRepo := new(mocksU.Repository)
			userRepo.On("FindUserByEmail", mock.Anything, tc.commonFriend.mockInEmail1).
				Return(tc.commonFriend.mockOutUser1, tc.commonFriend.mockCheckErr)
			userRepo.On("FindUserByEmail", mock.Anything, tc.commonFriend.mockInEmail2).
				Return(tc.commonFriend.mockOutUser2, tc.commonFriend.mockCheckErr)
			relaRepo.On("FindFriendList", mock.Anything, tc.commonFriend.mockOutUser1.ID).
				Return(tc.commonFriend.mockOutListUser1, tc.commonFriend.mockFindErr)
			relaRepo.On("FindFriendList", mock.Anything, tc.commonFriend.mockOutUser2.ID).
				Return(tc.commonFriend.mockOutListUser2, tc.commonFriend.mockFindErr)

			//WHEN
			svc := New(relaRepo, userRepo)
			list, err := svc.CommonFriend(context.Background(), tc.givenRequesterEmail, tc.givenAddresseeEmail)

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
