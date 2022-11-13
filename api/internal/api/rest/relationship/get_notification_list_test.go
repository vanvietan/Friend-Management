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

func TestGetNotificationList(t *testing.T) {
	type getNotificationList struct {
		mockInSender  string
		mockInText    string
		mockOutEmails []models.User
		mockErr       error
	}
	type arg struct {
		getNotificationList           getNotificationList
		getNotificationListMockCalled bool
		givenBody                     string
		expRs                         string
		expHTTPCode                   int
	}
	tcs := map[string]arg{
		"success": {
			getNotificationList: getNotificationList{
				mockInSender: "van1@gmail.com",
				mockInText:   "Hello World! van5@gmail.com",
				mockOutEmails: []models.User{
					{
						ID:    2,
						Email: "van2@gmail.com",
					},
					{
						ID:    5,
						Email: "van5@gmail.com",
					},
				},
			},
			getNotificationListMockCalled: true,
			givenBody: `{
							"sender":"van1@gmail.com",
							"text":"Hello World! van5@gmail.com"
						}`,
			expRs:       `{"success":"true","recipients":[{"email":"van2@gmail.com"},{"email":"van5@gmail.com"}]}`,
			expHTTPCode: http.StatusOK,
		},
		"fail: invalid request": {
			getNotificationListMockCalled: false,
			givenBody: `{
							"sender":"",
							"text":"Hello World! van5@gmail.com"
						}`,
			expRs:       `{"code":"invalid_request", "description":"invalid request, both sender and text must not be empty"}`,
			expHTTPCode: http.StatusBadRequest,
		},
		"fail: error from service": {
			getNotificationList: getNotificationList{
				mockInSender: "van1@gmail.com",
				mockInText:   "Hello World! van5@gmail.com",
				mockOutEmails: []models.User{
					{
						ID:    2,
						Email: "van2@gmail.com",
					},
					{
						ID:    5,
						Email: "van5@gmail.com",
					},
				},
				mockErr: errors.New("something wrong"),
			},
			getNotificationListMockCalled: true,
			givenBody: `{
							"sender":"van1@gmail.com",
							"text":"Hello World! van5@gmail.com"
						}`,
			expRs:       `{"code":"invalid_request", "description":"something wrong"}`,
			expHTTPCode: http.StatusInternalServerError,
		},
	}
	for s, tc := range tcs {
		t.Run(s, func(t *testing.T) {
			//MOCK
			mockSvc := new(mocks.Service)
			if tc.getNotificationListMockCalled {
				mockSvc.ExpectedCalls = []*mock.Call{
					mockSvc.On("NotificationList", mock.Anything, tc.getNotificationList.mockInSender, tc.getNotificationList.mockInText).
						Return(tc.getNotificationList.mockOutEmails, tc.getNotificationList.mockErr),
				}
			}
			//GIVEN
			req := httptest.NewRequest(http.MethodPost, "/notification-list", strings.NewReader(tc.givenBody))
			routeCtx := chi.NewRouteContext()
			ctx := context.WithValue(req.Context(), chi.RouteCtxKey, routeCtx)
			req = req.WithContext(ctx)
			res := httptest.NewRecorder()

			//WHEN
			instance := New(mockSvc)
			instance.GetNotificationList(res, req)

			//THEN
			require.Equal(t, tc.expHTTPCode, res.Code)
			require.JSONEq(t, tc.expRs, res.Body.String())
			mockSvc.AssertExpectations(t)
		})
	}
}
