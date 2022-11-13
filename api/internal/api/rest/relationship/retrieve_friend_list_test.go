package relationship

import (
	"context"
	"errors"
	mocks "fm/api/internal/mocks/controller/relationship"
	"fm/api/internal/models"
	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestRetrieveFriendList(t *testing.T) {
	type retrieveFriendList struct {
		mockIn  string
		mockOut []models.User
		mockErr error
	}
	type arg struct {
		retrieveFriendList           retrieveFriendList
		retrieveFriendListMockCalled bool
		givenBody                    string
		expRs                        string
		expHTTPCode                  int
	}
	tcs := map[string]arg{
		"success": {
			retrieveFriendList: retrieveFriendList{
				mockIn: "van1@gmail.com",
				mockOut: []models.User{
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
			retrieveFriendListMockCalled: true,
			givenBody: `{
							"email":"van1@gmail.com"
						}`,
			expRs: `{
						"success": "true",
						"friends": [
							{
								"email": "van2@gmail.com"
							},
							{
								"email": "van3@gmail.com"
							}
						],
						"count": 2
					}`,
			expHTTPCode: http.StatusOK,
		},
		"success: empty": {
			retrieveFriendList: retrieveFriendList{
				mockIn:  "van1@gmail.com",
				mockOut: []models.User{},
			},
			retrieveFriendListMockCalled: true,
			givenBody: `{
							"email":"van1@gmail.com"
						}`,
			expRs: `{
						"success": "true",
						"friends": null,
						"count": 0
					}`,
			expHTTPCode: http.StatusOK,
		},
		"fail:bad request": {
			retrieveFriendListMockCalled: false,
			givenBody: `{
							"email":""
						}`,
			expRs:       `{"code":"invalid_request", "description":"invalid email request"}`,
			expHTTPCode: http.StatusBadRequest,
		},
		"fail: error from service": {
			retrieveFriendList: retrieveFriendList{
				mockIn:  "van1@gmail.com",
				mockOut: nil,
				mockErr: errors.New("internal_error"),
			},
			retrieveFriendListMockCalled: true,
			givenBody: `{
							"email":"van1@gmail.com"
						}`,
			expRs:       `{"code":"invalid_request", "description":"internal_error"}`,
			expHTTPCode: http.StatusInternalServerError,
		},
	}
	for s, tc := range tcs {
		t.Run(s, func(t *testing.T) {
			//MOCK
			mockSvc := new(mocks.Service)
			if tc.retrieveFriendListMockCalled {
				mockSvc.ExpectedCalls = []*mock.Call{
					mockSvc.On("FriendList", mock.Anything, tc.retrieveFriendList.mockIn).
						Return(tc.retrieveFriendList.mockOut, tc.retrieveFriendList.mockErr),
				}
			}
			//GIVEN
			req := httptest.NewRequest(http.MethodPost, "/friend-list", strings.NewReader(tc.givenBody))
			routeCtx := chi.NewRouteContext()
			ctx := context.WithValue(req.Context(), chi.RouteCtxKey, routeCtx)
			req = req.WithContext(ctx)
			res := httptest.NewRecorder()

			//WHEN
			instance := New(mockSvc)
			instance.RetrieveFriendList(res, req)

			//THEN
			require.Equal(t, tc.expHTTPCode, res.Code)
			require.JSONEq(t, tc.expRs, res.Body.String())
			mockSvc.AssertExpectations(t)
		})
	}
}
