// Code generated by mockery. DO NOT EDIT.

package mocks

import (
	context "context"

	entity "github.com/abc-valera/netsly-golang/internal/domain/entity"
	mock "github.com/stretchr/testify/mock"

	model "github.com/abc-valera/netsly-golang/internal/domain/model"

	selector "github.com/abc-valera/netsly-golang/internal/domain/persistence/query/queryUtil/selector"
)

// Room is an autogenerated mock type for the IRoom type
type Room struct {
	mock.Mock
}

type Room_Expecter struct {
	mock *mock.Mock
}

func (_m *Room) EXPECT() *Room_Expecter {
	return &Room_Expecter{mock: &_m.Mock}
}

// Create provides a mock function with given fields: ctx, req
func (_m *Room) Create(ctx context.Context, req entity.RoomCreateRequest) (model.Room, error) {
	ret := _m.Called(ctx, req)

	if len(ret) == 0 {
		panic("no return value specified for Create")
	}

	var r0 model.Room
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, entity.RoomCreateRequest) (model.Room, error)); ok {
		return rf(ctx, req)
	}
	if rf, ok := ret.Get(0).(func(context.Context, entity.RoomCreateRequest) model.Room); ok {
		r0 = rf(ctx, req)
	} else {
		r0 = ret.Get(0).(model.Room)
	}

	if rf, ok := ret.Get(1).(func(context.Context, entity.RoomCreateRequest) error); ok {
		r1 = rf(ctx, req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Room_Create_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Create'
type Room_Create_Call struct {
	*mock.Call
}

// Create is a helper method to define mock.On call
//   - ctx context.Context
//   - req entity.RoomCreateRequest
func (_e *Room_Expecter) Create(ctx interface{}, req interface{}) *Room_Create_Call {
	return &Room_Create_Call{Call: _e.mock.On("Create", ctx, req)}
}

func (_c *Room_Create_Call) Run(run func(ctx context.Context, req entity.RoomCreateRequest)) *Room_Create_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(entity.RoomCreateRequest))
	})
	return _c
}

func (_c *Room_Create_Call) Return(_a0 model.Room, _a1 error) *Room_Create_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Room_Create_Call) RunAndReturn(run func(context.Context, entity.RoomCreateRequest) (model.Room, error)) *Room_Create_Call {
	_c.Call.Return(run)
	return _c
}

// Delete provides a mock function with given fields: ctx, roomID
func (_m *Room) Delete(ctx context.Context, roomID string) error {
	ret := _m.Called(ctx, roomID)

	if len(ret) == 0 {
		panic("no return value specified for Delete")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, roomID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Room_Delete_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Delete'
type Room_Delete_Call struct {
	*mock.Call
}

// Delete is a helper method to define mock.On call
//   - ctx context.Context
//   - roomID string
func (_e *Room_Expecter) Delete(ctx interface{}, roomID interface{}) *Room_Delete_Call {
	return &Room_Delete_Call{Call: _e.mock.On("Delete", ctx, roomID)}
}

func (_c *Room_Delete_Call) Run(run func(ctx context.Context, roomID string)) *Room_Delete_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *Room_Delete_Call) Return(_a0 error) *Room_Delete_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Room_Delete_Call) RunAndReturn(run func(context.Context, string) error) *Room_Delete_Call {
	_c.Call.Return(run)
	return _c
}

// Get provides a mock function with given fields: _a0, _a1
func (_m *Room) Get(_a0 context.Context, _a1 model.Room) (model.Room, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for Get")
	}

	var r0 model.Room
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, model.Room) (model.Room, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, model.Room) model.Room); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Get(0).(model.Room)
	}

	if rf, ok := ret.Get(1).(func(context.Context, model.Room) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Room_Get_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Get'
type Room_Get_Call struct {
	*mock.Call
}

// Get is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 model.Room
func (_e *Room_Expecter) Get(_a0 interface{}, _a1 interface{}) *Room_Get_Call {
	return &Room_Get_Call{Call: _e.mock.On("Get", _a0, _a1)}
}

func (_c *Room_Get_Call) Run(run func(_a0 context.Context, _a1 model.Room)) *Room_Get_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(model.Room))
	})
	return _c
}

func (_c *Room_Get_Call) Return(_a0 model.Room, _a1 error) *Room_Get_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Room_Get_Call) RunAndReturn(run func(context.Context, model.Room) (model.Room, error)) *Room_Get_Call {
	_c.Call.Return(run)
	return _c
}

// GetMany provides a mock function with given fields: _a0, _a1
func (_m *Room) GetMany(_a0 context.Context, _a1 ...selector.Option[model.Room]) ([]model.Room, error) {
	_va := make([]interface{}, len(_a1))
	for _i := range _a1 {
		_va[_i] = _a1[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetMany")
	}

	var r0 []model.Room
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, ...selector.Option[model.Room]) ([]model.Room, error)); ok {
		return rf(_a0, _a1...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, ...selector.Option[model.Room]) []model.Room); ok {
		r0 = rf(_a0, _a1...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.Room)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, ...selector.Option[model.Room]) error); ok {
		r1 = rf(_a0, _a1...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Room_GetMany_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetMany'
type Room_GetMany_Call struct {
	*mock.Call
}

// GetMany is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 ...selector.Option[model.Room]
func (_e *Room_Expecter) GetMany(_a0 interface{}, _a1 ...interface{}) *Room_GetMany_Call {
	return &Room_GetMany_Call{Call: _e.mock.On("GetMany",
		append([]interface{}{_a0}, _a1...)...)}
}

func (_c *Room_GetMany_Call) Run(run func(_a0 context.Context, _a1 ...selector.Option[model.Room])) *Room_GetMany_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]selector.Option[model.Room], len(args)-1)
		for i, a := range args[1:] {
			if a != nil {
				variadicArgs[i] = a.(selector.Option[model.Room])
			}
		}
		run(args[0].(context.Context), variadicArgs...)
	})
	return _c
}

func (_c *Room_GetMany_Call) Return(_a0 []model.Room, _a1 error) *Room_GetMany_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Room_GetMany_Call) RunAndReturn(run func(context.Context, ...selector.Option[model.Room]) ([]model.Room, error)) *Room_GetMany_Call {
	_c.Call.Return(run)
	return _c
}

// Update provides a mock function with given fields: ctx, roomID, req
func (_m *Room) Update(ctx context.Context, roomID string, req entity.RoomUpdateRequest) (model.Room, error) {
	ret := _m.Called(ctx, roomID, req)

	if len(ret) == 0 {
		panic("no return value specified for Update")
	}

	var r0 model.Room
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, entity.RoomUpdateRequest) (model.Room, error)); ok {
		return rf(ctx, roomID, req)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, entity.RoomUpdateRequest) model.Room); ok {
		r0 = rf(ctx, roomID, req)
	} else {
		r0 = ret.Get(0).(model.Room)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, entity.RoomUpdateRequest) error); ok {
		r1 = rf(ctx, roomID, req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Room_Update_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Update'
type Room_Update_Call struct {
	*mock.Call
}

// Update is a helper method to define mock.On call
//   - ctx context.Context
//   - roomID string
//   - req entity.RoomUpdateRequest
func (_e *Room_Expecter) Update(ctx interface{}, roomID interface{}, req interface{}) *Room_Update_Call {
	return &Room_Update_Call{Call: _e.mock.On("Update", ctx, roomID, req)}
}

func (_c *Room_Update_Call) Run(run func(ctx context.Context, roomID string, req entity.RoomUpdateRequest)) *Room_Update_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(entity.RoomUpdateRequest))
	})
	return _c
}

func (_c *Room_Update_Call) Return(_a0 model.Room, _a1 error) *Room_Update_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Room_Update_Call) RunAndReturn(run func(context.Context, string, entity.RoomUpdateRequest) (model.Room, error)) *Room_Update_Call {
	_c.Call.Return(run)
	return _c
}

// NewRoom creates a new instance of Room. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewRoom(t interface {
	mock.TestingT
	Cleanup(func())
}) *Room {
	mock := &Room{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
