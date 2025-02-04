// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: kannon/stats/apiv1/statsapiv1.proto

package apiv1

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

// StatsApiV1Client is the client API for StatsApiV1 service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type StatsApiV1Client interface {
	GetStats(ctx context.Context, in *GetStatsReq, opts ...grpc.CallOption) (*GetStatsRes, error)
	GetStatsAggregated(ctx context.Context, in *GetStatsAggregatedReq, opts ...grpc.CallOption) (*GetStatsAggregatedRes, error)
}

type statsApiV1Client struct {
	cc grpc.ClientConnInterface
}

func NewStatsApiV1Client(cc grpc.ClientConnInterface) StatsApiV1Client {
	return &statsApiV1Client{cc}
}

func (c *statsApiV1Client) GetStats(ctx context.Context, in *GetStatsReq, opts ...grpc.CallOption) (*GetStatsRes, error) {
	out := new(GetStatsRes)
	err := c.cc.Invoke(ctx, "/kannon.StatsApiV1/GetStats", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *statsApiV1Client) GetStatsAggregated(ctx context.Context, in *GetStatsAggregatedReq, opts ...grpc.CallOption) (*GetStatsAggregatedRes, error) {
	out := new(GetStatsAggregatedRes)
	err := c.cc.Invoke(ctx, "/kannon.StatsApiV1/GetStatsAggregated", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StatsApiV1Server is the server API for StatsApiV1 service.
// All implementations should embed UnimplementedStatsApiV1Server
// for forward compatibility
type StatsApiV1Server interface {
	GetStats(context.Context, *GetStatsReq) (*GetStatsRes, error)
	GetStatsAggregated(context.Context, *GetStatsAggregatedReq) (*GetStatsAggregatedRes, error)
}

// UnimplementedStatsApiV1Server should be embedded to have forward compatible implementations.
type UnimplementedStatsApiV1Server struct {
}

func (UnimplementedStatsApiV1Server) GetStats(context.Context, *GetStatsReq) (*GetStatsRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStats not implemented")
}
func (UnimplementedStatsApiV1Server) GetStatsAggregated(context.Context, *GetStatsAggregatedReq) (*GetStatsAggregatedRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStatsAggregated not implemented")
}

// UnsafeStatsApiV1Server may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to StatsApiV1Server will
// result in compilation errors.
type UnsafeStatsApiV1Server interface {
	mustEmbedUnimplementedStatsApiV1Server()
}

func RegisterStatsApiV1Server(s grpc.ServiceRegistrar, srv StatsApiV1Server) {
	s.RegisterService(&StatsApiV1_ServiceDesc, srv)
}

func _StatsApiV1_GetStats_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetStatsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StatsApiV1Server).GetStats(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kannon.StatsApiV1/GetStats",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StatsApiV1Server).GetStats(ctx, req.(*GetStatsReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _StatsApiV1_GetStatsAggregated_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetStatsAggregatedReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StatsApiV1Server).GetStatsAggregated(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kannon.StatsApiV1/GetStatsAggregated",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StatsApiV1Server).GetStatsAggregated(ctx, req.(*GetStatsAggregatedReq))
	}
	return interceptor(ctx, in, info, handler)
}

// StatsApiV1_ServiceDesc is the grpc.ServiceDesc for StatsApiV1 service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var StatsApiV1_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "kannon.StatsApiV1",
	HandlerType: (*StatsApiV1Server)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetStats",
			Handler:    _StatsApiV1_GetStats_Handler,
		},
		{
			MethodName: "GetStatsAggregated",
			Handler:    _StatsApiV1_GetStatsAggregated_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "kannon/stats/apiv1/statsapiv1.proto",
}
