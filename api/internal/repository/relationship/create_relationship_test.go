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

func TestCreateRelationship(t *testing.T) {
	type arg struct {
		givenInput models.Relationship
		expRs      models.Relationship
		expErr     error
	}
	tcs := map[string]arg{
		"success": {
			givenInput: models.Relationship{
				ID:          11,
				AddresseeID: 101,
				RequesterID: 105,
				Type:        "Friend",
			},
			expRs: models.Relationship{
				ID:          11,
				AddresseeID: 101,
				RequesterID: 105,
				Type:        "Friend",
			},
		},
		"fail:  violates check constraint ": {
			givenInput: models.Relationship{
				ID:          11,
				AddresseeID: 101,
				RequesterID: 101,
				Type:        "Friend",
			},
			expRs:  models.Relationship{},
			expErr: errors.New("ERROR: new row for relation \"relationships\" violates check constraint \"relationships_check\" (SQLSTATE 23514)"),
		},
		"fail: violates unique constraint": {
			givenInput: models.Relationship{
				ID:          1,
				AddresseeID: 105,
				RequesterID: 101,
				Type:        "Friend",
			},
			expRs:  models.Relationship{},
			expErr: errors.New("ERROR: duplicate key value violates unique constraint \"relationships_pkey\" (SQLSTATE 23505)"),
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
			rs, err := instance.CreateRelationship(context.Background(), tc.givenInput)

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
