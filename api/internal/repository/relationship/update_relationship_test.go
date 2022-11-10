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

func TestUpdateRelationship(t *testing.T) {
	type arg struct {
		givenInput models.Relationship
		expRs      models.Relationship
		expErr     error
	}
	tcs := map[string]arg{
		"success": {
			givenInput: models.Relationship{
				ID:          1,
				AddresseeID: 102,
				RequesterID: 101,
				Type:        "Blocked",
			},
			expRs: models.Relationship{
				ID:          1,
				AddresseeID: 102,
				RequesterID: 101,
				Type:        "Blocked",
			},
		},
		"fail: violates check constraint": {
			givenInput: models.Relationship{
				ID:          1,
				AddresseeID: 102,
				RequesterID: 102,
				Type:        "Blocked",
			},
			expRs:  models.Relationship{},
			expErr: errors.New("ERROR: new row for relation \"relationships\" violates check constraint \"relationships_check\" (SQLSTATE 23514)"),
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
			rs, err := instance.UpdateRelationship(context.Background(), tc.givenInput)

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
