// Code generated by mockery v1.0.0. DO NOT EDIT.

package auth

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// MockRoleSource is an autogenerated mock type for the RoleSource type
type MockRoleSource struct {
	mock.Mock
}

// FindUserRole provides a mock function with given fields: ctx, organizationID, userID
func (_m *MockRoleSource) FindUserRole(ctx context.Context, organizationID uint, userID uint) (string, bool, error) {
	ret := _m.Called(ctx, organizationID, userID)

	var r0 string
	if rf, ok := ret.Get(0).(func(context.Context, uint, uint) string); ok {
		r0 = rf(ctx, organizationID, userID)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 bool
	if rf, ok := ret.Get(1).(func(context.Context, uint, uint) bool); ok {
		r1 = rf(ctx, organizationID, userID)
	} else {
		r1 = ret.Get(1).(bool)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context, uint, uint) error); ok {
		r2 = rf(ctx, organizationID, userID)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}
