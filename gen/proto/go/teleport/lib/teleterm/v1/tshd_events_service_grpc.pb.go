// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: teleport/lib/teleterm/v1/tshd_events_service.proto

package v1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// TshdEventsServiceClient is the client API for TshdEventsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TshdEventsServiceClient interface {
	// Relogin makes the Electron app display a login modal for the specific root cluster. The request
	// returns a response after the relogin procedure has been successfully finished.
	Relogin(ctx context.Context, in *ReloginRequest, opts ...grpc.CallOption) (*ReloginResponse, error)
	// SendNotification causes the Electron app to display a notification in the UI. The request
	// accepts a specific message rather than a generic string so that the Electron is in control as
	// to what message is displayed and how exactly it looks.
	SendNotification(ctx context.Context, in *SendNotificationRequest, opts ...grpc.CallOption) (*SendNotificationResponse, error)
}

type tshdEventsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTshdEventsServiceClient(cc grpc.ClientConnInterface) TshdEventsServiceClient {
	return &tshdEventsServiceClient{cc}
}

func (c *tshdEventsServiceClient) Relogin(ctx context.Context, in *ReloginRequest, opts ...grpc.CallOption) (*ReloginResponse, error) {
	out := new(ReloginResponse)
	err := c.cc.Invoke(ctx, "/teleport.lib.teleterm.v1.TshdEventsService/Relogin", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tshdEventsServiceClient) SendNotification(ctx context.Context, in *SendNotificationRequest, opts ...grpc.CallOption) (*SendNotificationResponse, error) {
	out := new(SendNotificationResponse)
	err := c.cc.Invoke(ctx, "/teleport.lib.teleterm.v1.TshdEventsService/SendNotification", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TshdEventsServiceServer is the server API for TshdEventsService service.
// All implementations must embed UnimplementedTshdEventsServiceServer
// for forward compatibility
type TshdEventsServiceServer interface {
	// Relogin makes the Electron app display a login modal for the specific root cluster. The request
	// returns a response after the relogin procedure has been successfully finished.
	Relogin(context.Context, *ReloginRequest) (*ReloginResponse, error)
	// SendNotification causes the Electron app to display a notification in the UI. The request
	// accepts a specific message rather than a generic string so that the Electron is in control as
	// to what message is displayed and how exactly it looks.
	SendNotification(context.Context, *SendNotificationRequest) (*SendNotificationResponse, error)
	mustEmbedUnimplementedTshdEventsServiceServer()
}

// UnimplementedTshdEventsServiceServer must be embedded to have forward compatible implementations.
type UnimplementedTshdEventsServiceServer struct {
}

func (UnimplementedTshdEventsServiceServer) Relogin(context.Context, *ReloginRequest) (*ReloginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Relogin not implemented")
}
func (UnimplementedTshdEventsServiceServer) SendNotification(context.Context, *SendNotificationRequest) (*SendNotificationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendNotification not implemented")
}
func (UnimplementedTshdEventsServiceServer) mustEmbedUnimplementedTshdEventsServiceServer() {}

// UnsafeTshdEventsServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TshdEventsServiceServer will
// result in compilation errors.
type UnsafeTshdEventsServiceServer interface {
	mustEmbedUnimplementedTshdEventsServiceServer()
}

func RegisterTshdEventsServiceServer(s grpc.ServiceRegistrar, srv TshdEventsServiceServer) {
	s.RegisterService(&TshdEventsService_ServiceDesc, srv)
}

func _TshdEventsService_Relogin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReloginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TshdEventsServiceServer).Relogin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/teleport.lib.teleterm.v1.TshdEventsService/Relogin",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TshdEventsServiceServer).Relogin(ctx, req.(*ReloginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TshdEventsService_SendNotification_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendNotificationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TshdEventsServiceServer).SendNotification(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/teleport.lib.teleterm.v1.TshdEventsService/SendNotification",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TshdEventsServiceServer).SendNotification(ctx, req.(*SendNotificationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// TshdEventsService_ServiceDesc is the grpc.ServiceDesc for TshdEventsService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TshdEventsService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "teleport.lib.teleterm.v1.TshdEventsService",
	HandlerType: (*TshdEventsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Relogin",
			Handler:    _TshdEventsService_Relogin_Handler,
		},
		{
			MethodName: "SendNotification",
			Handler:    _TshdEventsService_SendNotification_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "teleport/lib/teleterm/v1/tshd_events_service.proto",
}
