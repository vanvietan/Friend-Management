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

func TestRetrieveCommonFriend(t *testing.T) {
	type retrieveCommonFriend struct {
		mockInEmail1 string
		mockInEmail2 string
		mockOut      []models.User
		mockErr      error
	}
	type arg struct {
		retrieveCommonFriend           retrieveCommonFriend
		retrieveCommonFriendMockCalled bool
		givenBody                      string
		expRs                          string
		expHTTPCode                    int
	}
	tcs := map[string]arg{
		"success": {
			retrieveCommonFriend: retrieveCommonFriend{
				mockInEmail1: "van1@gmail.com",
				mockInEmail2: "van2@gmail.com",
				mockOut: []models.User{
					{
						ID:    3,
						Email: "van3@gmail.com",
					},
					{
						ID:    4,
						Email: "van4@gmail.com",
					},
				},
			},
			retrieveCommonFriendMockCalled: true,
			givenBody: `{ 
							"friends":
								[
									"van1@gmail.com", 
									"van2@gmail.com"
								] 
						}`,
			expRs: `{
						"success": "true",
						"friends": [
							{
								"email": "van3@gmail.com"
							},
							{
								"email": "van4@gmail.com"
							}
						],
						"count": 2
					}`,
			expHTTPCode: http.StatusOK,
		},
		"fail:invalid email": {
			retrieveCommonFriendMockCalled: false,
			givenBody: `{ 
							"friends":
								[
									"van2@gmail.com"
								] 
						}`,
			expRs:       `{"code":"invalid_request", "description":"need 2 emails"}`,
			expHTTPCode: http.StatusBadRequest,
		},
		"fail: error from repo": {
			retrieveCommonFriend: retrieveCommonFriend{
				mockInEmail1: "van1@gmail.com",
				mockInEmail2: "van2@gmail.com",
				mockOut:      nil,
				mockErr:      errors.New("something wrong"),
			},
			retrieveCommonFriendMockCalled: true,
			givenBody: `{ 
							"friends":
								[
									"van1@gmail.com", 
									"van2@gmail.com"
								] 
						}`,
			expRs:       `{"code":"invalid_request", "description":"something wrong"}`,
			expHTTPCode: http.StatusInternalServerError,
		},
		"success:empty": {
			retrieveCommonFriend: retrieveCommonFriend{
				mockInEmail1: "van1@gmail.com",
				mockInEmail2: "van2@gmail.com",
				mockOut:      []models.User{},
			},
			retrieveCommonFriendMockCalled: true,
			givenBody: `{ 
							"friends":
								[
									"van1@gmail.com", 
									"van2@gmail.com"
								] 
						}`,
			expRs:       `{"success":"true","friends":null,"count":0}`,
			expHTTPCode: http.StatusOK,
		},
	}
	for s, tc := range tcs {
		t.Run(s, func(t *testing.T) {
			//MOCK
			mockSvc := new(mocks.Service)
			if tc.retrieveCommonFriendMockCalled {
				mockSvc.ExpectedCalls = []*mock.Call{
					mockSvc.On("CommonFriend", mock.Anything, tc.retrieveCommonFriend.mockInEmail1, tc.retrieveCommonFriend.mockInEmail2).
						Return(tc.retrieveCommonFriend.mockOut, tc.retrieveCommonFriend.mockErr),
				}
			}
			//GIVEN
			req := httptest.NewRequest(http.MethodPost, "/common-friend", strings.NewReader(tc.givenBody))
			routeCtx := chi.NewRouteContext()
			ctx := context.WithValue(req.Context(), chi.RouteCtxKey, routeCtx)
			req = req.WithContext(ctx)
			res := httptest.NewRecorder()

			//WHEN
			instance := New(mockSvc)
			instance.RetrieveCommonFriend(res, req)

			//THEN
			require.Equal(t, tc.expHTTPCode, res.Code)
			require.JSONEq(t, tc.expRs, res.Body.String())
			mockSvc.AssertExpectations(t)
		})
	}
}
