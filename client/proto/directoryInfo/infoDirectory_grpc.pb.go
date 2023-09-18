// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: client/proto/directoryInfo/infoDirectory.proto

package directoryInfo

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

// InfoDirectoryClient is the client API for InfoDirectory service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type InfoDirectoryClient interface {
	InfoDir(ctx context.Context, in *PathRequest, opts ...grpc.CallOption) (*DirInfoResponse, error)
	InfoDirStreamClient(ctx context.Context, opts ...grpc.CallOption) (InfoDirectory_InfoDirStreamClientClient, error)
}

type infoDirectoryClient struct {
	cc grpc.ClientConnInterface
}

func NewInfoDirectoryClient(cc grpc.ClientConnInterface) InfoDirectoryClient {
	return &infoDirectoryClient{cc}
}

func (c *infoDirectoryClient) InfoDir(ctx context.Context, in *PathRequest, opts ...grpc.CallOption) (*DirInfoResponse, error) {
	out := new(DirInfoResponse)
	err := c.cc.Invoke(ctx, "/InfoDirectory/InfoDir", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *infoDirectoryClient) InfoDirStreamClient(ctx context.Context, opts ...grpc.CallOption) (InfoDirectory_InfoDirStreamClientClient, error) {
	stream, err := c.cc.NewStream(ctx, &InfoDirectory_ServiceDesc.Streams[0], "/InfoDirectory/InfoDirStreamClient", opts...)
	if err != nil {
		return nil, err
	}
	x := &infoDirectoryInfoDirStreamClientClient{stream}
	return x, nil
}

type InfoDirectory_InfoDirStreamClientClient interface {
	Send(*PathRequest) error
	CloseAndRecv() (*DirInfoStreamClientResponse, error)
	grpc.ClientStream
}

type infoDirectoryInfoDirStreamClientClient struct {
	grpc.ClientStream
}

func (x *infoDirectoryInfoDirStreamClientClient) Send(m *PathRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *infoDirectoryInfoDirStreamClientClient) CloseAndRecv() (*DirInfoStreamClientResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(DirInfoStreamClientResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// InfoDirectoryServer is the server API for InfoDirectory service.
// All implementations must embed UnimplementedInfoDirectoryServer
// for forward compatibility
type InfoDirectoryServer interface {
	InfoDir(context.Context, *PathRequest) (*DirInfoResponse, error)
	InfoDirStreamClient(InfoDirectory_InfoDirStreamClientServer) error
	mustEmbedUnimplementedInfoDirectoryServer()
}

// UnimplementedInfoDirectoryServer must be embedded to have forward compatible implementations.
type UnimplementedInfoDirectoryServer struct {
}

func (UnimplementedInfoDirectoryServer) InfoDir(context.Context, *PathRequest) (*DirInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InfoDir not implemented")
}
func (UnimplementedInfoDirectoryServer) InfoDirStreamClient(InfoDirectory_InfoDirStreamClientServer) error {
	return status.Errorf(codes.Unimplemented, "method InfoDirStreamClient not implemented")
}
func (UnimplementedInfoDirectoryServer) mustEmbedUnimplementedInfoDirectoryServer() {}

// UnsafeInfoDirectoryServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to InfoDirectoryServer will
// result in compilation errors.
type UnsafeInfoDirectoryServer interface {
	mustEmbedUnimplementedInfoDirectoryServer()
}

func RegisterInfoDirectoryServer(s grpc.ServiceRegistrar, srv InfoDirectoryServer) {
	s.RegisterService(&InfoDirectory_ServiceDesc, srv)
}

func _InfoDirectory_InfoDir_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PathRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InfoDirectoryServer).InfoDir(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/InfoDirectory/InfoDir",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InfoDirectoryServer).InfoDir(ctx, req.(*PathRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InfoDirectory_InfoDirStreamClient_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(InfoDirectoryServer).InfoDirStreamClient(&infoDirectoryInfoDirStreamClientServer{stream})
}

type InfoDirectory_InfoDirStreamClientServer interface {
	SendAndClose(*DirInfoStreamClientResponse) error
	Recv() (*PathRequest, error)
	grpc.ServerStream
}

type infoDirectoryInfoDirStreamClientServer struct {
	grpc.ServerStream
}

func (x *infoDirectoryInfoDirStreamClientServer) SendAndClose(m *DirInfoStreamClientResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *infoDirectoryInfoDirStreamClientServer) Recv() (*PathRequest, error) {
	m := new(PathRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// InfoDirectory_ServiceDesc is the grpc.ServiceDesc for InfoDirectory service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var InfoDirectory_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "InfoDirectory",
	HandlerType: (*InfoDirectoryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "InfoDir",
			Handler:    _InfoDirectory_InfoDir_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "InfoDirStreamClient",
			Handler:       _InfoDirectory_InfoDirStreamClient_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "client/proto/directoryInfo/infoDirectory.proto",
}
