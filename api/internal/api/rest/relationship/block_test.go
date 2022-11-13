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

func TestBlock(t *testing.T) {
	type block struct {
		mockInRequester string
		mockInAddressee string
		mockErr         error
	}
	type arg struct {
		block           block
		blockMockCalled bool
		givenBody       string
		expRs           string
		expHTTPCode     int
	}
	tcs := map[string]arg{
		"success": {
			block: block{
				mockInRequester: "van1@gmail.com",
				mockInAddressee: "van2@gmail.com",
			},
			blockMockCalled: true,
			givenBody: `{
							"requester":"van1@gmail.com",
							"target":"van2@gmail.com"
						}`,
			expRs:       `{"success":"true"}`,
			expHTTPCode: http.StatusOK,
		},
		"fail:invalid email": {
			blockMockCalled: false,
			givenBody: `{
							"requester":"",
							"target":"van2@gmail.com"
						}`,
			expRs:       `{"code":"invalid_request", "description":"invalid email"}`,
			expHTTPCode: http.StatusBadRequest,
		},
		"fail: error from repo": {
			block: block{
				mockInRequester: "van1@gmail.com",
				mockInAddressee: "van2@gmail.com",
				mockErr:         errors.New("something wrong"),
			},
			blockMockCalled: true,
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
			if tc.blockMockCalled {
				mockSvc.ExpectedCalls = []*mock.Call{
					mockSvc.On("Block", mock.Anything, tc.block.mockInRequester, tc.block.mockInAddressee).
						Return(tc.block.mockErr),
				}
			}

			//GIVEN
			req := httptest.NewRequest(http.MethodPost, "/block", strings.NewReader(tc.givenBody))
			routeCtx := chi.NewRouteContext()
			ctx := context.WithValue(req.Context(), chi.RouteCtxKey, routeCtx)
			req = req.WithContext(ctx)
			res := httptest.NewRecorder()

			//WHEN
			instance := New(mockSvc)
			instance.Block(res, req)

			//THEN
			require.Equal(t, tc.expHTTPCode, res.Code)
			require.JSONEq(t, tc.expRs, res.Body.String())
			mockSvc.AssertExpectations(t)
		})
	}
}
