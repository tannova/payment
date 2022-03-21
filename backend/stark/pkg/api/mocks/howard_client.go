// Code generated by mockery v2.3.0. DO NOT EDIT.

package mocks

import (
	context "context"

	grpc "google.golang.org/grpc"

	mock "github.com/stretchr/testify/mock"

	stark "gitlab.com/mcuc/monorepo/backend/stark/pkg/api"
)

// HowardClient is an autogenerated mock type for the HowardClient type
type HowardClient struct {
	mock.Mock
}

// GetAllocationTopUpRate provides a mock function with given fields: ctx, in, opts
func (_m *HowardClient) GetAllocationTopUpRate(ctx context.Context, in *stark.GetReportRequest, opts ...grpc.CallOption) (*stark.GetAllocationTopUpRateReply, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *stark.GetAllocationTopUpRateReply
	if rf, ok := ret.Get(0).(func(context.Context, *stark.GetReportRequest, ...grpc.CallOption) *stark.GetAllocationTopUpRateReply); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*stark.GetAllocationTopUpRateReply)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *stark.GetReportRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAllocationWithdrawRate provides a mock function with given fields: ctx, in, opts
func (_m *HowardClient) GetAllocationWithdrawRate(ctx context.Context, in *stark.GetReportRequest, opts ...grpc.CallOption) (*stark.GetAllocationWithdrawRateReply, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *stark.GetAllocationWithdrawRateReply
	if rf, ok := ret.Get(0).(func(context.Context, *stark.GetReportRequest, ...grpc.CallOption) *stark.GetAllocationWithdrawRateReply); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*stark.GetAllocationWithdrawRateReply)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *stark.GetReportRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetIncomeStatement provides a mock function with given fields: ctx, in, opts
func (_m *HowardClient) GetIncomeStatement(ctx context.Context, in *stark.GetReportRequest, opts ...grpc.CallOption) (*stark.GetIncomeStatementReply, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *stark.GetIncomeStatementReply
	if rf, ok := ret.Get(0).(func(context.Context, *stark.GetReportRequest, ...grpc.CallOption) *stark.GetIncomeStatementReply); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*stark.GetIncomeStatementReply)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *stark.GetReportRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetPaymentToday provides a mock function with given fields: ctx, in, opts
func (_m *HowardClient) GetPaymentToday(ctx context.Context, in *stark.GetReportRequest, opts ...grpc.CallOption) (*stark.GetPaymentTodayReply, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *stark.GetPaymentTodayReply
	if rf, ok := ret.Get(0).(func(context.Context, *stark.GetReportRequest, ...grpc.CallOption) *stark.GetPaymentTodayReply); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*stark.GetPaymentTodayReply)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *stark.GetReportRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetProcessingPerformance provides a mock function with given fields: ctx, in, opts
func (_m *HowardClient) GetProcessingPerformance(ctx context.Context, in *stark.GetProcessingPerformanceRequest, opts ...grpc.CallOption) (*stark.GetProcessingPerformanceReply, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *stark.GetProcessingPerformanceReply
	if rf, ok := ret.Get(0).(func(context.Context, *stark.GetProcessingPerformanceRequest, ...grpc.CallOption) *stark.GetProcessingPerformanceReply); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*stark.GetProcessingPerformanceReply)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *stark.GetProcessingPerformanceRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetProfitRate provides a mock function with given fields: ctx, in, opts
func (_m *HowardClient) GetProfitRate(ctx context.Context, in *stark.GetReportRequest, opts ...grpc.CallOption) (*stark.GetProfitRateReply, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *stark.GetProfitRateReply
	if rf, ok := ret.Get(0).(func(context.Context, *stark.GetReportRequest, ...grpc.CallOption) *stark.GetProfitRateReply); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*stark.GetProfitRateReply)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *stark.GetReportRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetSellReportByMerchant provides a mock function with given fields: ctx, in, opts
func (_m *HowardClient) GetSellReportByMerchant(ctx context.Context, in *stark.GetSellReportByMerchantRequest, opts ...grpc.CallOption) (*stark.GetSellReportByMerchantReply, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *stark.GetSellReportByMerchantReply
	if rf, ok := ret.Get(0).(func(context.Context, *stark.GetSellReportByMerchantRequest, ...grpc.CallOption) *stark.GetSellReportByMerchantReply); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*stark.GetSellReportByMerchantReply)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *stark.GetSellReportByMerchantRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetSellReportByPaymentMethod provides a mock function with given fields: ctx, in, opts
func (_m *HowardClient) GetSellReportByPaymentMethod(ctx context.Context, in *stark.GetSellReportByPaymentMethodRequest, opts ...grpc.CallOption) (*stark.GetSellReportByPaymentMethodReply, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *stark.GetSellReportByPaymentMethodReply
	if rf, ok := ret.Get(0).(func(context.Context, *stark.GetSellReportByPaymentMethodRequest, ...grpc.CallOption) *stark.GetSellReportByPaymentMethodReply); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*stark.GetSellReportByPaymentMethodReply)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *stark.GetSellReportByPaymentMethodRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetSellReportByTeller provides a mock function with given fields: ctx, in, opts
func (_m *HowardClient) GetSellReportByTeller(ctx context.Context, in *stark.GetSellReportByTellerRequest, opts ...grpc.CallOption) (*stark.GetSellReportByTellerReply, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *stark.GetSellReportByTellerReply
	if rf, ok := ret.Get(0).(func(context.Context, *stark.GetSellReportByTellerRequest, ...grpc.CallOption) *stark.GetSellReportByTellerReply); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*stark.GetSellReportByTellerReply)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *stark.GetSellReportByTellerRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetSellReportByTimeRange provides a mock function with given fields: ctx, in, opts
func (_m *HowardClient) GetSellReportByTimeRange(ctx context.Context, in *stark.GetSellReportByTimeRangeRequest, opts ...grpc.CallOption) (*stark.GetSellReportByTimeRangeReply, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *stark.GetSellReportByTimeRangeReply
	if rf, ok := ret.Get(0).(func(context.Context, *stark.GetSellReportByTimeRangeRequest, ...grpc.CallOption) *stark.GetSellReportByTimeRangeReply); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*stark.GetSellReportByTimeRangeReply)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *stark.GetSellReportByTimeRangeRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetStatistic provides a mock function with given fields: ctx, in, opts
func (_m *HowardClient) GetStatistic(ctx context.Context, in *stark.GetStatisticRequest, opts ...grpc.CallOption) (*stark.GetStatisticReply, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *stark.GetStatisticReply
	if rf, ok := ret.Get(0).(func(context.Context, *stark.GetStatisticRequest, ...grpc.CallOption) *stark.GetStatisticReply); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*stark.GetStatisticReply)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *stark.GetStatisticRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTopPaymentMethod provides a mock function with given fields: ctx, in, opts
func (_m *HowardClient) GetTopPaymentMethod(ctx context.Context, in *stark.GetReportRequest, opts ...grpc.CallOption) (*stark.GetTopPaymentMethodReply, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *stark.GetTopPaymentMethodReply
	if rf, ok := ret.Get(0).(func(context.Context, *stark.GetReportRequest, ...grpc.CallOption) *stark.GetTopPaymentMethodReply); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*stark.GetTopPaymentMethodReply)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *stark.GetReportRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTopTeller provides a mock function with given fields: ctx, in, opts
func (_m *HowardClient) GetTopTeller(ctx context.Context, in *stark.GetReportRequest, opts ...grpc.CallOption) (*stark.GetTopTellerReply, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *stark.GetTopTellerReply
	if rf, ok := ret.Get(0).(func(context.Context, *stark.GetReportRequest, ...grpc.CallOption) *stark.GetTopTellerReply); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*stark.GetTopTellerReply)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *stark.GetReportRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTotalAmount provides a mock function with given fields: ctx, in, opts
func (_m *HowardClient) GetTotalAmount(ctx context.Context, in *stark.GetTotalAmountRequest, opts ...grpc.CallOption) (*stark.GetTotalAmountReply, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *stark.GetTotalAmountReply
	if rf, ok := ret.Get(0).(func(context.Context, *stark.GetTotalAmountRequest, ...grpc.CallOption) *stark.GetTotalAmountReply); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*stark.GetTotalAmountReply)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *stark.GetTotalAmountRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}