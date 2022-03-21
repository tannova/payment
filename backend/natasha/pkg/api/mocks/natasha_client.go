// Code generated by mockery v2.4.0. DO NOT EDIT.

package mocks

import (
	context "context"

	grpc "google.golang.org/grpc"

	mock "github.com/stretchr/testify/mock"

	natasha "gitlab.com/mcuc/monorepo/backend/natasha/pkg/api"
)

// NatashaClient is an autogenerated mock type for the NatashaClient type
type NatashaClient struct {
	mock.Mock
}

// CreateMerchant provides a mock function with given fields: ctx, in, opts
func (_m *NatashaClient) CreateMerchant(ctx context.Context, in *natasha.CreateMerchantRequest, opts ...grpc.CallOption) (*natasha.CreateMerchantReply, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *natasha.CreateMerchantReply
	if rf, ok := ret.Get(0).(func(context.Context, *natasha.CreateMerchantRequest, ...grpc.CallOption) *natasha.CreateMerchantReply); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*natasha.CreateMerchantReply)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *natasha.CreateMerchantRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetMerchant provides a mock function with given fields: ctx, in, opts
func (_m *NatashaClient) GetMerchant(ctx context.Context, in *natasha.GetMerchantRequest, opts ...grpc.CallOption) (*natasha.GetMerchantReply, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *natasha.GetMerchantReply
	if rf, ok := ret.Get(0).(func(context.Context, *natasha.GetMerchantRequest, ...grpc.CallOption) *natasha.GetMerchantReply); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*natasha.GetMerchantReply)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *natasha.GetMerchantRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListMerchants provides a mock function with given fields: ctx, in, opts
func (_m *NatashaClient) ListMerchants(ctx context.Context, in *natasha.ListMerchantsRequest, opts ...grpc.CallOption) (*natasha.ListMerchantsReply, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *natasha.ListMerchantsReply
	if rf, ok := ret.Get(0).(func(context.Context, *natasha.ListMerchantsRequest, ...grpc.CallOption) *natasha.ListMerchantsReply); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*natasha.ListMerchantsReply)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *natasha.ListMerchantsRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateMerchant provides a mock function with given fields: ctx, in, opts
func (_m *NatashaClient) UpdateMerchant(ctx context.Context, in *natasha.UpdateMerchantRequest, opts ...grpc.CallOption) (*natasha.UpdateMerchantReply, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *natasha.UpdateMerchantReply
	if rf, ok := ret.Get(0).(func(context.Context, *natasha.UpdateMerchantRequest, ...grpc.CallOption) *natasha.UpdateMerchantReply); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*natasha.UpdateMerchantReply)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *natasha.UpdateMerchantRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}