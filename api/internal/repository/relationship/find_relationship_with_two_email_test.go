package relationship

import (
	"context"
	"errors"
	"fm/api/internal/config"
	"fm/api/internal/models"
	"fm/api/internal/pkg"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestFindRelationshipWithTwoEmail(t *testing.T) {
	type arg struct {
		givenRequesterID int64
		givenAddresseeID int64
		expRs            models.Relationship
		expErr           error
	}
	tcs := map[string]arg{
		"success": {
			givenRequesterID: 101,
			givenAddresseeID: 102,
			expRs: models.Relationship{
				ID:          1,
				AddresseeID: 102,
				RequesterID: 101,
				Type:        "Friend",
			},
		},
		"fail: record not found": {
			givenRequesterID: 101,
			givenAddresseeID: 105,
			expRs:            models.Relationship{},
			expErr:           errors.New("record not found"),
		},
	}
	dbConn, errDB := config.GetDatabaseConnection()
	require.NoError(t, errDB)

	errExe := pkg.ExecuteTestData(dbConn, "./testdata/relationships.sql")
	require.NoError(t, errExe)

	for s, tc := range tcs {
		t.Run(s, func(t *testing.T) {
			//GiVEN
			instance := New(dbConn)

			//WHEN
			rs, err := instance.FindRelationshipWithTwoEmail(context.Background(), tc.givenRequesterID, tc.givenAddresseeID)

			//THEN
			if tc.expErr != nil {
				require.EqualError(t, err, tc.expErr.Error())
			} else {
				require.NoError(t, err)
				require.Equal(t, tc.expRs, rs)
			}
		})
	}
}
