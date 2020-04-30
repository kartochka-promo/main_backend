// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	models "2020_1_drop_table/internal/microservices/staff/models"
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// StaffClientInterface is an autogenerated mock type for the StaffClientInterface type
type StaffClientInterface struct {
	mock.Mock
}

// AddSessionInMetadata provides a mock function with given fields: ctx
func (_m *StaffClientInterface) AddSessionInMetadata(ctx context.Context) context.Context {
	ret := _m.Called(ctx)

	var r0 context.Context
	if rf, ok := ret.Get(0).(func(context.Context) context.Context); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(context.Context)
		}
	}

	return r0
}

// GetById provides a mock function with given fields: ctx, id
func (_m *StaffClientInterface) GetById(ctx context.Context, id int) (models.SafeStaff, error) {
	ret := _m.Called(ctx, id)

	var r0 models.SafeStaff
	if rf, ok := ret.Get(0).(func(context.Context, int) models.SafeStaff); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(models.SafeStaff)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetFromSession provides a mock function with given fields: ctx
func (_m *StaffClientInterface) GetFromSession(ctx context.Context) (models.SafeStaff, error) {
	ret := _m.Called(ctx)

	var r0 models.SafeStaff
	if rf, ok := ret.Get(0).(func(context.Context) models.SafeStaff); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Get(0).(models.SafeStaff)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
