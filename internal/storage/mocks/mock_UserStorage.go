// Code generated by mockery v2.35.4. DO NOT EDIT.

package mocks

import (
	context "context"

	entity "github.com/rrgmc/debefix-sample-app/internal/entity"
	mock "github.com/stretchr/testify/mock"

	uuid "github.com/google/uuid"
)

// UserStorage is an autogenerated mock type for the UserStorage type
type UserStorage struct {
	mock.Mock
}

type UserStorage_Expecter struct {
	mock *mock.Mock
}

func (_m *UserStorage) EXPECT() *UserStorage_Expecter {
	return &UserStorage_Expecter{mock: &_m.Mock}
}

// AddUser provides a mock function with given fields: ctx, user
func (_m *UserStorage) AddUser(ctx context.Context, user entity.User) (entity.User, error) {
	ret := _m.Called(ctx, user)

	var r0 entity.User
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, entity.User) (entity.User, error)); ok {
		return rf(ctx, user)
	}
	if rf, ok := ret.Get(0).(func(context.Context, entity.User) entity.User); ok {
		r0 = rf(ctx, user)
	} else {
		r0 = ret.Get(0).(entity.User)
	}

	if rf, ok := ret.Get(1).(func(context.Context, entity.User) error); ok {
		r1 = rf(ctx, user)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UserStorage_AddUser_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'AddUser'
type UserStorage_AddUser_Call struct {
	*mock.Call
}

// AddUser is a helper method to define mock.On call
//   - ctx context.Context
//   - user entity.User
func (_e *UserStorage_Expecter) AddUser(ctx interface{}, user interface{}) *UserStorage_AddUser_Call {
	return &UserStorage_AddUser_Call{Call: _e.mock.On("AddUser", ctx, user)}
}

func (_c *UserStorage_AddUser_Call) Run(run func(ctx context.Context, user entity.User)) *UserStorage_AddUser_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(entity.User))
	})
	return _c
}

func (_c *UserStorage_AddUser_Call) Return(_a0 entity.User, _a1 error) *UserStorage_AddUser_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *UserStorage_AddUser_Call) RunAndReturn(run func(context.Context, entity.User) (entity.User, error)) *UserStorage_AddUser_Call {
	_c.Call.Return(run)
	return _c
}

// DeleteUserByID provides a mock function with given fields: ctx, userID
func (_m *UserStorage) DeleteUserByID(ctx context.Context, userID uuid.UUID) error {
	ret := _m.Called(ctx, userID)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) error); ok {
		r0 = rf(ctx, userID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UserStorage_DeleteUserByID_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteUserByID'
type UserStorage_DeleteUserByID_Call struct {
	*mock.Call
}

// DeleteUserByID is a helper method to define mock.On call
//   - ctx context.Context
//   - userID uuid.UUID
func (_e *UserStorage_Expecter) DeleteUserByID(ctx interface{}, userID interface{}) *UserStorage_DeleteUserByID_Call {
	return &UserStorage_DeleteUserByID_Call{Call: _e.mock.On("DeleteUserByID", ctx, userID)}
}

func (_c *UserStorage_DeleteUserByID_Call) Run(run func(ctx context.Context, userID uuid.UUID)) *UserStorage_DeleteUserByID_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uuid.UUID))
	})
	return _c
}

func (_c *UserStorage_DeleteUserByID_Call) Return(_a0 error) *UserStorage_DeleteUserByID_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *UserStorage_DeleteUserByID_Call) RunAndReturn(run func(context.Context, uuid.UUID) error) *UserStorage_DeleteUserByID_Call {
	_c.Call.Return(run)
	return _c
}

// GetUserByID provides a mock function with given fields: ctx, userID
func (_m *UserStorage) GetUserByID(ctx context.Context, userID uuid.UUID) (entity.User, error) {
	ret := _m.Called(ctx, userID)

	var r0 entity.User
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) (entity.User, error)); ok {
		return rf(ctx, userID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) entity.User); ok {
		r0 = rf(ctx, userID)
	} else {
		r0 = ret.Get(0).(entity.User)
	}

	if rf, ok := ret.Get(1).(func(context.Context, uuid.UUID) error); ok {
		r1 = rf(ctx, userID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UserStorage_GetUserByID_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetUserByID'
type UserStorage_GetUserByID_Call struct {
	*mock.Call
}

// GetUserByID is a helper method to define mock.On call
//   - ctx context.Context
//   - userID uuid.UUID
func (_e *UserStorage_Expecter) GetUserByID(ctx interface{}, userID interface{}) *UserStorage_GetUserByID_Call {
	return &UserStorage_GetUserByID_Call{Call: _e.mock.On("GetUserByID", ctx, userID)}
}

func (_c *UserStorage_GetUserByID_Call) Run(run func(ctx context.Context, userID uuid.UUID)) *UserStorage_GetUserByID_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uuid.UUID))
	})
	return _c
}

func (_c *UserStorage_GetUserByID_Call) Return(_a0 entity.User, _a1 error) *UserStorage_GetUserByID_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *UserStorage_GetUserByID_Call) RunAndReturn(run func(context.Context, uuid.UUID) (entity.User, error)) *UserStorage_GetUserByID_Call {
	_c.Call.Return(run)
	return _c
}

// GetUserList provides a mock function with given fields: ctx, filter
func (_m *UserStorage) GetUserList(ctx context.Context, filter entity.UserFilter) ([]entity.User, error) {
	ret := _m.Called(ctx, filter)

	var r0 []entity.User
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, entity.UserFilter) ([]entity.User, error)); ok {
		return rf(ctx, filter)
	}
	if rf, ok := ret.Get(0).(func(context.Context, entity.UserFilter) []entity.User); ok {
		r0 = rf(ctx, filter)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]entity.User)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, entity.UserFilter) error); ok {
		r1 = rf(ctx, filter)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UserStorage_GetUserList_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetUserList'
type UserStorage_GetUserList_Call struct {
	*mock.Call
}

// GetUserList is a helper method to define mock.On call
//   - ctx context.Context
//   - filter entity.UserFilter
func (_e *UserStorage_Expecter) GetUserList(ctx interface{}, filter interface{}) *UserStorage_GetUserList_Call {
	return &UserStorage_GetUserList_Call{Call: _e.mock.On("GetUserList", ctx, filter)}
}

func (_c *UserStorage_GetUserList_Call) Run(run func(ctx context.Context, filter entity.UserFilter)) *UserStorage_GetUserList_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(entity.UserFilter))
	})
	return _c
}

func (_c *UserStorage_GetUserList_Call) Return(_a0 []entity.User, _a1 error) *UserStorage_GetUserList_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *UserStorage_GetUserList_Call) RunAndReturn(run func(context.Context, entity.UserFilter) ([]entity.User, error)) *UserStorage_GetUserList_Call {
	_c.Call.Return(run)
	return _c
}

// UpdateUserByID provides a mock function with given fields: ctx, userID, user
func (_m *UserStorage) UpdateUserByID(ctx context.Context, userID uuid.UUID, user entity.User) (entity.User, error) {
	ret := _m.Called(ctx, userID, user)

	var r0 entity.User
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID, entity.User) (entity.User, error)); ok {
		return rf(ctx, userID, user)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID, entity.User) entity.User); ok {
		r0 = rf(ctx, userID, user)
	} else {
		r0 = ret.Get(0).(entity.User)
	}

	if rf, ok := ret.Get(1).(func(context.Context, uuid.UUID, entity.User) error); ok {
		r1 = rf(ctx, userID, user)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UserStorage_UpdateUserByID_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdateUserByID'
type UserStorage_UpdateUserByID_Call struct {
	*mock.Call
}

// UpdateUserByID is a helper method to define mock.On call
//   - ctx context.Context
//   - userID uuid.UUID
//   - user entity.User
func (_e *UserStorage_Expecter) UpdateUserByID(ctx interface{}, userID interface{}, user interface{}) *UserStorage_UpdateUserByID_Call {
	return &UserStorage_UpdateUserByID_Call{Call: _e.mock.On("UpdateUserByID", ctx, userID, user)}
}

func (_c *UserStorage_UpdateUserByID_Call) Run(run func(ctx context.Context, userID uuid.UUID, user entity.User)) *UserStorage_UpdateUserByID_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uuid.UUID), args[2].(entity.User))
	})
	return _c
}

func (_c *UserStorage_UpdateUserByID_Call) Return(_a0 entity.User, _a1 error) *UserStorage_UpdateUserByID_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *UserStorage_UpdateUserByID_Call) RunAndReturn(run func(context.Context, uuid.UUID, entity.User) (entity.User, error)) *UserStorage_UpdateUserByID_Call {
	_c.Call.Return(run)
	return _c
}

// NewUserStorage creates a new instance of UserStorage. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewUserStorage(t interface {
	mock.TestingT
	Cleanup(func())
}) *UserStorage {
	mock := &UserStorage{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
