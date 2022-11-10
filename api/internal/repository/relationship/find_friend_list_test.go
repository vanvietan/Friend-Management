package relationship

import (
	"context"
	"fm/api/internal/config"
	"fm/api/internal/models"
	"fm/api/internal/pkg"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestFindFriendList(t *testing.T) {
	type arg struct {
		givenID int64
		expRs   []models.User
		expErr  error
	}
	tcs := map[string]arg{
		"success": {
			givenID: 101,
			expRs: []models.User{
				{
					ID:    102,
					Email: "van2@gmail.com",
				},
			},
		},
		"success: empty": {
			givenID: 105,
			expRs:   []models.User{},
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
			rs, err := instance.FindFriendList(context.Background(), tc.givenID)

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
