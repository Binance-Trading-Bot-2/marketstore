// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package proto

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

// ReplicationClient is the client API for Replication service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ReplicationClient interface {
	GetWALStream(ctx context.Context, in *GetWALStreamRequest, opts ...grpc.CallOption) (Replication_GetWALStreamClient, error)
}

type replicationClient struct {
	cc grpc.ClientConnInterface
}

func NewReplicationClient(cc grpc.ClientConnInterface) ReplicationClient {
	return &replicationClient{cc}
}

func (c *replicationClient) GetWALStream(ctx context.Context, in *GetWALStreamRequest, opts ...grpc.CallOption) (Replication_GetWALStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &Replication_ServiceDesc.Streams[0], "/proto.Replication/GetWALStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &replicationGetWALStreamClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Replication_GetWALStreamClient interface {
	Recv() (*GetWALStreamResponse, error)
	grpc.ClientStream
}

type replicationGetWALStreamClient struct {
	grpc.ClientStream
}

func (x *replicationGetWALStreamClient) Recv() (*GetWALStreamResponse, error) {
	m := new(GetWALStreamResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ReplicationServer is the server API for Replication service.
// All implementations must embed UnimplementedReplicationServer
// for forward compatibility
type ReplicationServer interface {
	GetWALStream(*GetWALStreamRequest, Replication_GetWALStreamServer) error
	mustEmbedUnimplementedReplicationServer()
}

// UnimplementedReplicationServer must be embedded to have forward compatible implementations.
type UnimplementedReplicationServer struct {
}

func (UnimplementedReplicationServer) GetWALStream(*GetWALStreamRequest, Replication_GetWALStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method GetWALStream not implemented")
}
func (UnimplementedReplicationServer) mustEmbedUnimplementedReplicationServer() {}

// UnsafeReplicationServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ReplicationServer will
// result in compilation errors.
type UnsafeReplicationServer interface {
	mustEmbedUnimplementedReplicationServer()
}

func RegisterReplicationServer(s grpc.ServiceRegistrar, srv ReplicationServer) {
	s.RegisterService(&Replication_ServiceDesc, srv)
}

func _Replication_GetWALStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GetWALStreamRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ReplicationServer).GetWALStream(m, &replicationGetWALStreamServer{stream})
}

type Replication_GetWALStreamServer interface {
	Send(*GetWALStreamResponse) error
	grpc.ServerStream
}

type replicationGetWALStreamServer struct {
	grpc.ServerStream
}

func (x *replicationGetWALStreamServer) Send(m *GetWALStreamResponse) error {
	return x.ServerStream.SendMsg(m)
}

// Replication_ServiceDesc is the grpc.ServiceDesc for Replication service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Replication_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Replication",
	HandlerType: (*ReplicationServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetWALStream",
			Handler:       _Replication_GetWALStream_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "replication.proto",
}
