// Code generated by mockery v1.0.0. DO NOT EDIT.

package ffcapimocks

import (
	context "context"

	ffcapi "github.com/hyperledger/firefly-transaction-manager/pkg/ffcapi"
	mock "github.com/stretchr/testify/mock"
)

// API is an autogenerated mock type for the API type
type API struct {
	mock.Mock
}

// BlockInfoByHash provides a mock function with given fields: ctx, req
func (_m *API) BlockInfoByHash(ctx context.Context, req *ffcapi.BlockInfoByHashRequest) (*ffcapi.BlockInfoByHashResponse, ffcapi.ErrorReason, error) {
	ret := _m.Called(ctx, req)

	var r0 *ffcapi.BlockInfoByHashResponse
	if rf, ok := ret.Get(0).(func(context.Context, *ffcapi.BlockInfoByHashRequest) *ffcapi.BlockInfoByHashResponse); ok {
		r0 = rf(ctx, req)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ffcapi.BlockInfoByHashResponse)
		}
	}

	var r1 ffcapi.ErrorReason
	if rf, ok := ret.Get(1).(func(context.Context, *ffcapi.BlockInfoByHashRequest) ffcapi.ErrorReason); ok {
		r1 = rf(ctx, req)
	} else {
		r1 = ret.Get(1).(ffcapi.ErrorReason)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context, *ffcapi.BlockInfoByHashRequest) error); ok {
		r2 = rf(ctx, req)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// BlockInfoByNumber provides a mock function with given fields: ctx, req
func (_m *API) BlockInfoByNumber(ctx context.Context, req *ffcapi.BlockInfoByNumberRequest) (*ffcapi.BlockInfoByNumberResponse, ffcapi.ErrorReason, error) {
	ret := _m.Called(ctx, req)

	var r0 *ffcapi.BlockInfoByNumberResponse
	if rf, ok := ret.Get(0).(func(context.Context, *ffcapi.BlockInfoByNumberRequest) *ffcapi.BlockInfoByNumberResponse); ok {
		r0 = rf(ctx, req)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ffcapi.BlockInfoByNumberResponse)
		}
	}

	var r1 ffcapi.ErrorReason
	if rf, ok := ret.Get(1).(func(context.Context, *ffcapi.BlockInfoByNumberRequest) ffcapi.ErrorReason); ok {
		r1 = rf(ctx, req)
	} else {
		r1 = ret.Get(1).(ffcapi.ErrorReason)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context, *ffcapi.BlockInfoByNumberRequest) error); ok {
		r2 = rf(ctx, req)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// DeployContractPrepare provides a mock function with given fields: ctx, req
func (_m *API) DeployContractPrepare(ctx context.Context, req *ffcapi.ContractDeployPrepareRequest) (*ffcapi.TransactionPrepareResponse, ffcapi.ErrorReason, error) {
	ret := _m.Called(ctx, req)

	var r0 *ffcapi.TransactionPrepareResponse
	if rf, ok := ret.Get(0).(func(context.Context, *ffcapi.ContractDeployPrepareRequest) *ffcapi.TransactionPrepareResponse); ok {
		r0 = rf(ctx, req)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ffcapi.TransactionPrepareResponse)
		}
	}

	var r1 ffcapi.ErrorReason
	if rf, ok := ret.Get(1).(func(context.Context, *ffcapi.ContractDeployPrepareRequest) ffcapi.ErrorReason); ok {
		r1 = rf(ctx, req)
	} else {
		r1 = ret.Get(1).(ffcapi.ErrorReason)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context, *ffcapi.ContractDeployPrepareRequest) error); ok {
		r2 = rf(ctx, req)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// EventListenerAdd provides a mock function with given fields: ctx, req
func (_m *API) EventListenerAdd(ctx context.Context, req *ffcapi.EventListenerAddRequest) (*ffcapi.EventListenerAddResponse, ffcapi.ErrorReason, error) {
	ret := _m.Called(ctx, req)

	var r0 *ffcapi.EventListenerAddResponse
	if rf, ok := ret.Get(0).(func(context.Context, *ffcapi.EventListenerAddRequest) *ffcapi.EventListenerAddResponse); ok {
		r0 = rf(ctx, req)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ffcapi.EventListenerAddResponse)
		}
	}

	var r1 ffcapi.ErrorReason
	if rf, ok := ret.Get(1).(func(context.Context, *ffcapi.EventListenerAddRequest) ffcapi.ErrorReason); ok {
		r1 = rf(ctx, req)
	} else {
		r1 = ret.Get(1).(ffcapi.ErrorReason)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context, *ffcapi.EventListenerAddRequest) error); ok {
		r2 = rf(ctx, req)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// EventListenerHWM provides a mock function with given fields: ctx, req
func (_m *API) EventListenerHWM(ctx context.Context, req *ffcapi.EventListenerHWMRequest) (*ffcapi.EventListenerHWMResponse, ffcapi.ErrorReason, error) {
	ret := _m.Called(ctx, req)

	var r0 *ffcapi.EventListenerHWMResponse
	if rf, ok := ret.Get(0).(func(context.Context, *ffcapi.EventListenerHWMRequest) *ffcapi.EventListenerHWMResponse); ok {
		r0 = rf(ctx, req)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ffcapi.EventListenerHWMResponse)
		}
	}

	var r1 ffcapi.ErrorReason
	if rf, ok := ret.Get(1).(func(context.Context, *ffcapi.EventListenerHWMRequest) ffcapi.ErrorReason); ok {
		r1 = rf(ctx, req)
	} else {
		r1 = ret.Get(1).(ffcapi.ErrorReason)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context, *ffcapi.EventListenerHWMRequest) error); ok {
		r2 = rf(ctx, req)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// EventListenerRemove provides a mock function with given fields: ctx, req
func (_m *API) EventListenerRemove(ctx context.Context, req *ffcapi.EventListenerRemoveRequest) (*ffcapi.EventListenerRemoveResponse, ffcapi.ErrorReason, error) {
	ret := _m.Called(ctx, req)

	var r0 *ffcapi.EventListenerRemoveResponse
	if rf, ok := ret.Get(0).(func(context.Context, *ffcapi.EventListenerRemoveRequest) *ffcapi.EventListenerRemoveResponse); ok {
		r0 = rf(ctx, req)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ffcapi.EventListenerRemoveResponse)
		}
	}

	var r1 ffcapi.ErrorReason
	if rf, ok := ret.Get(1).(func(context.Context, *ffcapi.EventListenerRemoveRequest) ffcapi.ErrorReason); ok {
		r1 = rf(ctx, req)
	} else {
		r1 = ret.Get(1).(ffcapi.ErrorReason)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context, *ffcapi.EventListenerRemoveRequest) error); ok {
		r2 = rf(ctx, req)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// EventListenerVerifyOptions provides a mock function with given fields: ctx, req
func (_m *API) EventListenerVerifyOptions(ctx context.Context, req *ffcapi.EventListenerVerifyOptionsRequest) (*ffcapi.EventListenerVerifyOptionsResponse, ffcapi.ErrorReason, error) {
	ret := _m.Called(ctx, req)

	var r0 *ffcapi.EventListenerVerifyOptionsResponse
	if rf, ok := ret.Get(0).(func(context.Context, *ffcapi.EventListenerVerifyOptionsRequest) *ffcapi.EventListenerVerifyOptionsResponse); ok {
		r0 = rf(ctx, req)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ffcapi.EventListenerVerifyOptionsResponse)
		}
	}

	var r1 ffcapi.ErrorReason
	if rf, ok := ret.Get(1).(func(context.Context, *ffcapi.EventListenerVerifyOptionsRequest) ffcapi.ErrorReason); ok {
		r1 = rf(ctx, req)
	} else {
		r1 = ret.Get(1).(ffcapi.ErrorReason)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context, *ffcapi.EventListenerVerifyOptionsRequest) error); ok {
		r2 = rf(ctx, req)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// EventStreamNewCheckpointStruct provides a mock function with given fields:
func (_m *API) EventStreamNewCheckpointStruct() ffcapi.EventListenerCheckpoint {
	ret := _m.Called()

	var r0 ffcapi.EventListenerCheckpoint
	if rf, ok := ret.Get(0).(func() ffcapi.EventListenerCheckpoint); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(ffcapi.EventListenerCheckpoint)
		}
	}

	return r0
}

// EventStreamStart provides a mock function with given fields: ctx, req
func (_m *API) EventStreamStart(ctx context.Context, req *ffcapi.EventStreamStartRequest) (*ffcapi.EventStreamStartResponse, ffcapi.ErrorReason, error) {
	ret := _m.Called(ctx, req)

	var r0 *ffcapi.EventStreamStartResponse
	if rf, ok := ret.Get(0).(func(context.Context, *ffcapi.EventStreamStartRequest) *ffcapi.EventStreamStartResponse); ok {
		r0 = rf(ctx, req)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ffcapi.EventStreamStartResponse)
		}
	}

	var r1 ffcapi.ErrorReason
	if rf, ok := ret.Get(1).(func(context.Context, *ffcapi.EventStreamStartRequest) ffcapi.ErrorReason); ok {
		r1 = rf(ctx, req)
	} else {
		r1 = ret.Get(1).(ffcapi.ErrorReason)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context, *ffcapi.EventStreamStartRequest) error); ok {
		r2 = rf(ctx, req)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// EventStreamStopped provides a mock function with given fields: ctx, req
func (_m *API) EventStreamStopped(ctx context.Context, req *ffcapi.EventStreamStoppedRequest) (*ffcapi.EventStreamStoppedResponse, ffcapi.ErrorReason, error) {
	ret := _m.Called(ctx, req)

	var r0 *ffcapi.EventStreamStoppedResponse
	if rf, ok := ret.Get(0).(func(context.Context, *ffcapi.EventStreamStoppedRequest) *ffcapi.EventStreamStoppedResponse); ok {
		r0 = rf(ctx, req)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ffcapi.EventStreamStoppedResponse)
		}
	}

	var r1 ffcapi.ErrorReason
	if rf, ok := ret.Get(1).(func(context.Context, *ffcapi.EventStreamStoppedRequest) ffcapi.ErrorReason); ok {
		r1 = rf(ctx, req)
	} else {
		r1 = ret.Get(1).(ffcapi.ErrorReason)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context, *ffcapi.EventStreamStoppedRequest) error); ok {
		r2 = rf(ctx, req)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// GasPriceEstimate provides a mock function with given fields: ctx, req
func (_m *API) GasPriceEstimate(ctx context.Context, req *ffcapi.GasPriceEstimateRequest) (*ffcapi.GasPriceEstimateResponse, ffcapi.ErrorReason, error) {
	ret := _m.Called(ctx, req)

	var r0 *ffcapi.GasPriceEstimateResponse
	if rf, ok := ret.Get(0).(func(context.Context, *ffcapi.GasPriceEstimateRequest) *ffcapi.GasPriceEstimateResponse); ok {
		r0 = rf(ctx, req)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ffcapi.GasPriceEstimateResponse)
		}
	}

	var r1 ffcapi.ErrorReason
	if rf, ok := ret.Get(1).(func(context.Context, *ffcapi.GasPriceEstimateRequest) ffcapi.ErrorReason); ok {
		r1 = rf(ctx, req)
	} else {
		r1 = ret.Get(1).(ffcapi.ErrorReason)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context, *ffcapi.GasPriceEstimateRequest) error); ok {
		r2 = rf(ctx, req)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// IsLive provides a mock function with given fields: ctx
func (_m *API) IsLive(ctx context.Context) (*ffcapi.LiveResponse, ffcapi.ErrorReason, error) {
	ret := _m.Called(ctx)

	var r0 *ffcapi.LiveResponse
	if rf, ok := ret.Get(0).(func(context.Context) *ffcapi.LiveResponse); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ffcapi.LiveResponse)
		}
	}

	var r1 ffcapi.ErrorReason
	if rf, ok := ret.Get(1).(func(context.Context) ffcapi.ErrorReason); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Get(1).(ffcapi.ErrorReason)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context) error); ok {
		r2 = rf(ctx)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// IsReady provides a mock function with given fields: ctx
func (_m *API) IsReady(ctx context.Context) (*ffcapi.ReadyResponse, ffcapi.ErrorReason, error) {
	ret := _m.Called(ctx)

	var r0 *ffcapi.ReadyResponse
	if rf, ok := ret.Get(0).(func(context.Context) *ffcapi.ReadyResponse); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ffcapi.ReadyResponse)
		}
	}

	var r1 ffcapi.ErrorReason
	if rf, ok := ret.Get(1).(func(context.Context) ffcapi.ErrorReason); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Get(1).(ffcapi.ErrorReason)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context) error); ok {
		r2 = rf(ctx)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// NewBlockListener provides a mock function with given fields: ctx, req
func (_m *API) NewBlockListener(ctx context.Context, req *ffcapi.NewBlockListenerRequest) (*ffcapi.NewBlockListenerResponse, ffcapi.ErrorReason, error) {
	ret := _m.Called(ctx, req)

	var r0 *ffcapi.NewBlockListenerResponse
	if rf, ok := ret.Get(0).(func(context.Context, *ffcapi.NewBlockListenerRequest) *ffcapi.NewBlockListenerResponse); ok {
		r0 = rf(ctx, req)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ffcapi.NewBlockListenerResponse)
		}
	}

	var r1 ffcapi.ErrorReason
	if rf, ok := ret.Get(1).(func(context.Context, *ffcapi.NewBlockListenerRequest) ffcapi.ErrorReason); ok {
		r1 = rf(ctx, req)
	} else {
		r1 = ret.Get(1).(ffcapi.ErrorReason)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context, *ffcapi.NewBlockListenerRequest) error); ok {
		r2 = rf(ctx, req)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// NextNonceForSigner provides a mock function with given fields: ctx, req
func (_m *API) NextNonceForSigner(ctx context.Context, req *ffcapi.NextNonceForSignerRequest) (*ffcapi.NextNonceForSignerResponse, ffcapi.ErrorReason, error) {
	ret := _m.Called(ctx, req)

	var r0 *ffcapi.NextNonceForSignerResponse
	if rf, ok := ret.Get(0).(func(context.Context, *ffcapi.NextNonceForSignerRequest) *ffcapi.NextNonceForSignerResponse); ok {
		r0 = rf(ctx, req)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ffcapi.NextNonceForSignerResponse)
		}
	}

	var r1 ffcapi.ErrorReason
	if rf, ok := ret.Get(1).(func(context.Context, *ffcapi.NextNonceForSignerRequest) ffcapi.ErrorReason); ok {
		r1 = rf(ctx, req)
	} else {
		r1 = ret.Get(1).(ffcapi.ErrorReason)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context, *ffcapi.NextNonceForSignerRequest) error); ok {
		r2 = rf(ctx, req)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// QueryInvoke provides a mock function with given fields: ctx, req
func (_m *API) QueryInvoke(ctx context.Context, req *ffcapi.QueryInvokeRequest) (*ffcapi.QueryInvokeResponse, ffcapi.ErrorReason, error) {
	ret := _m.Called(ctx, req)

	var r0 *ffcapi.QueryInvokeResponse
	if rf, ok := ret.Get(0).(func(context.Context, *ffcapi.QueryInvokeRequest) *ffcapi.QueryInvokeResponse); ok {
		r0 = rf(ctx, req)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ffcapi.QueryInvokeResponse)
		}
	}

	var r1 ffcapi.ErrorReason
	if rf, ok := ret.Get(1).(func(context.Context, *ffcapi.QueryInvokeRequest) ffcapi.ErrorReason); ok {
		r1 = rf(ctx, req)
	} else {
		r1 = ret.Get(1).(ffcapi.ErrorReason)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context, *ffcapi.QueryInvokeRequest) error); ok {
		r2 = rf(ctx, req)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// TransactionPrepare provides a mock function with given fields: ctx, req
func (_m *API) TransactionPrepare(ctx context.Context, req *ffcapi.TransactionPrepareRequest) (*ffcapi.TransactionPrepareResponse, ffcapi.ErrorReason, error) {
	ret := _m.Called(ctx, req)

	var r0 *ffcapi.TransactionPrepareResponse
	if rf, ok := ret.Get(0).(func(context.Context, *ffcapi.TransactionPrepareRequest) *ffcapi.TransactionPrepareResponse); ok {
		r0 = rf(ctx, req)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ffcapi.TransactionPrepareResponse)
		}
	}

	var r1 ffcapi.ErrorReason
	if rf, ok := ret.Get(1).(func(context.Context, *ffcapi.TransactionPrepareRequest) ffcapi.ErrorReason); ok {
		r1 = rf(ctx, req)
	} else {
		r1 = ret.Get(1).(ffcapi.ErrorReason)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context, *ffcapi.TransactionPrepareRequest) error); ok {
		r2 = rf(ctx, req)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// TransactionReceipt provides a mock function with given fields: ctx, req
func (_m *API) TransactionReceipt(ctx context.Context, req *ffcapi.TransactionReceiptRequest) (*ffcapi.TransactionReceiptResponse, ffcapi.ErrorReason, error) {
	ret := _m.Called(ctx, req)

	var r0 *ffcapi.TransactionReceiptResponse
	if rf, ok := ret.Get(0).(func(context.Context, *ffcapi.TransactionReceiptRequest) *ffcapi.TransactionReceiptResponse); ok {
		r0 = rf(ctx, req)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ffcapi.TransactionReceiptResponse)
		}
	}

	var r1 ffcapi.ErrorReason
	if rf, ok := ret.Get(1).(func(context.Context, *ffcapi.TransactionReceiptRequest) ffcapi.ErrorReason); ok {
		r1 = rf(ctx, req)
	} else {
		r1 = ret.Get(1).(ffcapi.ErrorReason)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context, *ffcapi.TransactionReceiptRequest) error); ok {
		r2 = rf(ctx, req)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// TransactionSend provides a mock function with given fields: ctx, req
func (_m *API) TransactionSend(ctx context.Context, req *ffcapi.TransactionSendRequest) (*ffcapi.TransactionSendResponse, ffcapi.ErrorReason, error) {
	ret := _m.Called(ctx, req)

	var r0 *ffcapi.TransactionSendResponse
	if rf, ok := ret.Get(0).(func(context.Context, *ffcapi.TransactionSendRequest) *ffcapi.TransactionSendResponse); ok {
		r0 = rf(ctx, req)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ffcapi.TransactionSendResponse)
		}
	}

	var r1 ffcapi.ErrorReason
	if rf, ok := ret.Get(1).(func(context.Context, *ffcapi.TransactionSendRequest) ffcapi.ErrorReason); ok {
		r1 = rf(ctx, req)
	} else {
		r1 = ret.Get(1).(ffcapi.ErrorReason)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context, *ffcapi.TransactionSendRequest) error); ok {
		r2 = rf(ctx, req)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}
