// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	context "context"

	compile "github.com/charithe/menshen/pkg/compile"

	mock "github.com/stretchr/testify/mock"

	policyv1 "github.com/charithe/menshen/pkg/generated/policy/v1"
)

// Index is an autogenerated mock type for the Index type
type Index struct {
	mock.Mock
}

// Add provides a mock function with given fields: _a0, _a1
func (_m *Index) Add(_a0 string, _a1 *policyv1.Policy) (*compile.Incremental, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *compile.Incremental
	if rf, ok := ret.Get(0).(func(string, *policyv1.Policy) *compile.Incremental); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*compile.Incremental)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, *policyv1.Policy) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FilenameFor provides a mock function with given fields: _a0
func (_m *Index) FilenameFor(_a0 *policyv1.Policy) string {
	ret := _m.Called(_a0)

	var r0 string
	if rf, ok := ret.Get(0).(func(*policyv1.Policy) string); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// GetAllPolicies provides a mock function with given fields: _a0
func (_m *Index) GetAllPolicies(_a0 context.Context) <-chan *compile.Unit {
	ret := _m.Called(_a0)

	var r0 <-chan *compile.Unit
	if rf, ok := ret.Get(0).(func(context.Context) <-chan *compile.Unit); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(<-chan *compile.Unit)
		}
	}

	return r0
}

// Remove provides a mock function with given fields: _a0
func (_m *Index) Remove(_a0 string) (*compile.Incremental, error) {
	ret := _m.Called(_a0)

	var r0 *compile.Incremental
	if rf, ok := ret.Get(0).(func(string) *compile.Incremental); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*compile.Incremental)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RemoveIfSafe provides a mock function with given fields: _a0
func (_m *Index) RemoveIfSafe(_a0 string) (*compile.Incremental, error) {
	ret := _m.Called(_a0)

	var r0 *compile.Incremental
	if rf, ok := ret.Get(0).(func(string) *compile.Incremental); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*compile.Incremental)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
