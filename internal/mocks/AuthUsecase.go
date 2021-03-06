// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	internal "github.com/ardafirdausr/todo-server/internal"
	entity "github.com/ardafirdausr/todo-server/internal/entity"

	mock "github.com/stretchr/testify/mock"
)

// AuthUsecase is an autogenerated mock type for the AuthUsecase type
type AuthUsecase struct {
	mock.Mock
}

// GenerateAuthToken provides a mock function with given fields: user, tokenizer
func (_m *AuthUsecase) GenerateAuthToken(user entity.User, tokenizer internal.Tokenizer) (string, error) {
	ret := _m.Called(user, tokenizer)

	var r0 string
	if rf, ok := ret.Get(0).(func(entity.User, internal.Tokenizer) string); ok {
		r0 = rf(user, tokenizer)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(entity.User, internal.Tokenizer) error); ok {
		r1 = rf(user, tokenizer)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SSO provides a mock function with given fields: token, authenticator
func (_m *AuthUsecase) SSO(token string, authenticator internal.SSOAuthenticator) (*entity.User, error) {
	ret := _m.Called(token, authenticator)

	var r0 *entity.User
	if rf, ok := ret.Get(0).(func(string, internal.SSOAuthenticator) *entity.User); ok {
		r0 = rf(token, authenticator)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, internal.SSOAuthenticator) error); ok {
		r1 = rf(token, authenticator)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
