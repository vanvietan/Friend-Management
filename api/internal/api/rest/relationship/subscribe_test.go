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

func TestSubscribe(t *testing.T) {
	type subscribe struct {
		mockRequester string
		mockAddressee string
		mockErr       error
	}
	type arg struct {
		subscribe           subscribe
		subscribeMockCalled bool
		givenBody           string
		expRs               string
		expHTTPCode         int
	}
	tcs := map[string]arg{
		"success": {
			subscribe: subscribe{
				mockRequester: "van1@gmail.com",
				mockAddressee: "van2@gmail.com",
			},
			subscribeMockCalled: true,
			givenBody: `{
							"requester":"van1@gmail.com",
							"target":"van2@gmail.com"
						}`,
			expRs:       `{"success":"true"}`,
			expHTTPCode: http.StatusOK,
		},
		"fail: bad request": {
			subscribeMockCalled: false,
			givenBody: `{
							"target":"van2@gmail.com"
						}`,
			expRs:       `{"code":"invalid_request", "description":"invalid email"}`,
			expHTTPCode: http.StatusBadRequest,
		},
		"fail: error from service": {
			subscribe: subscribe{
				mockRequester: "van1@gmail.com",
				mockAddressee: "van2@gmail.com",
				mockErr:       errors.New("something wrong"),
			},
			subscribeMockCalled: true,
			givenBody: `{
							"requester":"van1@gmail.com",
							"target":"van2@gmail.com"
						}`,
			expRs:       `{"code":"internal_server_error", "description":"something wrong"}`,
			expHTTPCode: http.StatusInternalServerError,
		},
	}
	for s, tc := range tcs {
		t.Run(s, func(t *testing.T) {
			//MOCK
			mockSvc := new(mocks.Service)
			if tc.subscribeMockCalled {
				mockSvc.ExpectedCalls = []*mock.Call{
					mockSvc.On("Subscribe", mock.Anything, tc.subscribe.mockRequester, tc.subscribe.mockAddressee).
						Return(tc.subscribe.mockErr),
				}
			}
			//GIVEN
			req := httptest.NewRequest(http.MethodPost, "/subscribe", strings.NewReader(tc.givenBody))
			routeCtx := chi.NewRouteContext()
			ctx := context.WithValue(req.Context(), chi.RouteCtxKey, routeCtx)
			req = req.WithContext(ctx)
			res := httptest.NewRecorder()

			//WHEN
			instance := New(mockSvc)
			instance.Subscribe(res, req)

			//THEN
			require.Equal(t, tc.expHTTPCode, res.Code)
			require.JSONEq(t, tc.expRs, res.Body.String())
			mockSvc.AssertExpectations(t)
		})
	}
}
