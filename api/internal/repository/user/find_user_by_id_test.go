package user

import (
	"context"
	"errors"
	"fm/api/internal/config"
	"fm/api/internal/models"
	"fm/api/internal/pkg"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestFindUserByID(t *testing.T) {
	type arg struct {
		givenID   int64
		expResult models.User
		expErr    error
	}

	tcs := map[string]arg{
		"success: ": {
			givenID: 101,
			expResult: models.User{
				ID:    101,
				Email: "van1@gmail.com",
			},
		},
		"fail: record not found": {
			givenID:   200,
			expResult: models.User{},
			expErr:    errors.New("record not found"),
		},
	}
	dbConn, errDB := config.GetDatabaseConnection()
	require.NoError(t, errDB)

	errExe := pkg.ExecuteTestData(dbConn, "./testdata/users.sql")
	require.NoError(t, errExe)

	for s, tc := range tcs {
		t.Run(s, func(t *testing.T) {
			//GiVEN
			instance := New(dbConn)

			//WHEN
			rs, err := instance.FindUserByID(context.Background(), tc.givenID)

			//THEN
			if tc.expErr != nil {
				require.EqualError(t, err, tc.expErr.Error())
			} else {
				require.NoError(t, err)
				require.Equal(t, tc.expResult, rs)
			}
		})
	}

}
