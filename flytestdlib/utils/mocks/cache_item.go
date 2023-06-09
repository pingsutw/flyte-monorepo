// Code generated by mockery v1.0.1. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// CacheItem is an autogenerated mock type for the CacheItem type
type CacheItem struct {
	mock.Mock
}

type CacheItem_ID struct {
	*mock.Call
}

func (_m CacheItem_ID) Return(_a0 string) *CacheItem_ID {
	return &CacheItem_ID{Call: _m.Call.Return(_a0)}
}

func (_m *CacheItem) OnID() *CacheItem_ID {
	c_call := _m.On("ID")
	return &CacheItem_ID{Call: c_call}
}

func (_m *CacheItem) OnIDMatch(matchers ...interface{}) *CacheItem_ID {
	c_call := _m.On("ID", matchers...)
	return &CacheItem_ID{Call: c_call}
}

// ID provides a mock function with given fields:
func (_m *CacheItem) ID() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}
