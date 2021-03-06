// Code generated by mockery 2.9.0. DO NOT EDIT.

package mocks

import (
	context "context"

	domain "github.com/hezbymuhammad/golang-marvel-demo/domain"
	mock "github.com/stretchr/testify/mock"
)

// CharacterReadRepository is an autogenerated mock type for the CharacterReadRepository type
type CharacterReadRepository struct {
	mock.Mock
}

// Fetch provides a mock function with given fields: ctx, page
func (_m *CharacterReadRepository) Fetch(ctx context.Context, page int) ([]int, error) {
	ret := _m.Called(ctx, page)

	var r0 []int
	if rf, ok := ret.Get(0).(func(context.Context, int) []int); ok {
		r0 = rf(ctx, page)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]int)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int) error); ok {
		r1 = rf(ctx, page)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByID provides a mock function with given fields: ctx, id
func (_m *CharacterReadRepository) GetByID(ctx context.Context, id int) (domain.Character, error) {
	ret := _m.Called(ctx, id)

	var r0 domain.Character
	if rf, ok := ret.Get(0).(func(context.Context, int) domain.Character); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(domain.Character)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
