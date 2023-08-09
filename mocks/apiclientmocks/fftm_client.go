// Code generated by mockery v2.32.0. DO NOT EDIT.

package apiclientmocks

import (
	context "context"

	apitypes "github.com/hyperledger/firefly-transaction-manager/pkg/apitypes"

	mock "github.com/stretchr/testify/mock"
)

// FFTMClient is an autogenerated mock type for the FFTMClient type
type FFTMClient struct {
	mock.Mock
}

// DeleteEventStream provides a mock function with given fields: ctx, eventStreamID
func (_m *FFTMClient) DeleteEventStream(ctx context.Context, eventStreamID string) error {
	ret := _m.Called(ctx, eventStreamID)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, eventStreamID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteEventStreamsByName provides a mock function with given fields: ctx, nameRegex
func (_m *FFTMClient) DeleteEventStreamsByName(ctx context.Context, nameRegex string) error {
	ret := _m.Called(ctx, nameRegex)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, nameRegex)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteListener provides a mock function with given fields: ctx, eventStreamID, listenerID
func (_m *FFTMClient) DeleteListener(ctx context.Context, eventStreamID string, listenerID string) error {
	ret := _m.Called(ctx, eventStreamID, listenerID)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) error); ok {
		r0 = rf(ctx, eventStreamID, listenerID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteListenersByName provides a mock function with given fields: ctx, eventStreamID, nameRegex
func (_m *FFTMClient) DeleteListenersByName(ctx context.Context, eventStreamID string, nameRegex string) error {
	ret := _m.Called(ctx, eventStreamID, nameRegex)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) error); ok {
		r0 = rf(ctx, eventStreamID, nameRegex)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetEventStreams provides a mock function with given fields: ctx
func (_m *FFTMClient) GetEventStreams(ctx context.Context) ([]apitypes.EventStream, error) {
	ret := _m.Called(ctx)

	var r0 []apitypes.EventStream
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) ([]apitypes.EventStream, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) []apitypes.EventStream); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]apitypes.EventStream)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetListeners provides a mock function with given fields: ctx, eventStreamID
func (_m *FFTMClient) GetListeners(ctx context.Context, eventStreamID string) ([]apitypes.Listener, error) {
	ret := _m.Called(ctx, eventStreamID)

	var r0 []apitypes.Listener
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) ([]apitypes.Listener, error)); ok {
		return rf(ctx, eventStreamID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) []apitypes.Listener); ok {
		r0 = rf(ctx, eventStreamID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]apitypes.Listener)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, eventStreamID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewFFTMClient creates a new instance of FFTMClient. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewFFTMClient(t interface {
	mock.TestingT
	Cleanup(func())
}) *FFTMClient {
	mock := &FFTMClient{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
