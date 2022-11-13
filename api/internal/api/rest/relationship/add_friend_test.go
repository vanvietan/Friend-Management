package relationship

import (
	"context"
	"errors"
	mocks "fm/api/internal/mocks/controller/relationship"
	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestAddFriend(t *testing.T) {
	type addFriend struct {
		mockRequesterEmail string
		mockAddressEmail   string
		mockErr            error
	}
	type arg struct {
		addFriend           addFriend
		addFriendMockCalled bool
		givenBody           string
		expRs               string
		expHTTPCode         int
	}
	tcs := map[string]arg{
		"success": {
			addFriend: addFriend{
				mockRequesterEmail: "van1@gmail.com",
				mockAddressEmail:   "van2@gmail.com",
			},
			addFriendMockCalled: true,
			givenBody: `{ 
							"friends":
								[
									"van1@gmail.com", 
									"van2@gmail.com"
								] 
						}`,
			expRs:       `{"success":"true"}`,
			expHTTPCode: http.StatusOK,
		},
		"fail: invalid email address": {
			addFriendMockCalled: false,
			givenBody: `{ 
							"friends":
								[
									"van2@gmail.com"
								] 
						}`,
			expRs:       `{"code":"invalid_request", "description":"need 2 emails"}`,
			expHTTPCode: http.StatusBadRequest,
		},
		"fail: error from service": {
			addFriend: addFriend{
				mockRequesterEmail: "van1@gmail.com",
				mockAddressEmail:   "van2@gmail.com",
				mockErr:            errors.New("something wrong"),
			},
			addFriendMockCalled: true,
			givenBody: `{ 
							"friends":
								[
									"van1@gmail.com", 
									"van2@gmail.com"
								] 
						}`,
			expRs:       `{"code":"internal_server_error", "description":"something wrong"}`,
			expHTTPCode: http.StatusInternalServerError,
		},
	}
	for s, tc := range tcs {
		t.Run(s, func(t *testing.T) {
			//MOCK
			mockSvc := new(mocks.Service)
			if tc.addFriendMockCalled {
				mockSvc.ExpectedCalls = []*mock.Call{
					mockSvc.On("AddFriend", mock.Anything, tc.addFriend.mockRequesterEmail, tc.addFriend.mockAddressEmail).
						Return(tc.addFriend.mockErr),
				}
			}

			//GIVEN
			req := httptest.NewRequest(http.MethodPost, "/add-friend", strings.NewReader(tc.givenBody))
			routeCtx := chi.NewRouteContext()
			ctx := context.WithValue(req.Context(), chi.RouteCtxKey, routeCtx)
			req = req.WithContext(ctx)
			res := httptest.NewRecorder()

			//WHEN
			instance := New(mockSvc)
			instance.AddFriend(res, req)

			//THEN
			require.Equal(t, tc.expHTTPCode, res.Code)
			require.JSONEq(t, tc.expRs, res.Body.String())
			mockSvc.AssertExpectations(t)
		})
	}
}
