// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	context "context"
	models "fm/api/internal/models"

	mock "github.com/stretchr/testify/mock"
)

// Service is an autogenerated mock type for the Service type
type Service struct {
	mock.Mock
}

// AddFriend provides a mock function with given fields: ctx, requesterEmail, addresseeEmail
func (_m *Service) AddFriend(ctx context.Context, requesterEmail string, addresseeEmail string) error {
	ret := _m.Called(ctx, requesterEmail, addresseeEmail)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) error); ok {
		r0 = rf(ctx, requesterEmail, addresseeEmail)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Block provides a mock function with given fields: ctx, requester, addressee
func (_m *Service) Block(ctx context.Context, requester string, addressee string) error {
	ret := _m.Called(ctx, requester, addressee)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) error); ok {
		r0 = rf(ctx, requester, addressee)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CommonFriend provides a mock function with given fields: ctx, requesterEmail, addresseeEmail
func (_m *Service) CommonFriend(ctx context.Context, requesterEmail string, addresseeEmail string) ([]models.User, error) {
	ret := _m.Called(ctx, requesterEmail, addresseeEmail)

	var r0 []models.User
	if rf, ok := ret.Get(0).(func(context.Context, string, string) []models.User); ok {
		r0 = rf(ctx, requesterEmail, addresseeEmail)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, requesterEmail, addresseeEmail)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FriendList provides a mock function with given fields: ctx, input
func (_m *Service) FriendList(ctx context.Context, input string) ([]models.User, error) {
	ret := _m.Called(ctx, input)

	var r0 []models.User
	if rf, ok := ret.Get(0).(func(context.Context, string) []models.User); ok {
		r0 = rf(ctx, input)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, input)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NotificationList provides a mock function with given fields: ctx, sender, text
func (_m *Service) NotificationList(ctx context.Context, sender string, text string) ([]models.User, error) {
	ret := _m.Called(ctx, sender, text)

	var r0 []models.User
	if rf, ok := ret.Get(0).(func(context.Context, string, string) []models.User); ok {
		r0 = rf(ctx, sender, text)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, sender, text)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Subscribe provides a mock function with given fields: ctx, requester, addressee
func (_m *Service) Subscribe(ctx context.Context, requester string, addressee string) error {
	ret := _m.Called(ctx, requester, addressee)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) error); ok {
		r0 = rf(ctx, requester, addressee)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewService interface {
	mock.TestingT
	Cleanup(func())
}

// NewService creates a new instance of Service. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewService(t mockConstructorTestingTNewService) *Service {
	mock := &Service{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
