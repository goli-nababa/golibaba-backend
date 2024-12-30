// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             (unknown)
// source: bank/v1/payment.proto

package bankv1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	PaymentService_ProcessPayment_FullMethodName        = "/bank.v1.PaymentService/ProcessPayment"
	PaymentService_ProcessRefund_FullMethodName         = "/bank.v1.PaymentService/ProcessRefund"
	PaymentService_VerifyPayment_FullMethodName         = "/bank.v1.PaymentService/VerifyPayment"
	PaymentService_RefundPayment_FullMethodName         = "/bank.v1.PaymentService/RefundPayment"
	PaymentService_ChargeWallet_FullMethodName          = "/bank.v1.PaymentService/ChargeWallet"
	PaymentService_WithdrawFromWallet_FullMethodName    = "/bank.v1.PaymentService/WithdrawFromWallet"
	PaymentService_GetTransactionHistory_FullMethodName = "/bank.v1.PaymentService/GetTransactionHistory"
)

// PaymentServiceClient is the client API for PaymentService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PaymentServiceClient interface {
	ProcessPayment(ctx context.Context, in *ProcessPaymentRequest, opts ...grpc.CallOption) (*ProcessPaymentResponse, error)
	ProcessRefund(ctx context.Context, in *ProcessRefundRequest, opts ...grpc.CallOption) (*ProcessRefundResponse, error)
	VerifyPayment(ctx context.Context, in *VerifyPaymentRequest, opts ...grpc.CallOption) (*VerifyPaymentResponse, error)
	RefundPayment(ctx context.Context, in *RefundPaymentRequest, opts ...grpc.CallOption) (*RefundPaymentResponse, error)
	ChargeWallet(ctx context.Context, in *ChargeWalletRequest, opts ...grpc.CallOption) (*ChargeWalletResponse, error)
	WithdrawFromWallet(ctx context.Context, in *WithdrawFromWalletRequest, opts ...grpc.CallOption) (*WithdrawFromWalletResponse, error)
	GetTransactionHistory(ctx context.Context, in *GetTransactionHistoryRequest, opts ...grpc.CallOption) (*GetTransactionHistoryResponse, error)
}

type paymentServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPaymentServiceClient(cc grpc.ClientConnInterface) PaymentServiceClient {
	return &paymentServiceClient{cc}
}

func (c *paymentServiceClient) ProcessPayment(ctx context.Context, in *ProcessPaymentRequest, opts ...grpc.CallOption) (*ProcessPaymentResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ProcessPaymentResponse)
	err := c.cc.Invoke(ctx, PaymentService_ProcessPayment_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *paymentServiceClient) ProcessRefund(ctx context.Context, in *ProcessRefundRequest, opts ...grpc.CallOption) (*ProcessRefundResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ProcessRefundResponse)
	err := c.cc.Invoke(ctx, PaymentService_ProcessRefund_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *paymentServiceClient) VerifyPayment(ctx context.Context, in *VerifyPaymentRequest, opts ...grpc.CallOption) (*VerifyPaymentResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(VerifyPaymentResponse)
	err := c.cc.Invoke(ctx, PaymentService_VerifyPayment_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *paymentServiceClient) RefundPayment(ctx context.Context, in *RefundPaymentRequest, opts ...grpc.CallOption) (*RefundPaymentResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(RefundPaymentResponse)
	err := c.cc.Invoke(ctx, PaymentService_RefundPayment_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *paymentServiceClient) ChargeWallet(ctx context.Context, in *ChargeWalletRequest, opts ...grpc.CallOption) (*ChargeWalletResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ChargeWalletResponse)
	err := c.cc.Invoke(ctx, PaymentService_ChargeWallet_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *paymentServiceClient) WithdrawFromWallet(ctx context.Context, in *WithdrawFromWalletRequest, opts ...grpc.CallOption) (*WithdrawFromWalletResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(WithdrawFromWalletResponse)
	err := c.cc.Invoke(ctx, PaymentService_WithdrawFromWallet_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *paymentServiceClient) GetTransactionHistory(ctx context.Context, in *GetTransactionHistoryRequest, opts ...grpc.CallOption) (*GetTransactionHistoryResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetTransactionHistoryResponse)
	err := c.cc.Invoke(ctx, PaymentService_GetTransactionHistory_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PaymentServiceServer is the server API for PaymentService service.
// All implementations must embed UnimplementedPaymentServiceServer
// for forward compatibility.
type PaymentServiceServer interface {
	ProcessPayment(context.Context, *ProcessPaymentRequest) (*ProcessPaymentResponse, error)
	ProcessRefund(context.Context, *ProcessRefundRequest) (*ProcessRefundResponse, error)
	VerifyPayment(context.Context, *VerifyPaymentRequest) (*VerifyPaymentResponse, error)
	RefundPayment(context.Context, *RefundPaymentRequest) (*RefundPaymentResponse, error)
	ChargeWallet(context.Context, *ChargeWalletRequest) (*ChargeWalletResponse, error)
	WithdrawFromWallet(context.Context, *WithdrawFromWalletRequest) (*WithdrawFromWalletResponse, error)
	GetTransactionHistory(context.Context, *GetTransactionHistoryRequest) (*GetTransactionHistoryResponse, error)
	mustEmbedUnimplementedPaymentServiceServer()
}

// UnimplementedPaymentServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedPaymentServiceServer struct{}

func (UnimplementedPaymentServiceServer) ProcessPayment(context.Context, *ProcessPaymentRequest) (*ProcessPaymentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ProcessPayment not implemented")
}
func (UnimplementedPaymentServiceServer) ProcessRefund(context.Context, *ProcessRefundRequest) (*ProcessRefundResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ProcessRefund not implemented")
}
func (UnimplementedPaymentServiceServer) VerifyPayment(context.Context, *VerifyPaymentRequest) (*VerifyPaymentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VerifyPayment not implemented")
}
func (UnimplementedPaymentServiceServer) RefundPayment(context.Context, *RefundPaymentRequest) (*RefundPaymentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RefundPayment not implemented")
}
func (UnimplementedPaymentServiceServer) ChargeWallet(context.Context, *ChargeWalletRequest) (*ChargeWalletResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChargeWallet not implemented")
}
func (UnimplementedPaymentServiceServer) WithdrawFromWallet(context.Context, *WithdrawFromWalletRequest) (*WithdrawFromWalletResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method WithdrawFromWallet not implemented")
}
func (UnimplementedPaymentServiceServer) GetTransactionHistory(context.Context, *GetTransactionHistoryRequest) (*GetTransactionHistoryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTransactionHistory not implemented")
}
func (UnimplementedPaymentServiceServer) mustEmbedUnimplementedPaymentServiceServer() {}
func (UnimplementedPaymentServiceServer) testEmbeddedByValue()                        {}

// UnsafePaymentServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PaymentServiceServer will
// result in compilation errors.
type UnsafePaymentServiceServer interface {
	mustEmbedUnimplementedPaymentServiceServer()
}

func RegisterPaymentServiceServer(s grpc.ServiceRegistrar, srv PaymentServiceServer) {
	// If the following call pancis, it indicates UnimplementedPaymentServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&PaymentService_ServiceDesc, srv)
}

func _PaymentService_ProcessPayment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProcessPaymentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PaymentServiceServer).ProcessPayment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PaymentService_ProcessPayment_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PaymentServiceServer).ProcessPayment(ctx, req.(*ProcessPaymentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PaymentService_ProcessRefund_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProcessRefundRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PaymentServiceServer).ProcessRefund(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PaymentService_ProcessRefund_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PaymentServiceServer).ProcessRefund(ctx, req.(*ProcessRefundRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PaymentService_VerifyPayment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VerifyPaymentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PaymentServiceServer).VerifyPayment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PaymentService_VerifyPayment_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PaymentServiceServer).VerifyPayment(ctx, req.(*VerifyPaymentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PaymentService_RefundPayment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RefundPaymentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PaymentServiceServer).RefundPayment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PaymentService_RefundPayment_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PaymentServiceServer).RefundPayment(ctx, req.(*RefundPaymentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PaymentService_ChargeWallet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChargeWalletRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PaymentServiceServer).ChargeWallet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PaymentService_ChargeWallet_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PaymentServiceServer).ChargeWallet(ctx, req.(*ChargeWalletRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PaymentService_WithdrawFromWallet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WithdrawFromWalletRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PaymentServiceServer).WithdrawFromWallet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PaymentService_WithdrawFromWallet_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PaymentServiceServer).WithdrawFromWallet(ctx, req.(*WithdrawFromWalletRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PaymentService_GetTransactionHistory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTransactionHistoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PaymentServiceServer).GetTransactionHistory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PaymentService_GetTransactionHistory_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PaymentServiceServer).GetTransactionHistory(ctx, req.(*GetTransactionHistoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// PaymentService_ServiceDesc is the grpc.ServiceDesc for PaymentService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PaymentService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "bank.v1.PaymentService",
	HandlerType: (*PaymentServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ProcessPayment",
			Handler:    _PaymentService_ProcessPayment_Handler,
		},
		{
			MethodName: "ProcessRefund",
			Handler:    _PaymentService_ProcessRefund_Handler,
		},
		{
			MethodName: "VerifyPayment",
			Handler:    _PaymentService_VerifyPayment_Handler,
		},
		{
			MethodName: "RefundPayment",
			Handler:    _PaymentService_RefundPayment_Handler,
		},
		{
			MethodName: "ChargeWallet",
			Handler:    _PaymentService_ChargeWallet_Handler,
		},
		{
			MethodName: "WithdrawFromWallet",
			Handler:    _PaymentService_WithdrawFromWallet_Handler,
		},
		{
			MethodName: "GetTransactionHistory",
			Handler:    _PaymentService_GetTransactionHistory_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "bank/v1/payment.proto",
}
