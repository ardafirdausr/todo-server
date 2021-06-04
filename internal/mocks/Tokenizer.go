// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	entity "github.com/ardafirdausr/todo-server/internal/entity"

	mock "github.com/stretchr/testify/mock"
)

// Tokenizer is an autogenerated mock type for the Tokenizer type
type Tokenizer struct {
	mock.Mock
}

// Generate provides a mock function with given fields: _a0
func (_m *Tokenizer) Generate(_a0 entity.TokenPayload) (string, error) {
	ret := _m.Called(_a0)

	var r0 string
	if rf, ok := ret.Get(0).(func(entity.TokenPayload) string); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(entity.TokenPayload) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
