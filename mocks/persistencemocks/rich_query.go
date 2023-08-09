// Code generated by mockery v2.32.0. DO NOT EDIT.

package persistencemocks

import (
	context "context"

	apitypes "github.com/hyperledger/firefly-transaction-manager/pkg/apitypes"

	ffapi "github.com/hyperledger/firefly-common/pkg/ffapi"

	fftypes "github.com/hyperledger/firefly-common/pkg/fftypes"

	mock "github.com/stretchr/testify/mock"
)

// RichQuery is an autogenerated mock type for the RichQuery type
type RichQuery struct {
	mock.Mock
}

// ListListeners provides a mock function with given fields: ctx, filter
func (_m *RichQuery) ListListeners(ctx context.Context, filter ffapi.AndFilter) ([]*apitypes.Listener, *ffapi.FilterResult, error) {
	ret := _m.Called(ctx, filter)

	var r0 []*apitypes.Listener
	var r1 *ffapi.FilterResult
	var r2 error
	if rf, ok := ret.Get(0).(func(context.Context, ffapi.AndFilter) ([]*apitypes.Listener, *ffapi.FilterResult, error)); ok {
		return rf(ctx, filter)
	}
	if rf, ok := ret.Get(0).(func(context.Context, ffapi.AndFilter) []*apitypes.Listener); ok {
		r0 = rf(ctx, filter)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*apitypes.Listener)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, ffapi.AndFilter) *ffapi.FilterResult); ok {
		r1 = rf(ctx, filter)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ffapi.FilterResult)
		}
	}

	if rf, ok := ret.Get(2).(func(context.Context, ffapi.AndFilter) error); ok {
		r2 = rf(ctx, filter)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// ListStreamListeners provides a mock function with given fields: ctx, streamID, filter
func (_m *RichQuery) ListStreamListeners(ctx context.Context, streamID *fftypes.UUID, filter ffapi.AndFilter) ([]*apitypes.Listener, *ffapi.FilterResult, error) {
	ret := _m.Called(ctx, streamID, filter)

	var r0 []*apitypes.Listener
	var r1 *ffapi.FilterResult
	var r2 error
	if rf, ok := ret.Get(0).(func(context.Context, *fftypes.UUID, ffapi.AndFilter) ([]*apitypes.Listener, *ffapi.FilterResult, error)); ok {
		return rf(ctx, streamID, filter)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *fftypes.UUID, ffapi.AndFilter) []*apitypes.Listener); ok {
		r0 = rf(ctx, streamID, filter)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*apitypes.Listener)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *fftypes.UUID, ffapi.AndFilter) *ffapi.FilterResult); ok {
		r1 = rf(ctx, streamID, filter)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ffapi.FilterResult)
		}
	}

	if rf, ok := ret.Get(2).(func(context.Context, *fftypes.UUID, ffapi.AndFilter) error); ok {
		r2 = rf(ctx, streamID, filter)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// ListStreams provides a mock function with given fields: ctx, filter
func (_m *RichQuery) ListStreams(ctx context.Context, filter ffapi.AndFilter) ([]*apitypes.EventStream, *ffapi.FilterResult, error) {
	ret := _m.Called(ctx, filter)

	var r0 []*apitypes.EventStream
	var r1 *ffapi.FilterResult
	var r2 error
	if rf, ok := ret.Get(0).(func(context.Context, ffapi.AndFilter) ([]*apitypes.EventStream, *ffapi.FilterResult, error)); ok {
		return rf(ctx, filter)
	}
	if rf, ok := ret.Get(0).(func(context.Context, ffapi.AndFilter) []*apitypes.EventStream); ok {
		r0 = rf(ctx, filter)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*apitypes.EventStream)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, ffapi.AndFilter) *ffapi.FilterResult); ok {
		r1 = rf(ctx, filter)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ffapi.FilterResult)
		}
	}

	if rf, ok := ret.Get(2).(func(context.Context, ffapi.AndFilter) error); ok {
		r2 = rf(ctx, filter)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// ListTransactionConfirmations provides a mock function with given fields: ctx, txID, filter
func (_m *RichQuery) ListTransactionConfirmations(ctx context.Context, txID string, filter ffapi.AndFilter) ([]*apitypes.ConfirmationRecord, *ffapi.FilterResult, error) {
	ret := _m.Called(ctx, txID, filter)

	var r0 []*apitypes.ConfirmationRecord
	var r1 *ffapi.FilterResult
	var r2 error
	if rf, ok := ret.Get(0).(func(context.Context, string, ffapi.AndFilter) ([]*apitypes.ConfirmationRecord, *ffapi.FilterResult, error)); ok {
		return rf(ctx, txID, filter)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, ffapi.AndFilter) []*apitypes.ConfirmationRecord); ok {
		r0 = rf(ctx, txID, filter)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*apitypes.ConfirmationRecord)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, ffapi.AndFilter) *ffapi.FilterResult); ok {
		r1 = rf(ctx, txID, filter)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ffapi.FilterResult)
		}
	}

	if rf, ok := ret.Get(2).(func(context.Context, string, ffapi.AndFilter) error); ok {
		r2 = rf(ctx, txID, filter)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// ListTransactionHistory provides a mock function with given fields: ctx, txID, filter
func (_m *RichQuery) ListTransactionHistory(ctx context.Context, txID string, filter ffapi.AndFilter) ([]*apitypes.TXHistoryRecord, *ffapi.FilterResult, error) {
	ret := _m.Called(ctx, txID, filter)

	var r0 []*apitypes.TXHistoryRecord
	var r1 *ffapi.FilterResult
	var r2 error
	if rf, ok := ret.Get(0).(func(context.Context, string, ffapi.AndFilter) ([]*apitypes.TXHistoryRecord, *ffapi.FilterResult, error)); ok {
		return rf(ctx, txID, filter)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, ffapi.AndFilter) []*apitypes.TXHistoryRecord); ok {
		r0 = rf(ctx, txID, filter)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*apitypes.TXHistoryRecord)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, ffapi.AndFilter) *ffapi.FilterResult); ok {
		r1 = rf(ctx, txID, filter)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ffapi.FilterResult)
		}
	}

	if rf, ok := ret.Get(2).(func(context.Context, string, ffapi.AndFilter) error); ok {
		r2 = rf(ctx, txID, filter)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// ListTransactions provides a mock function with given fields: ctx, filter
func (_m *RichQuery) ListTransactions(ctx context.Context, filter ffapi.AndFilter) ([]*apitypes.ManagedTX, *ffapi.FilterResult, error) {
	ret := _m.Called(ctx, filter)

	var r0 []*apitypes.ManagedTX
	var r1 *ffapi.FilterResult
	var r2 error
	if rf, ok := ret.Get(0).(func(context.Context, ffapi.AndFilter) ([]*apitypes.ManagedTX, *ffapi.FilterResult, error)); ok {
		return rf(ctx, filter)
	}
	if rf, ok := ret.Get(0).(func(context.Context, ffapi.AndFilter) []*apitypes.ManagedTX); ok {
		r0 = rf(ctx, filter)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*apitypes.ManagedTX)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, ffapi.AndFilter) *ffapi.FilterResult); ok {
		r1 = rf(ctx, filter)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ffapi.FilterResult)
		}
	}

	if rf, ok := ret.Get(2).(func(context.Context, ffapi.AndFilter) error); ok {
		r2 = rf(ctx, filter)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// NewConfirmationFilter provides a mock function with given fields: ctx
func (_m *RichQuery) NewConfirmationFilter(ctx context.Context) ffapi.FilterBuilder {
	ret := _m.Called(ctx)

	var r0 ffapi.FilterBuilder
	if rf, ok := ret.Get(0).(func(context.Context) ffapi.FilterBuilder); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(ffapi.FilterBuilder)
		}
	}

	return r0
}

// NewListenerFilter provides a mock function with given fields: ctx
func (_m *RichQuery) NewListenerFilter(ctx context.Context) ffapi.FilterBuilder {
	ret := _m.Called(ctx)

	var r0 ffapi.FilterBuilder
	if rf, ok := ret.Get(0).(func(context.Context) ffapi.FilterBuilder); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(ffapi.FilterBuilder)
		}
	}

	return r0
}

// NewStreamFilter provides a mock function with given fields: ctx
func (_m *RichQuery) NewStreamFilter(ctx context.Context) ffapi.FilterBuilder {
	ret := _m.Called(ctx)

	var r0 ffapi.FilterBuilder
	if rf, ok := ret.Get(0).(func(context.Context) ffapi.FilterBuilder); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(ffapi.FilterBuilder)
		}
	}

	return r0
}

// NewTransactionFilter provides a mock function with given fields: ctx
func (_m *RichQuery) NewTransactionFilter(ctx context.Context) ffapi.FilterBuilder {
	ret := _m.Called(ctx)

	var r0 ffapi.FilterBuilder
	if rf, ok := ret.Get(0).(func(context.Context) ffapi.FilterBuilder); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(ffapi.FilterBuilder)
		}
	}

	return r0
}

// NewTxHistoryFilter provides a mock function with given fields: ctx
func (_m *RichQuery) NewTxHistoryFilter(ctx context.Context) ffapi.FilterBuilder {
	ret := _m.Called(ctx)

	var r0 ffapi.FilterBuilder
	if rf, ok := ret.Get(0).(func(context.Context) ffapi.FilterBuilder); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(ffapi.FilterBuilder)
		}
	}

	return r0
}

// NewRichQuery creates a new instance of RichQuery. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewRichQuery(t interface {
	mock.TestingT
	Cleanup(func())
}) *RichQuery {
	mock := &RichQuery{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
