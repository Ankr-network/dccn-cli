// Code generated by mockery v1.0.0. DO NOT EDIT.

// Generated: please do not edit by hand

package mocks

import do "github.com/Ankr-network/dccn-cli/do"
import mock "github.com/stretchr/testify/mock"

// RegionsService is an autogenerated mock type for the RegionsService type
type RegionsService struct {
	mock.Mock
}

// List provides a mock function with given fields:
func (_m *RegionsService) List() (do.Regions, error) {
	ret := _m.Called()

	var r0 do.Regions
	if rf, ok := ret.Get(0).(func() do.Regions); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(do.Regions)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
