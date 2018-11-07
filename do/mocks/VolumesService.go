// Code generated by mockery v1.0.0. DO NOT EDIT.

// Generated: please do not edit by hand

package mocks

import do "github.com/Ankr-network/dccn-cli/do"
import godo "github.com/Ankr-network/godo"
import mock "github.com/stretchr/testify/mock"

// VolumesService is an autogenerated mock type for the VolumesService type
type VolumesService struct {
	mock.Mock
}

// CreateSnapshot provides a mock function with given fields: _a0
func (_m *VolumesService) CreateSnapshot(_a0 *godo.SnapshotCreateRequest) (*do.Snapshot, error) {
	ret := _m.Called(_a0)

	var r0 *do.Snapshot
	if rf, ok := ret.Get(0).(func(*godo.SnapshotCreateRequest) *do.Snapshot); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*do.Snapshot)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*godo.SnapshotCreateRequest) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateVolume provides a mock function with given fields: _a0
func (_m *VolumesService) CreateVolume(_a0 *godo.VolumeCreateRequest) (*do.Volume, error) {
	ret := _m.Called(_a0)

	var r0 *do.Volume
	if rf, ok := ret.Get(0).(func(*godo.VolumeCreateRequest) *do.Volume); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*do.Volume)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*godo.VolumeCreateRequest) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteSnapshot provides a mock function with given fields: _a0
func (_m *VolumesService) DeleteSnapshot(_a0 string) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteVolume provides a mock function with given fields: _a0
func (_m *VolumesService) DeleteVolume(_a0 string) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Get provides a mock function with given fields: _a0
func (_m *VolumesService) Get(_a0 string) (*do.Volume, error) {
	ret := _m.Called(_a0)

	var r0 *do.Volume
	if rf, ok := ret.Get(0).(func(string) *do.Volume); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*do.Volume)
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

// GetSnapshot provides a mock function with given fields: _a0
func (_m *VolumesService) GetSnapshot(_a0 string) (*do.Snapshot, error) {
	ret := _m.Called(_a0)

	var r0 *do.Snapshot
	if rf, ok := ret.Get(0).(func(string) *do.Snapshot); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*do.Snapshot)
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

// List provides a mock function with given fields:
func (_m *VolumesService) List() ([]do.Volume, error) {
	ret := _m.Called()

	var r0 []do.Volume
	if rf, ok := ret.Get(0).(func() []do.Volume); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]do.Volume)
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

// ListSnapshots provides a mock function with given fields: _a0, _a1
func (_m *VolumesService) ListSnapshots(_a0 string, _a1 *godo.ListOptions) ([]do.Snapshot, error) {
	ret := _m.Called(_a0, _a1)

	var r0 []do.Snapshot
	if rf, ok := ret.Get(0).(func(string, *godo.ListOptions) []do.Snapshot); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]do.Snapshot)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, *godo.ListOptions) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
