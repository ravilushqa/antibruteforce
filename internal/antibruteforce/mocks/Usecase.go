// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import context "context"
import mock "github.com/stretchr/testify/mock"

// Usecase is an autogenerated mock type for the Usecase type
type Usecase struct {
	mock.Mock
}

// BlacklistAdd provides a mock function with given fields: ctx, subnet
func (_m *Usecase) BlacklistAdd(ctx context.Context, subnet string) error {
	ret := _m.Called(ctx, subnet)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, subnet)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// BlacklistRemove provides a mock function with given fields: ctx, subnet
func (_m *Usecase) BlacklistRemove(ctx context.Context, subnet string) error {
	ret := _m.Called(ctx, subnet)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, subnet)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Check provides a mock function with given fields: ctx, login, password, ip
func (_m *Usecase) Check(ctx context.Context, login string, password string, ip string) error {
	ret := _m.Called(ctx, login, password, ip)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string, string) error); ok {
		r0 = rf(ctx, login, password, ip)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Reset provides a mock function with given fields: ctx, login, ip
func (_m *Usecase) Reset(ctx context.Context, login string, ip string) error {
	ret := _m.Called(ctx, login, ip)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) error); ok {
		r0 = rf(ctx, login, ip)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// WhitelistAdd provides a mock function with given fields: ctx, subnet
func (_m *Usecase) WhitelistAdd(ctx context.Context, subnet string) error {
	ret := _m.Called(ctx, subnet)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, subnet)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// WhitelistRemove provides a mock function with given fields: ctx, subnet
func (_m *Usecase) WhitelistRemove(ctx context.Context, subnet string) error {
	ret := _m.Called(ctx, subnet)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, subnet)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}