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

func TestBlock(t *testing.T) {
	type block struct {
		mockInRequesterEmail                 string
		mockInAddresseeEmail                 string
		mockOutFindUserByEmail1              models.User
		mockOutFindUserByEmail2              models.User
		mockFindUserByEmailErr               error
		mockOutFindRelationshipWithTwoEmail1 models.Relationship
		mockOutFindRelationshipWithTwoEmail2 models.Relationship
		mockFindRelationshipWithTwoEmailErr  error
		mockInUpdateRelationship             models.Relationship
		mockOutUpdateRelationship            models.Relationship
		mockOutUpdateRelationshipErr         error
		mockInCreateRelationship1            models.Relationship
		mockOutCreateRelationship1           models.Relationship
		mockInCreateRelationship2            models.Relationship
		mockOutCreateRelationship2           models.Relationship
		mockCreateRelationshipErr            error
	}
	type arg struct {
		block               block
		givenRequesterEmail string
		givenAddresseeEmail string
		expErr              error
	}
	tcs := map[string]arg{
		"success": {
			block: block{
				mockInRequesterEmail: "van1@gmail.com",
				mockInAddresseeEmail: "van2@gmail.com",
				mockOutFindUserByEmail1: models.User{
					ID:    101,
					Email: "van1@gmail.com",
				},
				mockOutFindUserByEmail2: models.User{
					ID:    102,
					Email: "van2@gmail.com",
				},
				mockInCreateRelationship1: models.Relationship{
					ID:          1,
					AddresseeID: 102,
					RequesterID: 101,
					Type:        "Blocked",
				},
				mockOutCreateRelationship1: models.Relationship{
					ID:          1,
					AddresseeID: 102,
					RequesterID: 101,
					Type:        "Blocked",
				},
				mockInCreateRelationship2: models.Relationship{
					ID:          2,
					AddresseeID: 101,
					RequesterID: 102,
					Type:        "Blocked",
				},
				mockOutCreateRelationship2: models.Relationship{
					ID:          2,
					AddresseeID: 101,
					RequesterID: 102,
					Type:        "Blocked",
				},
			},
			givenRequesterEmail: "van1@gmail.com",
			givenAddresseeEmail: "van2@gmail.com",
		},
	}
	for s, tc := range tcs {
		t.Run(s, func(t *testing.T) {
			//GIVEN
			relaRepo := new(mocksR.Repository)
			userRepo := new(mocksU.Repository)
			userRepo.On("FindUserByEmail", mock.Anything, tc.block.mockInRequesterEmail).
				Return(tc.block.mockOutFindUserByEmail1, tc.block.mockFindUserByEmailErr)
			userRepo.On("FindUserByEmail", mock.Anything, tc.block.mockInAddresseeEmail).
				Return(tc.block.mockOutFindUserByEmail2, tc.block.mockFindUserByEmailErr)
			relaRepo.On("FindRelationshipWithTwoEmail", mock.Anything, tc.block.mockOutFindUserByEmail1.ID, tc.block.mockOutFindUserByEmail2.ID).
				Return(tc.block.mockOutFindRelationshipWithTwoEmail1, tc.block.mockFindRelationshipWithTwoEmailErr)
			relaRepo.On("FindRelationshipWithTwoEmail", mock.Anything, tc.block.mockOutFindUserByEmail2.ID, tc.block.mockOutFindUserByEmail1.ID).
				Return(tc.block.mockOutFindRelationshipWithTwoEmail2, tc.block.mockFindRelationshipWithTwoEmailErr)
			relaRepo.On("UpdateRelationship", mock.Anything, tc.block.mockInUpdateRelationship).
				Return(tc.block.mockOutUpdateRelationship, tc.block.mockOutUpdateRelationshipErr)
			relaRepo.On("CreateRelationship", mock.Anything, tc.block.mockInCreateRelationship1).
				Return(tc.block.mockOutCreateRelationship1, tc.block.mockCreateRelationshipErr)
			relaRepo.On("CreateRelationship", mock.Anything, tc.block.mockInCreateRelationship2).
				Return(tc.block.mockOutCreateRelationship2, tc.block.mockCreateRelationshipErr)
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
			err := svc.Block(context.Background(), tc.givenRequesterEmail, tc.givenAddresseeEmail)

			//THEN
			if tc.expErr != nil {
				require.EqualError(t, err, tc.expErr.Error())
			} else {
				require.NoError(t, err)
			}
		})
	}
}
