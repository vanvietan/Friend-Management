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

func TestSubscribe(t *testing.T) {
	type testSubscribe struct {
		mockInRequesterEmail      string
		mockInAddresseeEmail      string
		mockOutRequester          models.User
		mockOutAddressee          models.User
		mockCheckErr              error
		mockFindRelationship      models.Relationship
		mockFindRelationshipErr   error
		mockInCreateRelationship  models.Relationship
		mockOutCreateRelationship models.Relationship
		mockCreateRelationshipErr error
	}
	type arg struct {
		testSubscribe       testSubscribe
		givenRequesterEmail string
		givenAddresseeEmail string
		expErr              error
	}
	tcs := map[string]arg{
		"success": {
			testSubscribe: testSubscribe{
				mockInRequesterEmail: "van1@gmail.com",
				mockInAddresseeEmail: "van2@gmail.com",
				mockOutRequester: models.User{
					ID:    101,
					Email: "van1@gmail.com",
				},
				mockOutAddressee: models.User{
					ID:    102,
					Email: "van2@gmail.com",
				},
				mockFindRelationship:    models.Relationship{},
				mockFindRelationshipErr: errors.New("record not found"),
				mockInCreateRelationship: models.Relationship{
					ID:          1,
					AddresseeID: 102,
					RequesterID: 101,
					Type:        "Subscribed",
				},
				mockOutCreateRelationship: models.Relationship{
					ID:          1,
					AddresseeID: 102,
					RequesterID: 101,
					Type:        "Subscribed",
				},
			},
			givenRequesterEmail: "van1@gmail.com",
			givenAddresseeEmail: "van2@gmail.com",
		},
		"fail: can't subscribe to yourself": {
			givenAddresseeEmail: "van1@gmail.com",
			givenRequesterEmail: "van1@gmail.com",
			expErr:              errors.New("can't subscribe to yourself"),
		},
		"fail: invalid email address": {
			givenAddresseeEmail: "van2@gmail.com",
			givenRequesterEmail: "@@gmail.com",
			expErr:              errors.New("invalid email address"),
		},
		"fail: error Check": {
			testSubscribe: testSubscribe{
				mockInRequesterEmail: "van10@gmail.com",
				mockInAddresseeEmail: "van2@gmail.com",
				mockOutRequester:     models.User{},
				mockOutAddressee: models.User{
					ID:    102,
					Email: "van2@gmail.com",
				},
				mockCheckErr: errors.New("something wrong"),
			},
			givenRequesterEmail: "van10@gmail.com",
			givenAddresseeEmail: "van2@gmail.com",
			expErr:              errors.New("can't find van10@gmail.com email"),
		},
		"fail: already friend": {
			testSubscribe: testSubscribe{
				mockInRequesterEmail: "van1@gmail.com",
				mockInAddresseeEmail: "van2@gmail.com",
				mockOutRequester: models.User{
					ID:    101,
					Email: "van1@gmail.com",
				},
				mockOutAddressee: models.User{
					ID:    102,
					Email: "van2@gmail.com",
				},
				mockFindRelationship: models.Relationship{
					ID:          1,
					AddresseeID: 102,
					RequesterID: 101,
					Type:        "Friend",
				},
			},
			givenRequesterEmail: "van1@gmail.com",
			givenAddresseeEmail: "van2@gmail.com",
			expErr:              errors.New("van2@gmail.com is already your friend"),
		},
		"fail: requester blocked": {
			testSubscribe: testSubscribe{
				mockInRequesterEmail: "van1@gmail.com",
				mockInAddresseeEmail: "van2@gmail.com",
				mockOutRequester: models.User{
					ID:    101,
					Email: "van1@gmail.com",
				},
				mockOutAddressee: models.User{
					ID:    102,
					Email: "van2@gmail.com",
				},
				mockFindRelationship: models.Relationship{
					ID:          1,
					AddresseeID: 102,
					RequesterID: 101,
					Type:        "Blocked",
				},
			},
			givenRequesterEmail: "van1@gmail.com",
			givenAddresseeEmail: "van2@gmail.com",
			expErr:              errors.New("van1@gmail.com is blocked"),
		},
		"fail: error from repo": {
			testSubscribe: testSubscribe{
				mockInRequesterEmail: "van1@gmail.com",
				mockInAddresseeEmail: "van2@gmail.com",
				mockOutRequester: models.User{
					ID:    101,
					Email: "van1@gmail.com",
				},
				mockOutAddressee: models.User{
					ID:    102,
					Email: "van2@gmail.com",
				},
				mockFindRelationship:    models.Relationship{},
				mockFindRelationshipErr: errors.New("record not found"),
				mockInCreateRelationship: models.Relationship{
					ID:          1,
					AddresseeID: 102,
					RequesterID: 101,
					Type:        "Subscribed",
				},
				mockOutCreateRelationship: models.Relationship{},
				mockCreateRelationshipErr: errors.New("something wrong"),
			},
			givenRequesterEmail: "van1@gmail.com",
			givenAddresseeEmail: "van2@gmail.com",
			expErr:              errors.New("something wrong"),
		},
	}
	for s, tc := range tcs {
		t.Run(s, func(t *testing.T) {
			//GIVEN
			relaRepo := new(mocksR.Repository)
			userRepo := new(mocksU.Repository)
			userRepo.On("FindUserByEmail", mock.Anything, tc.testSubscribe.mockInRequesterEmail).
				Return(tc.testSubscribe.mockOutRequester, tc.testSubscribe.mockCheckErr)
			userRepo.On("FindUserByEmail", mock.Anything, tc.testSubscribe.mockInAddresseeEmail).
				Return(tc.testSubscribe.mockOutAddressee, tc.testSubscribe.mockCheckErr)
			relaRepo.On("FindRelationshipWithTwoEmail", mock.Anything, tc.testSubscribe.mockOutRequester.ID, tc.testSubscribe.mockOutAddressee.ID).
				Return(tc.testSubscribe.mockFindRelationship, tc.testSubscribe.mockFindRelationshipErr)
			relaRepo.On("CreateRelationship", mock.Anything, tc.testSubscribe.mockInCreateRelationship).
				Return(tc.testSubscribe.mockOutCreateRelationship, tc.testSubscribe.mockCreateRelationshipErr)

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
			err := svc.Subscribe(context.Background(), tc.givenRequesterEmail, tc.givenAddresseeEmail)

			//THEN
			if tc.expErr != nil {
				require.EqualError(t, err, tc.expErr.Error())
			} else {
				require.NoError(t, err)
			}
		})
	}
}
