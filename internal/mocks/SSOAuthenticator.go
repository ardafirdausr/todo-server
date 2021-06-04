// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	entity "github.com/ardafirdausr/todo-server/internal/entity"

	mock "github.com/stretchr/testify/mock"
)

// SSOAuthenticator is an autogenerated mock type for the SSOAuthenticator type
type SSOAuthenticator struct {
	mock.Mock
}

// Authenticate provides a mock function with given fields: token
func (_m *SSOAuthenticator) Authenticate(token string) (*entity.User, error) {
	ret := _m.Called(token)

	var r0 *entity.User
	if rf, ok := ret.Get(0).(func(string) *entity.User); ok {
		r0 = rf(token)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(token)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
