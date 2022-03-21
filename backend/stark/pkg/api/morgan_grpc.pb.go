// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package stark

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// MorganClient is the client API for Morgan service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MorganClient interface {
	// MEX call to request a payment for their user to make payment to telco
	CreateTelcoWithdraw(ctx context.Context, in *CreateTelcoWithdrawRequest, opts ...grpc.CallOption) (*CreateTelcoWithdrawReply, error)
	CreateTelcoTopUp(ctx context.Context, in *CreateTelcoTopUpRequest, opts ...grpc.CallOption) (*CreateTelcoTopUpReply, error)
	CancelTelcoWithdraw(ctx context.Context, in *CancelTelcoWithdrawRequest, opts ...grpc.CallOption) (*CancelTelcoWithdrawReply, error)
	// When telco provider callback to Bishop, it will call this api to complete a top up telco payment
	CompleteTelcoTopUp(ctx context.Context, in *CompleteTelcoTopUpRequest, opts ...grpc.CallOption) (*CompleteTelcoTopUpReply, error)
	// Teller manage top-up and withdraw payment
	ApproveTelcoTopUp(ctx context.Context, in *ApproveTelcoTopUpRequest, opts ...grpc.CallOption) (*ApproveTelcoTopUpReply, error)
	ApproveTelcoWithdraw(ctx context.Context, in *ApproveTelcoWithdrawRequest, opts ...grpc.CallOption) (*ApproveTelcoWithdrawReply, error)
	RejectTelcoTopUp(ctx context.Context, in *RejectTelcoTopUpRequest, opts ...grpc.CallOption) (*RejectTelcoTopUpReply, error)
	RejectTelcoWithdraw(ctx context.Context, in *RejectTelcoWithdrawRequest, opts ...grpc.CallOption) (*RejectTelcoWithdrawReply, error)
	GetSettings(ctx context.Context, in *GetSettingsRequest, opts ...grpc.CallOption) (*GetSettingsReply, error)
	UpdateTopUpAutoApprovalSetting(ctx context.Context, in *UpdateTopUpAutoApprovalSettingRequest, opts ...grpc.CallOption) (*UpdateTopUpAutoApprovalSettingReply, error)
	UpdateUsingThirdPartySetting(ctx context.Context, in *UpdateUsingThirdPartySettingRequest, opts ...grpc.CallOption) (*UpdateUsingThirdPartySettingReply, error)
	UpdateChargeCardProvidersSetting(ctx context.Context, in *UpdateChargeCardProvidersSettingRequest, opts ...grpc.CallOption) (*UpdateChargeCardProvidersSettingReply, error)
	UpdateGetCardProvidersSetting(ctx context.Context, in *UpdateGetCardProvidersSettingRequest, opts ...grpc.CallOption) (*UpdateGetCardProvidersSettingReply, error)
}

type morganClient struct {
	cc grpc.ClientConnInterface
}

func NewMorganClient(cc grpc.ClientConnInterface) MorganClient {
	return &morganClient{cc}
}

func (c *morganClient) CreateTelcoWithdraw(ctx context.Context, in *CreateTelcoWithdrawRequest, opts ...grpc.CallOption) (*CreateTelcoWithdrawReply, error) {
	out := new(CreateTelcoWithdrawReply)
	err := c.cc.Invoke(ctx, "/mcuc.stark.morgan.Morgan/CreateTelcoWithdraw", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *morganClient) CreateTelcoTopUp(ctx context.Context, in *CreateTelcoTopUpRequest, opts ...grpc.CallOption) (*CreateTelcoTopUpReply, error) {
	out := new(CreateTelcoTopUpReply)
	err := c.cc.Invoke(ctx, "/mcuc.stark.morgan.Morgan/CreateTelcoTopUp", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *morganClient) CancelTelcoWithdraw(ctx context.Context, in *CancelTelcoWithdrawRequest, opts ...grpc.CallOption) (*CancelTelcoWithdrawReply, error) {
	out := new(CancelTelcoWithdrawReply)
	err := c.cc.Invoke(ctx, "/mcuc.stark.morgan.Morgan/CancelTelcoWithdraw", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *morganClient) CompleteTelcoTopUp(ctx context.Context, in *CompleteTelcoTopUpRequest, opts ...grpc.CallOption) (*CompleteTelcoTopUpReply, error) {
	out := new(CompleteTelcoTopUpReply)
	err := c.cc.Invoke(ctx, "/mcuc.stark.morgan.Morgan/CompleteTelcoTopUp", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *morganClient) ApproveTelcoTopUp(ctx context.Context, in *ApproveTelcoTopUpRequest, opts ...grpc.CallOption) (*ApproveTelcoTopUpReply, error) {
	out := new(ApproveTelcoTopUpReply)
	err := c.cc.Invoke(ctx, "/mcuc.stark.morgan.Morgan/ApproveTelcoTopUp", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *morganClient) ApproveTelcoWithdraw(ctx context.Context, in *ApproveTelcoWithdrawRequest, opts ...grpc.CallOption) (*ApproveTelcoWithdrawReply, error) {
	out := new(ApproveTelcoWithdrawReply)
	err := c.cc.Invoke(ctx, "/mcuc.stark.morgan.Morgan/ApproveTelcoWithdraw", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *morganClient) RejectTelcoTopUp(ctx context.Context, in *RejectTelcoTopUpRequest, opts ...grpc.CallOption) (*RejectTelcoTopUpReply, error) {
	out := new(RejectTelcoTopUpReply)
	err := c.cc.Invoke(ctx, "/mcuc.stark.morgan.Morgan/RejectTelcoTopUp", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *morganClient) RejectTelcoWithdraw(ctx context.Context, in *RejectTelcoWithdrawRequest, opts ...grpc.CallOption) (*RejectTelcoWithdrawReply, error) {
	out := new(RejectTelcoWithdrawReply)
	err := c.cc.Invoke(ctx, "/mcuc.stark.morgan.Morgan/RejectTelcoWithdraw", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *morganClient) GetSettings(ctx context.Context, in *GetSettingsRequest, opts ...grpc.CallOption) (*GetSettingsReply, error) {
	out := new(GetSettingsReply)
	err := c.cc.Invoke(ctx, "/mcuc.stark.morgan.Morgan/GetSettings", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *morganClient) UpdateTopUpAutoApprovalSetting(ctx context.Context, in *UpdateTopUpAutoApprovalSettingRequest, opts ...grpc.CallOption) (*UpdateTopUpAutoApprovalSettingReply, error) {
	out := new(UpdateTopUpAutoApprovalSettingReply)
	err := c.cc.Invoke(ctx, "/mcuc.stark.morgan.Morgan/UpdateTopUpAutoApprovalSetting", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *morganClient) UpdateUsingThirdPartySetting(ctx context.Context, in *UpdateUsingThirdPartySettingRequest, opts ...grpc.CallOption) (*UpdateUsingThirdPartySettingReply, error) {
	out := new(UpdateUsingThirdPartySettingReply)
	err := c.cc.Invoke(ctx, "/mcuc.stark.morgan.Morgan/UpdateUsingThirdPartySetting", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *morganClient) UpdateChargeCardProvidersSetting(ctx context.Context, in *UpdateChargeCardProvidersSettingRequest, opts ...grpc.CallOption) (*UpdateChargeCardProvidersSettingReply, error) {
	out := new(UpdateChargeCardProvidersSettingReply)
	err := c.cc.Invoke(ctx, "/mcuc.stark.morgan.Morgan/UpdateChargeCardProvidersSetting", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *morganClient) UpdateGetCardProvidersSetting(ctx context.Context, in *UpdateGetCardProvidersSettingRequest, opts ...grpc.CallOption) (*UpdateGetCardProvidersSettingReply, error) {
	out := new(UpdateGetCardProvidersSettingReply)
	err := c.cc.Invoke(ctx, "/mcuc.stark.morgan.Morgan/UpdateGetCardProvidersSetting", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MorganServer is the server API for Morgan service.
// All implementations must embed UnimplementedMorganServer
// for forward compatibility
type MorganServer interface {
	// MEX call to request a payment for their user to make payment to telco
	CreateTelcoWithdraw(context.Context, *CreateTelcoWithdrawRequest) (*CreateTelcoWithdrawReply, error)
	CreateTelcoTopUp(context.Context, *CreateTelcoTopUpRequest) (*CreateTelcoTopUpReply, error)
	CancelTelcoWithdraw(context.Context, *CancelTelcoWithdrawRequest) (*CancelTelcoWithdrawReply, error)
	// When telco provider callback to Bishop, it will call this api to complete a top up telco payment
	CompleteTelcoTopUp(context.Context, *CompleteTelcoTopUpRequest) (*CompleteTelcoTopUpReply, error)
	// Teller manage top-up and withdraw payment
	ApproveTelcoTopUp(context.Context, *ApproveTelcoTopUpRequest) (*ApproveTelcoTopUpReply, error)
	ApproveTelcoWithdraw(context.Context, *ApproveTelcoWithdrawRequest) (*ApproveTelcoWithdrawReply, error)
	RejectTelcoTopUp(context.Context, *RejectTelcoTopUpRequest) (*RejectTelcoTopUpReply, error)
	RejectTelcoWithdraw(context.Context, *RejectTelcoWithdrawRequest) (*RejectTelcoWithdrawReply, error)
	GetSettings(context.Context, *GetSettingsRequest) (*GetSettingsReply, error)
	UpdateTopUpAutoApprovalSetting(context.Context, *UpdateTopUpAutoApprovalSettingRequest) (*UpdateTopUpAutoApprovalSettingReply, error)
	UpdateUsingThirdPartySetting(context.Context, *UpdateUsingThirdPartySettingRequest) (*UpdateUsingThirdPartySettingReply, error)
	UpdateChargeCardProvidersSetting(context.Context, *UpdateChargeCardProvidersSettingRequest) (*UpdateChargeCardProvidersSettingReply, error)
	UpdateGetCardProvidersSetting(context.Context, *UpdateGetCardProvidersSettingRequest) (*UpdateGetCardProvidersSettingReply, error)
	mustEmbedUnimplementedMorganServer()
}

// UnimplementedMorganServer must be embedded to have forward compatible implementations.
type UnimplementedMorganServer struct {
}

func (*UnimplementedMorganServer) CreateTelcoWithdraw(context.Context, *CreateTelcoWithdrawRequest) (*CreateTelcoWithdrawReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTelcoWithdraw not implemented")
}
func (*UnimplementedMorganServer) CreateTelcoTopUp(context.Context, *CreateTelcoTopUpRequest) (*CreateTelcoTopUpReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTelcoTopUp not implemented")
}
func (*UnimplementedMorganServer) CancelTelcoWithdraw(context.Context, *CancelTelcoWithdrawRequest) (*CancelTelcoWithdrawReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CancelTelcoWithdraw not implemented")
}
func (*UnimplementedMorganServer) CompleteTelcoTopUp(context.Context, *CompleteTelcoTopUpRequest) (*CompleteTelcoTopUpReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CompleteTelcoTopUp not implemented")
}
func (*UnimplementedMorganServer) ApproveTelcoTopUp(context.Context, *ApproveTelcoTopUpRequest) (*ApproveTelcoTopUpReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ApproveTelcoTopUp not implemented")
}
func (*UnimplementedMorganServer) ApproveTelcoWithdraw(context.Context, *ApproveTelcoWithdrawRequest) (*ApproveTelcoWithdrawReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ApproveTelcoWithdraw not implemented")
}
func (*UnimplementedMorganServer) RejectTelcoTopUp(context.Context, *RejectTelcoTopUpRequest) (*RejectTelcoTopUpReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RejectTelcoTopUp not implemented")
}
func (*UnimplementedMorganServer) RejectTelcoWithdraw(context.Context, *RejectTelcoWithdrawRequest) (*RejectTelcoWithdrawReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RejectTelcoWithdraw not implemented")
}
func (*UnimplementedMorganServer) GetSettings(context.Context, *GetSettingsRequest) (*GetSettingsReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSettings not implemented")
}
func (*UnimplementedMorganServer) UpdateTopUpAutoApprovalSetting(context.Context, *UpdateTopUpAutoApprovalSettingRequest) (*UpdateTopUpAutoApprovalSettingReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateTopUpAutoApprovalSetting not implemented")
}
func (*UnimplementedMorganServer) UpdateUsingThirdPartySetting(context.Context, *UpdateUsingThirdPartySettingRequest) (*UpdateUsingThirdPartySettingReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUsingThirdPartySetting not implemented")
}
func (*UnimplementedMorganServer) UpdateChargeCardProvidersSetting(context.Context, *UpdateChargeCardProvidersSettingRequest) (*UpdateChargeCardProvidersSettingReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateChargeCardProvidersSetting not implemented")
}
func (*UnimplementedMorganServer) UpdateGetCardProvidersSetting(context.Context, *UpdateGetCardProvidersSettingRequest) (*UpdateGetCardProvidersSettingReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateGetCardProvidersSetting not implemented")
}
func (*UnimplementedMorganServer) mustEmbedUnimplementedMorganServer() {}

func RegisterMorganServer(s *grpc.Server, srv MorganServer) {
	s.RegisterService(&_Morgan_serviceDesc, srv)
}

func _Morgan_CreateTelcoWithdraw_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTelcoWithdrawRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MorganServer).CreateTelcoWithdraw(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mcuc.stark.morgan.Morgan/CreateTelcoWithdraw",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MorganServer).CreateTelcoWithdraw(ctx, req.(*CreateTelcoWithdrawRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Morgan_CreateTelcoTopUp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTelcoTopUpRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MorganServer).CreateTelcoTopUp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mcuc.stark.morgan.Morgan/CreateTelcoTopUp",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MorganServer).CreateTelcoTopUp(ctx, req.(*CreateTelcoTopUpRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Morgan_CancelTelcoWithdraw_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CancelTelcoWithdrawRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MorganServer).CancelTelcoWithdraw(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mcuc.stark.morgan.Morgan/CancelTelcoWithdraw",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MorganServer).CancelTelcoWithdraw(ctx, req.(*CancelTelcoWithdrawRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Morgan_CompleteTelcoTopUp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CompleteTelcoTopUpRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MorganServer).CompleteTelcoTopUp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mcuc.stark.morgan.Morgan/CompleteTelcoTopUp",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MorganServer).CompleteTelcoTopUp(ctx, req.(*CompleteTelcoTopUpRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Morgan_ApproveTelcoTopUp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ApproveTelcoTopUpRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MorganServer).ApproveTelcoTopUp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mcuc.stark.morgan.Morgan/ApproveTelcoTopUp",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MorganServer).ApproveTelcoTopUp(ctx, req.(*ApproveTelcoTopUpRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Morgan_ApproveTelcoWithdraw_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ApproveTelcoWithdrawRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MorganServer).ApproveTelcoWithdraw(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mcuc.stark.morgan.Morgan/ApproveTelcoWithdraw",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MorganServer).ApproveTelcoWithdraw(ctx, req.(*ApproveTelcoWithdrawRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Morgan_RejectTelcoTopUp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RejectTelcoTopUpRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MorganServer).RejectTelcoTopUp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mcuc.stark.morgan.Morgan/RejectTelcoTopUp",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MorganServer).RejectTelcoTopUp(ctx, req.(*RejectTelcoTopUpRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Morgan_RejectTelcoWithdraw_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RejectTelcoWithdrawRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MorganServer).RejectTelcoWithdraw(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mcuc.stark.morgan.Morgan/RejectTelcoWithdraw",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MorganServer).RejectTelcoWithdraw(ctx, req.(*RejectTelcoWithdrawRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Morgan_GetSettings_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSettingsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MorganServer).GetSettings(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mcuc.stark.morgan.Morgan/GetSettings",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MorganServer).GetSettings(ctx, req.(*GetSettingsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Morgan_UpdateTopUpAutoApprovalSetting_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateTopUpAutoApprovalSettingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MorganServer).UpdateTopUpAutoApprovalSetting(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mcuc.stark.morgan.Morgan/UpdateTopUpAutoApprovalSetting",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MorganServer).UpdateTopUpAutoApprovalSetting(ctx, req.(*UpdateTopUpAutoApprovalSettingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Morgan_UpdateUsingThirdPartySetting_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateUsingThirdPartySettingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MorganServer).UpdateUsingThirdPartySetting(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mcuc.stark.morgan.Morgan/UpdateUsingThirdPartySetting",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MorganServer).UpdateUsingThirdPartySetting(ctx, req.(*UpdateUsingThirdPartySettingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Morgan_UpdateChargeCardProvidersSetting_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateChargeCardProvidersSettingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MorganServer).UpdateChargeCardProvidersSetting(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mcuc.stark.morgan.Morgan/UpdateChargeCardProvidersSetting",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MorganServer).UpdateChargeCardProvidersSetting(ctx, req.(*UpdateChargeCardProvidersSettingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Morgan_UpdateGetCardProvidersSetting_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateGetCardProvidersSettingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MorganServer).UpdateGetCardProvidersSetting(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mcuc.stark.morgan.Morgan/UpdateGetCardProvidersSetting",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MorganServer).UpdateGetCardProvidersSetting(ctx, req.(*UpdateGetCardProvidersSettingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Morgan_serviceDesc = grpc.ServiceDesc{
	ServiceName: "mcuc.stark.morgan.Morgan",
	HandlerType: (*MorganServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateTelcoWithdraw",
			Handler:    _Morgan_CreateTelcoWithdraw_Handler,
		},
		{
			MethodName: "CreateTelcoTopUp",
			Handler:    _Morgan_CreateTelcoTopUp_Handler,
		},
		{
			MethodName: "CancelTelcoWithdraw",
			Handler:    _Morgan_CancelTelcoWithdraw_Handler,
		},
		{
			MethodName: "CompleteTelcoTopUp",
			Handler:    _Morgan_CompleteTelcoTopUp_Handler,
		},
		{
			MethodName: "ApproveTelcoTopUp",
			Handler:    _Morgan_ApproveTelcoTopUp_Handler,
		},
		{
			MethodName: "ApproveTelcoWithdraw",
			Handler:    _Morgan_ApproveTelcoWithdraw_Handler,
		},
		{
			MethodName: "RejectTelcoTopUp",
			Handler:    _Morgan_RejectTelcoTopUp_Handler,
		},
		{
			MethodName: "RejectTelcoWithdraw",
			Handler:    _Morgan_RejectTelcoWithdraw_Handler,
		},
		{
			MethodName: "GetSettings",
			Handler:    _Morgan_GetSettings_Handler,
		},
		{
			MethodName: "UpdateTopUpAutoApprovalSetting",
			Handler:    _Morgan_UpdateTopUpAutoApprovalSetting_Handler,
		},
		{
			MethodName: "UpdateUsingThirdPartySetting",
			Handler:    _Morgan_UpdateUsingThirdPartySetting_Handler,
		},
		{
			MethodName: "UpdateChargeCardProvidersSetting",
			Handler:    _Morgan_UpdateChargeCardProvidersSetting_Handler,
		},
		{
			MethodName: "UpdateGetCardProvidersSetting",
			Handler:    _Morgan_UpdateGetCardProvidersSetting_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "stark/api/morgan.proto",
}