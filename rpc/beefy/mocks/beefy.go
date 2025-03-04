// Code generated by mockery v2.13.0-beta.1. DO NOT EDIT.

package mocks

import (
	beefy "github.com/kirill-sr/go-substrate-rpc-client/v4/rpc/beefy"
	mock "github.com/stretchr/testify/mock"

	types "github.com/kirill-sr/go-substrate-rpc-client/v4/types"
)

// Beefy is an autogenerated mock type for the Beefy type
type Beefy struct {
	mock.Mock
}

// GetFinalizedHead provides a mock function with given fields:
func (_m *Beefy) GetFinalizedHead() (types.Hash, error) {
	ret := _m.Called()

	var r0 types.Hash
	if rf, ok := ret.Get(0).(func() types.Hash); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(types.Hash)
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

// SubscribeJustifications provides a mock function with given fields:
func (_m *Beefy) SubscribeJustifications() (*beefy.JustificationsSubscription, error) {
	ret := _m.Called()

	var r0 *beefy.JustificationsSubscription
	if rf, ok := ret.Get(0).(func() *beefy.JustificationsSubscription); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*beefy.JustificationsSubscription)
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

type NewBeefyT interface {
	mock.TestingT
	Cleanup(func())
}

// NewBeefy creates a new instance of Beefy. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewBeefy(t NewBeefyT) *Beefy {
	mock := &Beefy{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
