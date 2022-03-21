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

// TonyClient is the client API for Tony service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TonyClient interface {
	// MEX call to requedt a payment for their user to make payment to e_wallet
	GetEWalletPaymentCode(ctx context.Context, in *GetEWalletPaymentCodeRequest, opts ...grpc.CallOption) (*GetEWalletPaymentCodeReply, error)
	CancelEWalletTopUp(ctx context.Context, in *CancelEWalletTopUpRequest, opts ...grpc.CallOption) (*CancelEWalletTopUpReply, error)
	// Teller call it when recevie sms or check e_wallet account
	CreateEWalletTopUp(ctx context.Context, in *CreateEWalletTopUpRequest, opts ...grpc.CallOption) (*CreateEWalletTopUpReply, error)
	// Teller approve it
	ApproveEWalletTopUp(ctx context.Context, in *ApproveEWalletTopUpRequest, opts ...grpc.CallOption) (*ApproveEWalletTopUpReply, error)
	// Teller reject it
	RejectEWalletTopUp(ctx context.Context, in *RejectEWalletTopUpRequest, opts ...grpc.CallOption) (*RejectEWalletTopUpReply, error)
	// Utility API to call for process payment
	GetSystemEWallets(ctx context.Context, in *GetSystemEWalletsRequest, opts ...grpc.CallOption) (*GetSystemEWalletsReply, error)
	// APIs for manage System EWallet for topup
	CreateSystemEWallet(ctx context.Context, in *CreateSystemEWalletRequest, opts ...grpc.CallOption) (*CreateSystemEWalletReply, error)
	UpdateSystemEWalletStatus(ctx context.Context, in *UpdateSystemEWalletStatusRequest, opts ...grpc.CallOption) (*UpdateSystemEWalletStatusReply, error)
	ListSystemEWallets(ctx context.Context, in *ListSystemEWalletsRequest, opts ...grpc.CallOption) (*ListSystemEWalletsReply, error)
	ValidateSystemEWallets(ctx context.Context, in *ValidateSystemEWalletsRequest, opts ...grpc.CallOption) (*ValidateSystemEWalletsReply, error)
	ImportSystemEWallets(ctx context.Context, in *ImportSystemEWalletsRequest, opts ...grpc.CallOption) (*ImportSystemEWalletsReply, error)
}

type tonyClient struct {
	cc grpc.ClientConnInterface
}

func NewTonyClient(cc grpc.ClientConnInterface) TonyClient {
	return &tonyClient{cc}
}

func (c *tonyClient) GetEWalletPaymentCode(ctx context.Context, in *GetEWalletPaymentCodeRequest, opts ...grpc.CallOption) (*GetEWalletPaymentCodeReply, error) {
	out := new(GetEWalletPaymentCodeReply)
	err := c.cc.Invoke(ctx, "/mcuc.stark.tony.Tony/GetEWalletPaymentCode", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tonyClient) CancelEWalletTopUp(ctx context.Context, in *CancelEWalletTopUpRequest, opts ...grpc.CallOption) (*CancelEWalletTopUpReply, error) {
	out := new(CancelEWalletTopUpReply)
	err := c.cc.Invoke(ctx, "/mcuc.stark.tony.Tony/CancelEWalletTopUp", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tonyClient) CreateEWalletTopUp(ctx context.Context, in *CreateEWalletTopUpRequest, opts ...grpc.CallOption) (*CreateEWalletTopUpReply, error) {
	out := new(CreateEWalletTopUpReply)
	err := c.cc.Invoke(ctx, "/mcuc.stark.tony.Tony/CreateEWalletTopUp", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tonyClient) ApproveEWalletTopUp(ctx context.Context, in *ApproveEWalletTopUpRequest, opts ...grpc.CallOption) (*ApproveEWalletTopUpReply, error) {
	out := new(ApproveEWalletTopUpReply)
	err := c.cc.Invoke(ctx, "/mcuc.stark.tony.Tony/ApproveEWalletTopUp", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tonyClient) RejectEWalletTopUp(ctx context.Context, in *RejectEWalletTopUpRequest, opts ...grpc.CallOption) (*RejectEWalletTopUpReply, error) {
	out := new(RejectEWalletTopUpReply)
	err := c.cc.Invoke(ctx, "/mcuc.stark.tony.Tony/RejectEWalletTopUp", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tonyClient) GetSystemEWallets(ctx context.Context, in *GetSystemEWalletsRequest, opts ...grpc.CallOption) (*GetSystemEWalletsReply, error) {
	out := new(GetSystemEWalletsReply)
	err := c.cc.Invoke(ctx, "/mcuc.stark.tony.Tony/GetSystemEWallets", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tonyClient) CreateSystemEWallet(ctx context.Context, in *CreateSystemEWalletRequest, opts ...grpc.CallOption) (*CreateSystemEWalletReply, error) {
	out := new(CreateSystemEWalletReply)
	err := c.cc.Invoke(ctx, "/mcuc.stark.tony.Tony/CreateSystemEWallet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tonyClient) UpdateSystemEWalletStatus(ctx context.Context, in *UpdateSystemEWalletStatusRequest, opts ...grpc.CallOption) (*UpdateSystemEWalletStatusReply, error) {
	out := new(UpdateSystemEWalletStatusReply)
	err := c.cc.Invoke(ctx, "/mcuc.stark.tony.Tony/UpdateSystemEWalletStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tonyClient) ListSystemEWallets(ctx context.Context, in *ListSystemEWalletsRequest, opts ...grpc.CallOption) (*ListSystemEWalletsReply, error) {
	out := new(ListSystemEWalletsReply)
	err := c.cc.Invoke(ctx, "/mcuc.stark.tony.Tony/ListSystemEWallets", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tonyClient) ValidateSystemEWallets(ctx context.Context, in *ValidateSystemEWalletsRequest, opts ...grpc.CallOption) (*ValidateSystemEWalletsReply, error) {
	out := new(ValidateSystemEWalletsReply)
	err := c.cc.Invoke(ctx, "/mcuc.stark.tony.Tony/ValidateSystemEWallets", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tonyClient) ImportSystemEWallets(ctx context.Context, in *ImportSystemEWalletsRequest, opts ...grpc.CallOption) (*ImportSystemEWalletsReply, error) {
	out := new(ImportSystemEWalletsReply)
	err := c.cc.Invoke(ctx, "/mcuc.stark.tony.Tony/ImportSystemEWallets", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TonyServer is the server API for Tony service.
// All implementations must embed UnimplementedTonyServer
// for forward compatibility
type TonyServer interface {
	// MEX call to requedt a payment for their user to make payment to e_wallet
	GetEWalletPaymentCode(context.Context, *GetEWalletPaymentCodeRequest) (*GetEWalletPaymentCodeReply, error)
	CancelEWalletTopUp(context.Context, *CancelEWalletTopUpRequest) (*CancelEWalletTopUpReply, error)
	// Teller call it when recevie sms or check e_wallet account
	CreateEWalletTopUp(context.Context, *CreateEWalletTopUpRequest) (*CreateEWalletTopUpReply, error)
	// Teller approve it
	ApproveEWalletTopUp(context.Context, *ApproveEWalletTopUpRequest) (*ApproveEWalletTopUpReply, error)
	// Teller reject it
	RejectEWalletTopUp(context.Context, *RejectEWalletTopUpRequest) (*RejectEWalletTopUpReply, error)
	// Utility API to call for process payment
	GetSystemEWallets(context.Context, *GetSystemEWalletsRequest) (*GetSystemEWalletsReply, error)
	// APIs for manage System EWallet for topup
	CreateSystemEWallet(context.Context, *CreateSystemEWalletRequest) (*CreateSystemEWalletReply, error)
	UpdateSystemEWalletStatus(context.Context, *UpdateSystemEWalletStatusRequest) (*UpdateSystemEWalletStatusReply, error)
	ListSystemEWallets(context.Context, *ListSystemEWalletsRequest) (*ListSystemEWalletsReply, error)
	ValidateSystemEWallets(context.Context, *ValidateSystemEWalletsRequest) (*ValidateSystemEWalletsReply, error)
	ImportSystemEWallets(context.Context, *ImportSystemEWalletsRequest) (*ImportSystemEWalletsReply, error)
	mustEmbedUnimplementedTonyServer()
}

// UnimplementedTonyServer must be embedded to have forward compatible implementations.
type UnimplementedTonyServer struct {
}

func (*UnimplementedTonyServer) GetEWalletPaymentCode(context.Context, *GetEWalletPaymentCodeRequest) (*GetEWalletPaymentCodeReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetEWalletPaymentCode not implemented")
}
func (*UnimplementedTonyServer) CancelEWalletTopUp(context.Context, *CancelEWalletTopUpRequest) (*CancelEWalletTopUpReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CancelEWalletTopUp not implemented")
}
func (*UnimplementedTonyServer) CreateEWalletTopUp(context.Context, *CreateEWalletTopUpRequest) (*CreateEWalletTopUpReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateEWalletTopUp not implemented")
}
func (*UnimplementedTonyServer) ApproveEWalletTopUp(context.Context, *ApproveEWalletTopUpRequest) (*ApproveEWalletTopUpReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ApproveEWalletTopUp not implemented")
}
func (*UnimplementedTonyServer) RejectEWalletTopUp(context.Context, *RejectEWalletTopUpRequest) (*RejectEWalletTopUpReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RejectEWalletTopUp not implemented")
}
func (*UnimplementedTonyServer) GetSystemEWallets(context.Context, *GetSystemEWalletsRequest) (*GetSystemEWalletsReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSystemEWallets not implemented")
}
func (*UnimplementedTonyServer) CreateSystemEWallet(context.Context, *CreateSystemEWalletRequest) (*CreateSystemEWalletReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateSystemEWallet not implemented")
}
func (*UnimplementedTonyServer) UpdateSystemEWalletStatus(context.Context, *UpdateSystemEWalletStatusRequest) (*UpdateSystemEWalletStatusReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateSystemEWalletStatus not implemented")
}
func (*UnimplementedTonyServer) ListSystemEWallets(context.Context, *ListSystemEWalletsRequest) (*ListSystemEWalletsReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListSystemEWallets not implemented")
}
func (*UnimplementedTonyServer) ValidateSystemEWallets(context.Context, *ValidateSystemEWalletsRequest) (*ValidateSystemEWalletsReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ValidateSystemEWallets not implemented")
}
func (*UnimplementedTonyServer) ImportSystemEWallets(context.Context, *ImportSystemEWalletsRequest) (*ImportSystemEWalletsReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ImportSystemEWallets not implemented")
}
func (*UnimplementedTonyServer) mustEmbedUnimplementedTonyServer() {}

func RegisterTonyServer(s *grpc.Server, srv TonyServer) {
	s.RegisterService(&_Tony_serviceDesc, srv)
}

func _Tony_GetEWalletPaymentCode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetEWalletPaymentCodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TonyServer).GetEWalletPaymentCode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mcuc.stark.tony.Tony/GetEWalletPaymentCode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TonyServer).GetEWalletPaymentCode(ctx, req.(*GetEWalletPaymentCodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Tony_CancelEWalletTopUp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CancelEWalletTopUpRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TonyServer).CancelEWalletTopUp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mcuc.stark.tony.Tony/CancelEWalletTopUp",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TonyServer).CancelEWalletTopUp(ctx, req.(*CancelEWalletTopUpRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Tony_CreateEWalletTopUp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateEWalletTopUpRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TonyServer).CreateEWalletTopUp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mcuc.stark.tony.Tony/CreateEWalletTopUp",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TonyServer).CreateEWalletTopUp(ctx, req.(*CreateEWalletTopUpRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Tony_ApproveEWalletTopUp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ApproveEWalletTopUpRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TonyServer).ApproveEWalletTopUp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mcuc.stark.tony.Tony/ApproveEWalletTopUp",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TonyServer).ApproveEWalletTopUp(ctx, req.(*ApproveEWalletTopUpRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Tony_RejectEWalletTopUp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RejectEWalletTopUpRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TonyServer).RejectEWalletTopUp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mcuc.stark.tony.Tony/RejectEWalletTopUp",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TonyServer).RejectEWalletTopUp(ctx, req.(*RejectEWalletTopUpRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Tony_GetSystemEWallets_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSystemEWalletsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TonyServer).GetSystemEWallets(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mcuc.stark.tony.Tony/GetSystemEWallets",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TonyServer).GetSystemEWallets(ctx, req.(*GetSystemEWalletsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Tony_CreateSystemEWallet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateSystemEWalletRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TonyServer).CreateSystemEWallet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mcuc.stark.tony.Tony/CreateSystemEWallet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TonyServer).CreateSystemEWallet(ctx, req.(*CreateSystemEWalletRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Tony_UpdateSystemEWalletStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateSystemEWalletStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TonyServer).UpdateSystemEWalletStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mcuc.stark.tony.Tony/UpdateSystemEWalletStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TonyServer).UpdateSystemEWalletStatus(ctx, req.(*UpdateSystemEWalletStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Tony_ListSystemEWallets_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListSystemEWalletsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TonyServer).ListSystemEWallets(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mcuc.stark.tony.Tony/ListSystemEWallets",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TonyServer).ListSystemEWallets(ctx, req.(*ListSystemEWalletsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Tony_ValidateSystemEWallets_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ValidateSystemEWalletsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TonyServer).ValidateSystemEWallets(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mcuc.stark.tony.Tony/ValidateSystemEWallets",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TonyServer).ValidateSystemEWallets(ctx, req.(*ValidateSystemEWalletsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Tony_ImportSystemEWallets_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ImportSystemEWalletsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TonyServer).ImportSystemEWallets(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mcuc.stark.tony.Tony/ImportSystemEWallets",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TonyServer).ImportSystemEWallets(ctx, req.(*ImportSystemEWalletsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Tony_serviceDesc = grpc.ServiceDesc{
	ServiceName: "mcuc.stark.tony.Tony",
	HandlerType: (*TonyServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetEWalletPaymentCode",
			Handler:    _Tony_GetEWalletPaymentCode_Handler,
		},
		{
			MethodName: "CancelEWalletTopUp",
			Handler:    _Tony_CancelEWalletTopUp_Handler,
		},
		{
			MethodName: "CreateEWalletTopUp",
			Handler:    _Tony_CreateEWalletTopUp_Handler,
		},
		{
			MethodName: "ApproveEWalletTopUp",
			Handler:    _Tony_ApproveEWalletTopUp_Handler,
		},
		{
			MethodName: "RejectEWalletTopUp",
			Handler:    _Tony_RejectEWalletTopUp_Handler,
		},
		{
			MethodName: "GetSystemEWallets",
			Handler:    _Tony_GetSystemEWallets_Handler,
		},
		{
			MethodName: "CreateSystemEWallet",
			Handler:    _Tony_CreateSystemEWallet_Handler,
		},
		{
			MethodName: "UpdateSystemEWalletStatus",
			Handler:    _Tony_UpdateSystemEWalletStatus_Handler,
		},
		{
			MethodName: "ListSystemEWallets",
			Handler:    _Tony_ListSystemEWallets_Handler,
		},
		{
			MethodName: "ValidateSystemEWallets",
			Handler:    _Tony_ValidateSystemEWallets_Handler,
		},
		{
			MethodName: "ImportSystemEWallets",
			Handler:    _Tony_ImportSystemEWallets_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "stark/api/tony.proto",
}
