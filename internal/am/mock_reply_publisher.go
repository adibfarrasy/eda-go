// Code generated by mockery v2.14.0. DO NOT EDIT.

package am

import (
	context "context"
	ddd "eda-in-golang/internal/ddd"

	mock "github.com/stretchr/testify/mock"
)

// MockReplyPublisher is an autogenerated mock type for the ReplyPublisher type
type MockReplyPublisher struct {
	mock.Mock
}

// Publish provides a mock function with given fields: ctx, topicName, reply
func (_m *MockReplyPublisher) Publish(ctx context.Context, topicName string, reply ddd.Reply) error {
	ret := _m.Called(ctx, topicName, reply)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, ddd.Reply) error); ok {
		r0 = rf(ctx, topicName, reply)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewMockReplyPublisher interface {
	mock.TestingT
	Cleanup(func())
}

// NewMockReplyPublisher creates a new instance of MockReplyPublisher. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockReplyPublisher(t mockConstructorTestingTNewMockReplyPublisher) *MockReplyPublisher {
	mock := &MockReplyPublisher{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
