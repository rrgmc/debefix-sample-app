// Code generated by mockery v2.35.4. DO NOT EDIT.

package mocks

import (
	context "context"

	entity "github.com/RangelReale/debefix-sample-app/internal/entity"
	mock "github.com/stretchr/testify/mock"

	uuid "github.com/google/uuid"
)

// TagStorage is an autogenerated mock type for the TagStorage type
type TagStorage struct {
	mock.Mock
}

type TagStorage_Expecter struct {
	mock *mock.Mock
}

func (_m *TagStorage) EXPECT() *TagStorage_Expecter {
	return &TagStorage_Expecter{mock: &_m.Mock}
}

// AddTag provides a mock function with given fields: ctx, tag
func (_m *TagStorage) AddTag(ctx context.Context, tag entity.Tag) (entity.Tag, error) {
	ret := _m.Called(ctx, tag)

	var r0 entity.Tag
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, entity.Tag) (entity.Tag, error)); ok {
		return rf(ctx, tag)
	}
	if rf, ok := ret.Get(0).(func(context.Context, entity.Tag) entity.Tag); ok {
		r0 = rf(ctx, tag)
	} else {
		r0 = ret.Get(0).(entity.Tag)
	}

	if rf, ok := ret.Get(1).(func(context.Context, entity.Tag) error); ok {
		r1 = rf(ctx, tag)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// TagStorage_AddTag_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'AddTag'
type TagStorage_AddTag_Call struct {
	*mock.Call
}

// AddTag is a helper method to define mock.On call
//   - ctx context.Context
//   - tag entity.Tag
func (_e *TagStorage_Expecter) AddTag(ctx interface{}, tag interface{}) *TagStorage_AddTag_Call {
	return &TagStorage_AddTag_Call{Call: _e.mock.On("AddTag", ctx, tag)}
}

func (_c *TagStorage_AddTag_Call) Run(run func(ctx context.Context, tag entity.Tag)) *TagStorage_AddTag_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(entity.Tag))
	})
	return _c
}

func (_c *TagStorage_AddTag_Call) Return(_a0 entity.Tag, _a1 error) *TagStorage_AddTag_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *TagStorage_AddTag_Call) RunAndReturn(run func(context.Context, entity.Tag) (entity.Tag, error)) *TagStorage_AddTag_Call {
	_c.Call.Return(run)
	return _c
}

// DeleteTagByID provides a mock function with given fields: ctx, tagID
func (_m *TagStorage) DeleteTagByID(ctx context.Context, tagID uuid.UUID) error {
	ret := _m.Called(ctx, tagID)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) error); ok {
		r0 = rf(ctx, tagID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// TagStorage_DeleteTagByID_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteTagByID'
type TagStorage_DeleteTagByID_Call struct {
	*mock.Call
}

// DeleteTagByID is a helper method to define mock.On call
//   - ctx context.Context
//   - tagID uuid.UUID
func (_e *TagStorage_Expecter) DeleteTagByID(ctx interface{}, tagID interface{}) *TagStorage_DeleteTagByID_Call {
	return &TagStorage_DeleteTagByID_Call{Call: _e.mock.On("DeleteTagByID", ctx, tagID)}
}

func (_c *TagStorage_DeleteTagByID_Call) Run(run func(ctx context.Context, tagID uuid.UUID)) *TagStorage_DeleteTagByID_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uuid.UUID))
	})
	return _c
}

func (_c *TagStorage_DeleteTagByID_Call) Return(_a0 error) *TagStorage_DeleteTagByID_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *TagStorage_DeleteTagByID_Call) RunAndReturn(run func(context.Context, uuid.UUID) error) *TagStorage_DeleteTagByID_Call {
	_c.Call.Return(run)
	return _c
}

// GetTagByID provides a mock function with given fields: ctx, tagID
func (_m *TagStorage) GetTagByID(ctx context.Context, tagID uuid.UUID) (entity.Tag, error) {
	ret := _m.Called(ctx, tagID)

	var r0 entity.Tag
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) (entity.Tag, error)); ok {
		return rf(ctx, tagID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) entity.Tag); ok {
		r0 = rf(ctx, tagID)
	} else {
		r0 = ret.Get(0).(entity.Tag)
	}

	if rf, ok := ret.Get(1).(func(context.Context, uuid.UUID) error); ok {
		r1 = rf(ctx, tagID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// TagStorage_GetTagByID_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetTagByID'
type TagStorage_GetTagByID_Call struct {
	*mock.Call
}

// GetTagByID is a helper method to define mock.On call
//   - ctx context.Context
//   - tagID uuid.UUID
func (_e *TagStorage_Expecter) GetTagByID(ctx interface{}, tagID interface{}) *TagStorage_GetTagByID_Call {
	return &TagStorage_GetTagByID_Call{Call: _e.mock.On("GetTagByID", ctx, tagID)}
}

func (_c *TagStorage_GetTagByID_Call) Run(run func(ctx context.Context, tagID uuid.UUID)) *TagStorage_GetTagByID_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uuid.UUID))
	})
	return _c
}

func (_c *TagStorage_GetTagByID_Call) Return(_a0 entity.Tag, _a1 error) *TagStorage_GetTagByID_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *TagStorage_GetTagByID_Call) RunAndReturn(run func(context.Context, uuid.UUID) (entity.Tag, error)) *TagStorage_GetTagByID_Call {
	_c.Call.Return(run)
	return _c
}

// GetTags provides a mock function with given fields: ctx, filter
func (_m *TagStorage) GetTags(ctx context.Context, filter entity.TagsFilter) ([]entity.Tag, error) {
	ret := _m.Called(ctx, filter)

	var r0 []entity.Tag
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, entity.TagsFilter) ([]entity.Tag, error)); ok {
		return rf(ctx, filter)
	}
	if rf, ok := ret.Get(0).(func(context.Context, entity.TagsFilter) []entity.Tag); ok {
		r0 = rf(ctx, filter)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]entity.Tag)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, entity.TagsFilter) error); ok {
		r1 = rf(ctx, filter)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// TagStorage_GetTags_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetTags'
type TagStorage_GetTags_Call struct {
	*mock.Call
}

// GetTags is a helper method to define mock.On call
//   - ctx context.Context
//   - filter entity.TagsFilter
func (_e *TagStorage_Expecter) GetTags(ctx interface{}, filter interface{}) *TagStorage_GetTags_Call {
	return &TagStorage_GetTags_Call{Call: _e.mock.On("GetTags", ctx, filter)}
}

func (_c *TagStorage_GetTags_Call) Run(run func(ctx context.Context, filter entity.TagsFilter)) *TagStorage_GetTags_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(entity.TagsFilter))
	})
	return _c
}

func (_c *TagStorage_GetTags_Call) Return(_a0 []entity.Tag, _a1 error) *TagStorage_GetTags_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *TagStorage_GetTags_Call) RunAndReturn(run func(context.Context, entity.TagsFilter) ([]entity.Tag, error)) *TagStorage_GetTags_Call {
	_c.Call.Return(run)
	return _c
}

// UpdateTagByID provides a mock function with given fields: ctx, tagID, tag
func (_m *TagStorage) UpdateTagByID(ctx context.Context, tagID uuid.UUID, tag entity.Tag) (entity.Tag, error) {
	ret := _m.Called(ctx, tagID, tag)

	var r0 entity.Tag
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID, entity.Tag) (entity.Tag, error)); ok {
		return rf(ctx, tagID, tag)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID, entity.Tag) entity.Tag); ok {
		r0 = rf(ctx, tagID, tag)
	} else {
		r0 = ret.Get(0).(entity.Tag)
	}

	if rf, ok := ret.Get(1).(func(context.Context, uuid.UUID, entity.Tag) error); ok {
		r1 = rf(ctx, tagID, tag)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// TagStorage_UpdateTagByID_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdateTagByID'
type TagStorage_UpdateTagByID_Call struct {
	*mock.Call
}

// UpdateTagByID is a helper method to define mock.On call
//   - ctx context.Context
//   - tagID uuid.UUID
//   - tag entity.Tag
func (_e *TagStorage_Expecter) UpdateTagByID(ctx interface{}, tagID interface{}, tag interface{}) *TagStorage_UpdateTagByID_Call {
	return &TagStorage_UpdateTagByID_Call{Call: _e.mock.On("UpdateTagByID", ctx, tagID, tag)}
}

func (_c *TagStorage_UpdateTagByID_Call) Run(run func(ctx context.Context, tagID uuid.UUID, tag entity.Tag)) *TagStorage_UpdateTagByID_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uuid.UUID), args[2].(entity.Tag))
	})
	return _c
}

func (_c *TagStorage_UpdateTagByID_Call) Return(_a0 entity.Tag, _a1 error) *TagStorage_UpdateTagByID_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *TagStorage_UpdateTagByID_Call) RunAndReturn(run func(context.Context, uuid.UUID, entity.Tag) (entity.Tag, error)) *TagStorage_UpdateTagByID_Call {
	_c.Call.Return(run)
	return _c
}

// NewTagStorage creates a new instance of TagStorage. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewTagStorage(t interface {
	mock.TestingT
	Cleanup(func())
}) *TagStorage {
	mock := &TagStorage{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}