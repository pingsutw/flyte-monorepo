// Code generated by mockery v1.0.1. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// CacheKey is an autogenerated mock type for the CacheKey type
type CacheKey struct {
	mock.Mock
}

type CacheKey_String struct {
	*mock.Call
}

func (_m CacheKey_String) Return(_a0 string) *CacheKey_String {
	return &CacheKey_String{Call: _m.Call.Return(_a0)}
}

func (_m *CacheKey) OnString() *CacheKey_String {
	c_call := _m.On("String")
	return &CacheKey_String{Call: c_call}
}

func (_m *CacheKey) OnStringMatch(matchers ...interface{}) *CacheKey_String {
	c_call := _m.On("String", matchers...)
	return &CacheKey_String{Call: c_call}
}

// String provides a mock function with given fields:
func (_m *CacheKey) String() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}
