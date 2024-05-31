// Code generated by mockery v2.40.2. DO NOT EDIT.

package metricsmocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// EventMetricsEmitter is an autogenerated mock type for the EventMetricsEmitter type
type EventMetricsEmitter struct {
	mock.Mock
}

// RecordBlockHashProcessMetrics provides a mock function with given fields: ctx, durationInSeconds
func (_m *EventMetricsEmitter) RecordBlockHashProcessMetrics(ctx context.Context, durationInSeconds float64) {
	_m.Called(ctx, durationInSeconds)
}

// RecordBlockHashQueueingMetrics provides a mock function with given fields: ctx, durationInSeconds
func (_m *EventMetricsEmitter) RecordBlockHashQueueingMetrics(ctx context.Context, durationInSeconds float64) {
	_m.Called(ctx, durationInSeconds)
}

// RecordConfirmationMetrics provides a mock function with given fields: ctx, durationInSeconds
func (_m *EventMetricsEmitter) RecordConfirmationMetrics(ctx context.Context, durationInSeconds float64) {
	_m.Called(ctx, durationInSeconds)
}

// RecordNotificationProcessMetrics provides a mock function with given fields: ctx, notificationType, durationInSeconds
func (_m *EventMetricsEmitter) RecordNotificationProcessMetrics(ctx context.Context, notificationType string, durationInSeconds float64) {
	_m.Called(ctx, notificationType, durationInSeconds)
}

// RecordNotificationQueueingMetrics provides a mock function with given fields: ctx, notificationType, durationInSeconds
func (_m *EventMetricsEmitter) RecordNotificationQueueingMetrics(ctx context.Context, notificationType string, durationInSeconds float64) {
	_m.Called(ctx, notificationType, durationInSeconds)
}

// RecordReceiptCheckMetrics provides a mock function with given fields: ctx, status, durationInSeconds
func (_m *EventMetricsEmitter) RecordReceiptCheckMetrics(ctx context.Context, status string, durationInSeconds float64) {
	_m.Called(ctx, status, durationInSeconds)
}

// RecordReceiptMetrics provides a mock function with given fields: ctx, durationInSeconds
func (_m *EventMetricsEmitter) RecordReceiptMetrics(ctx context.Context, durationInSeconds float64) {
	_m.Called(ctx, durationInSeconds)
}

// NewEventMetricsEmitter creates a new instance of EventMetricsEmitter. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewEventMetricsEmitter(t interface {
	mock.TestingT
	Cleanup(func())
}) *EventMetricsEmitter {
	mock := &EventMetricsEmitter{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
