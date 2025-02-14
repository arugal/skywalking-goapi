// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package v3

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	v3 "skywalking.apache.org/repo/goapi/collect/common/v3"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// EBPFProcessServiceClient is the client API for EBPFProcessService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type EBPFProcessServiceClient interface {
	// Report discovered process in Rover
	ReportProcesses(ctx context.Context, in *EBPFProcessReportList, opts ...grpc.CallOption) (*EBPFReportProcessDownstream, error)
	// Keep the process alive in the backend.
	KeepAlive(ctx context.Context, in *EBPFProcessPingPkgList, opts ...grpc.CallOption) (*v3.Commands, error)
}

type eBPFProcessServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewEBPFProcessServiceClient(cc grpc.ClientConnInterface) EBPFProcessServiceClient {
	return &eBPFProcessServiceClient{cc}
}

func (c *eBPFProcessServiceClient) ReportProcesses(ctx context.Context, in *EBPFProcessReportList, opts ...grpc.CallOption) (*EBPFReportProcessDownstream, error) {
	out := new(EBPFReportProcessDownstream)
	err := c.cc.Invoke(ctx, "/skywalking.v3.EBPFProcessService/reportProcesses", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eBPFProcessServiceClient) KeepAlive(ctx context.Context, in *EBPFProcessPingPkgList, opts ...grpc.CallOption) (*v3.Commands, error) {
	out := new(v3.Commands)
	err := c.cc.Invoke(ctx, "/skywalking.v3.EBPFProcessService/keepAlive", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EBPFProcessServiceServer is the server API for EBPFProcessService service.
// All implementations must embed UnimplementedEBPFProcessServiceServer
// for forward compatibility
type EBPFProcessServiceServer interface {
	// Report discovered process in Rover
	ReportProcesses(context.Context, *EBPFProcessReportList) (*EBPFReportProcessDownstream, error)
	// Keep the process alive in the backend.
	KeepAlive(context.Context, *EBPFProcessPingPkgList) (*v3.Commands, error)
	mustEmbedUnimplementedEBPFProcessServiceServer()
}

// UnimplementedEBPFProcessServiceServer must be embedded to have forward compatible implementations.
type UnimplementedEBPFProcessServiceServer struct {
}

func (UnimplementedEBPFProcessServiceServer) ReportProcesses(context.Context, *EBPFProcessReportList) (*EBPFReportProcessDownstream, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReportProcesses not implemented")
}
func (UnimplementedEBPFProcessServiceServer) KeepAlive(context.Context, *EBPFProcessPingPkgList) (*v3.Commands, error) {
	return nil, status.Errorf(codes.Unimplemented, "method KeepAlive not implemented")
}
func (UnimplementedEBPFProcessServiceServer) mustEmbedUnimplementedEBPFProcessServiceServer() {}

// UnsafeEBPFProcessServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to EBPFProcessServiceServer will
// result in compilation errors.
type UnsafeEBPFProcessServiceServer interface {
	mustEmbedUnimplementedEBPFProcessServiceServer()
}

func RegisterEBPFProcessServiceServer(s grpc.ServiceRegistrar, srv EBPFProcessServiceServer) {
	s.RegisterService(&EBPFProcessService_ServiceDesc, srv)
}

func _EBPFProcessService_ReportProcesses_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EBPFProcessReportList)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EBPFProcessServiceServer).ReportProcesses(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/skywalking.v3.EBPFProcessService/reportProcesses",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EBPFProcessServiceServer).ReportProcesses(ctx, req.(*EBPFProcessReportList))
	}
	return interceptor(ctx, in, info, handler)
}

func _EBPFProcessService_KeepAlive_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EBPFProcessPingPkgList)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EBPFProcessServiceServer).KeepAlive(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/skywalking.v3.EBPFProcessService/keepAlive",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EBPFProcessServiceServer).KeepAlive(ctx, req.(*EBPFProcessPingPkgList))
	}
	return interceptor(ctx, in, info, handler)
}

// EBPFProcessService_ServiceDesc is the grpc.ServiceDesc for EBPFProcessService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var EBPFProcessService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "skywalking.v3.EBPFProcessService",
	HandlerType: (*EBPFProcessServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "reportProcesses",
			Handler:    _EBPFProcessService_ReportProcesses_Handler,
		},
		{
			MethodName: "keepAlive",
			Handler:    _EBPFProcessService_KeepAlive_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ebpf/profiling/Process.proto",
}
