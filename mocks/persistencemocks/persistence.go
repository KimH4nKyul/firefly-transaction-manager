// Code generated by mockery v1.0.0. DO NOT EDIT.

package persistencemocks

import (
	context "context"

	apitypes "github.com/hyperledger/firefly-transaction-manager/pkg/apitypes"

	fftypes "github.com/hyperledger/firefly-common/pkg/fftypes"

	mock "github.com/stretchr/testify/mock"
)

// Persistence is an autogenerated mock type for the Persistence type
type Persistence struct {
	mock.Mock
}

// Close provides a mock function with given fields: ctx
func (_m *Persistence) Close(ctx context.Context) {
	_m.Called(ctx)
}

// DeleteCheckpoint provides a mock function with given fields: ctx, streamID
func (_m *Persistence) DeleteCheckpoint(ctx context.Context, streamID *fftypes.UUID) error {
	ret := _m.Called(ctx, streamID)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *fftypes.UUID) error); ok {
		r0 = rf(ctx, streamID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteListener provides a mock function with given fields: ctx, listenerID
func (_m *Persistence) DeleteListener(ctx context.Context, listenerID *fftypes.UUID) error {
	ret := _m.Called(ctx, listenerID)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *fftypes.UUID) error); ok {
		r0 = rf(ctx, listenerID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteStream provides a mock function with given fields: ctx, streamID
func (_m *Persistence) DeleteStream(ctx context.Context, streamID *fftypes.UUID) error {
	ret := _m.Called(ctx, streamID)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *fftypes.UUID) error); ok {
		r0 = rf(ctx, streamID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetCheckpoint provides a mock function with given fields: ctx, streamID
func (_m *Persistence) GetCheckpoint(ctx context.Context, streamID *fftypes.UUID) (*apitypes.EventStreamCheckpoint, error) {
	ret := _m.Called(ctx, streamID)

	var r0 *apitypes.EventStreamCheckpoint
	if rf, ok := ret.Get(0).(func(context.Context, *fftypes.UUID) *apitypes.EventStreamCheckpoint); ok {
		r0 = rf(ctx, streamID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*apitypes.EventStreamCheckpoint)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *fftypes.UUID) error); ok {
		r1 = rf(ctx, streamID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetListener provides a mock function with given fields: ctx, listenerID
func (_m *Persistence) GetListener(ctx context.Context, listenerID *fftypes.UUID) (*apitypes.Listener, error) {
	ret := _m.Called(ctx, listenerID)

	var r0 *apitypes.Listener
	if rf, ok := ret.Get(0).(func(context.Context, *fftypes.UUID) *apitypes.Listener); ok {
		r0 = rf(ctx, listenerID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*apitypes.Listener)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *fftypes.UUID) error); ok {
		r1 = rf(ctx, listenerID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetStream provides a mock function with given fields: ctx, streamID
func (_m *Persistence) GetStream(ctx context.Context, streamID *fftypes.UUID) (*apitypes.EventStream, error) {
	ret := _m.Called(ctx, streamID)

	var r0 *apitypes.EventStream
	if rf, ok := ret.Get(0).(func(context.Context, *fftypes.UUID) *apitypes.EventStream); ok {
		r0 = rf(ctx, streamID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*apitypes.EventStream)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *fftypes.UUID) error); ok {
		r1 = rf(ctx, streamID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListListeners provides a mock function with given fields: ctx, after, limit
func (_m *Persistence) ListListeners(ctx context.Context, after *fftypes.UUID, limit int) ([]*apitypes.Listener, error) {
	ret := _m.Called(ctx, after, limit)

	var r0 []*apitypes.Listener
	if rf, ok := ret.Get(0).(func(context.Context, *fftypes.UUID, int) []*apitypes.Listener); ok {
		r0 = rf(ctx, after, limit)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*apitypes.Listener)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *fftypes.UUID, int) error); ok {
		r1 = rf(ctx, after, limit)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListStreamListeners provides a mock function with given fields: ctx, after, limit, streamID
func (_m *Persistence) ListStreamListeners(ctx context.Context, after *fftypes.UUID, limit int, streamID *fftypes.UUID) ([]*apitypes.Listener, error) {
	ret := _m.Called(ctx, after, limit, streamID)

	var r0 []*apitypes.Listener
	if rf, ok := ret.Get(0).(func(context.Context, *fftypes.UUID, int, *fftypes.UUID) []*apitypes.Listener); ok {
		r0 = rf(ctx, after, limit, streamID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*apitypes.Listener)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *fftypes.UUID, int, *fftypes.UUID) error); ok {
		r1 = rf(ctx, after, limit, streamID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListStreams provides a mock function with given fields: ctx, after, limit
func (_m *Persistence) ListStreams(ctx context.Context, after *fftypes.UUID, limit int) ([]*apitypes.EventStream, error) {
	ret := _m.Called(ctx, after, limit)

	var r0 []*apitypes.EventStream
	if rf, ok := ret.Get(0).(func(context.Context, *fftypes.UUID, int) []*apitypes.EventStream); ok {
		r0 = rf(ctx, after, limit)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*apitypes.EventStream)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *fftypes.UUID, int) error); ok {
		r1 = rf(ctx, after, limit)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// WriteCheckpoint provides a mock function with given fields: ctx, checkpoint
func (_m *Persistence) WriteCheckpoint(ctx context.Context, checkpoint *apitypes.EventStreamCheckpoint) error {
	ret := _m.Called(ctx, checkpoint)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *apitypes.EventStreamCheckpoint) error); ok {
		r0 = rf(ctx, checkpoint)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// WriteListener provides a mock function with given fields: ctx, spec
func (_m *Persistence) WriteListener(ctx context.Context, spec *apitypes.Listener) error {
	ret := _m.Called(ctx, spec)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *apitypes.Listener) error); ok {
		r0 = rf(ctx, spec)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// WriteStream provides a mock function with given fields: ctx, spec
func (_m *Persistence) WriteStream(ctx context.Context, spec *apitypes.EventStream) error {
	ret := _m.Called(ctx, spec)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *apitypes.EventStream) error); ok {
		r0 = rf(ctx, spec)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
