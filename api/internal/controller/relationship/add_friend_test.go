package relationship

import (
	"context"
	"errors"
	mocksR "fm/api/internal/mocks/repository/relationship"
	mocksU "fm/api/internal/mocks/repository/user"
	"fm/api/internal/models"
	"fm/api/internal/pkg"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestAddFriend(t *testing.T) {
	type addFriend struct {
		mockRequesterEmail         string
		mockAddresseeEmail         string
		mockOutUser1               models.User
		mockOutUser2               models.User
		mockInCreateRelationship1  models.Relationship
		mockInCreateRelationship2  models.Relationship
		mockOutFindRelationship    models.Relationship
		mockOutCreateRelationship1 models.Relationship
		mockOutCreateRelationship2 models.Relationship
		mockErrFind                error
		mockErrFindRelationship    error
		mockErrCreate              error
	}
	type arg struct {
		addFriend           addFriend
		givenRequesterEmail string
		givenAddresseeEmail string
		expErr              error
	}
	tcs := map[string]arg{
		"success": {
			addFriend: addFriend{
				mockRequesterEmail: "van1@gmail.com",
				mockAddresseeEmail: "van5@gmail.com",
				mockOutUser1: models.User{
					ID:    101,
					Email: "van1@gmail.com",
				},
				mockOutUser2: models.User{
					ID:    105,
					Email: "van5@gmail.com",
				},
				mockInCreateRelationship1: models.Relationship{
					ID:          1,
					AddresseeID: 105,
					RequesterID: 101,
					Type:        "Friend",
				},
				mockInCreateRelationship2: models.Relationship{
					ID:          1,
					AddresseeID: 101,
					RequesterID: 105,
					Type:        "Friend",
				},
				mockOutFindRelationship: models.Relationship{},
				mockOutCreateRelationship1: models.Relationship{
					ID:          1,
					AddresseeID: 105,
					RequesterID: 101,
					Type:        "Friend",
				},
				mockOutCreateRelationship2: models.Relationship{
					ID:          1,
					AddresseeID: 101,
					RequesterID: 105,
					Type:        "Friend",
				},
				mockErrFindRelationship: errors.New("record not found"),
			},
			givenRequesterEmail: "van1@gmail.com",
			givenAddresseeEmail: "van5@gmail.com",
		},
		"fail: self add friend": {
			givenRequesterEmail: "van1@gmail.com",
			givenAddresseeEmail: "van1@gmail.com",
			expErr:              errors.New("can't add friend to yourself"),
		},
		"fail: invalid email address": {
			addFriend: addFriend{
				mockRequesterEmail: "@gmail.com",
				mockAddresseeEmail: "van5@gmail.com",
				mockOutUser1:       models.User{},
				mockOutUser2: models.User{
					ID:    105,
					Email: "van5@gmail.com",
				},
			},
			givenRequesterEmail: "@gmail.com",
			givenAddresseeEmail: "van5@gmail.com",
			expErr:              errors.New("invalid email address"),
		},
		"fail: record not found": {
			addFriend: addFriend{
				mockRequesterEmail: "van10@gmail.com",
				mockAddresseeEmail: "van5@gmail.com",
				mockOutUser1:       models.User{},
				mockOutUser2: models.User{
					ID:    105,
					Email: "van5@gmail.com",
				},
				mockErrFind: errors.New("record not found"),
			},
			givenRequesterEmail: "van10@gmail.com",
			givenAddresseeEmail: "van5@gmail.com",
			expErr:              errors.New("can't find van10@gmail.com email"),
		},
		"fail: requester is blocked": {
			addFriend: addFriend{
				mockRequesterEmail: "van1@gmail.com",
				mockAddresseeEmail: "van5@gmail.com",
				mockOutUser1: models.User{
					ID:    101,
					Email: "van1@gmail.com",
				},
				mockOutUser2: models.User{
					ID:    105,
					Email: "van5@gmail.com",
				},
				mockOutFindRelationship: models.Relationship{
					ID:          1,
					AddresseeID: 101,
					RequesterID: 105,
					Type:        "Blocked",
				},
				mockErrFindRelationship: errors.New("requester is blocked"),
			},
			givenRequesterEmail: "van1@gmail.com",
			givenAddresseeEmail: "van5@gmail.com",
			expErr:              errors.New("requester is blocked"),
		},
		"fail: error from repo": {
			addFriend: addFriend{
				mockRequesterEmail: "van1@gmail.com",
				mockAddresseeEmail: "van5@gmail.com",
				mockOutUser1: models.User{
					ID:    101,
					Email: "van1@gmail.com",
				},
				mockOutUser2: models.User{
					ID:    105,
					Email: "van5@gmail.com",
				},
				mockInCreateRelationship1: models.Relationship{
					ID:          1,
					AddresseeID: 105,
					RequesterID: 101,
					Type:        "Friend",
				},
				mockInCreateRelationship2: models.Relationship{
					ID:          1,
					AddresseeID: 101,
					RequesterID: 105,
					Type:        "Friend",
				},
				mockOutFindRelationship: models.Relationship{},
				mockOutCreateRelationship1: models.Relationship{
					ID:          1,
					AddresseeID: 105,
					RequesterID: 101,
					Type:        "Friend",
				},
				mockOutCreateRelationship2: models.Relationship{
					ID:          1,
					AddresseeID: 101,
					RequesterID: 105,
					Type:        "Friend",
				},
				mockErrFindRelationship: errors.New("record not found"),
				mockErrCreate:           errors.New("something wrong"),
			},
			givenRequesterEmail: "van1@gmail.com",
			givenAddresseeEmail: "van5@gmail.com",
			expErr:              errors.New("something wrong"),
		},
	}
	for s, tc := range tcs {
		t.Run(s, func(t *testing.T) {
			//GIVEN
			relaRepo := new(mocksR.Repository)
			userRepo := new(mocksU.Repository)
			userRepo.On("FindUserByEmail", mock.Anything, tc.addFriend.mockRequesterEmail).
				Return(tc.addFriend.mockOutUser1, tc.addFriend.mockErrFind)
			userRepo.On("FindUserByEmail", mock.Anything, tc.addFriend.mockAddresseeEmail).
				Return(tc.addFriend.mockOutUser2, tc.addFriend.mockErrFind)
			relaRepo.On("FindRelationshipWithTwoEmail", mock.Anything, tc.addFriend.mockOutUser1.ID, tc.addFriend.mockOutUser2.ID).
				Return(tc.addFriend.mockOutFindRelationship, tc.addFriend.mockErrFindRelationship)
			relaRepo.On("CreateRelationship", mock.Anything, tc.addFriend.mockInCreateRelationship1).
				Return(tc.addFriend.mockOutCreateRelationship1, tc.addFriend.mockErrCreate)
			relaRepo.On("CreateRelationship", mock.Anything, tc.addFriend.mockInCreateRelationship2).
				Return(tc.addFriend.mockOutCreateRelationship2, tc.addFriend.mockErrCreate)
			getNextIDFunc = func() (int64, error) {
				if s == "fail: generate id fail" {
					return 0, errors.New("something wrong")
				}
				return 1, nil
			}
			defer func() {
				getNextIDFunc = pkg.GetNextId
			}()
			//WHEN
			svc := New(relaRepo, userRepo)
			err := svc.AddFriend(context.Background(), tc.givenRequesterEmail, tc.givenAddresseeEmail)

			//THEN
			if tc.expErr != nil {
				require.EqualError(t, err, tc.expErr.Error())
			} else {
				require.NoError(t, err)
			}
		})
	}
}
