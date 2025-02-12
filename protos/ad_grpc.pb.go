// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v3.12.4
// source: protos/ad.proto

package Yektanet

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.62.0 or later.
const _ = grpc.SupportPackageIsVersion8

const (
	AdRetriever_GetAds_FullMethodName = "/Yektanet.AdRetriever/get_ads"
)

// AdRetrieverClient is the client API for AdRetriever service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AdRetrieverClient interface {
	GetAds(ctx context.Context, in *TargetingRequest, opts ...grpc.CallOption) (*TargetingResponse, error)
}

type adRetrieverClient struct {
	cc grpc.ClientConnInterface
}

func NewAdRetrieverClient(cc grpc.ClientConnInterface) AdRetrieverClient {
	return &adRetrieverClient{cc}
}

func (c *adRetrieverClient) GetAds(ctx context.Context, in *TargetingRequest, opts ...grpc.CallOption) (*TargetingResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(TargetingResponse)
	err := c.cc.Invoke(ctx, AdRetriever_GetAds_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AdRetrieverServer is the server API for AdRetriever service.
// All implementations must embed UnimplementedAdRetrieverServer
// for forward compatibility
type AdRetrieverServer interface {
	GetAds(context.Context, *TargetingRequest) (*TargetingResponse, error)
	mustEmbedUnimplementedAdRetrieverServer()
}

// UnimplementedAdRetrieverServer must be embedded to have forward compatible implementations.
type UnimplementedAdRetrieverServer struct {
}

func (UnimplementedAdRetrieverServer) GetAds(context.Context, *TargetingRequest) (*TargetingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAds not implemented")
}
func (UnimplementedAdRetrieverServer) mustEmbedUnimplementedAdRetrieverServer() {}

// UnsafeAdRetrieverServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AdRetrieverServer will
// result in compilation errors.
type UnsafeAdRetrieverServer interface {
	mustEmbedUnimplementedAdRetrieverServer()
}

func RegisterAdRetrieverServer(s grpc.ServiceRegistrar, srv AdRetrieverServer) {
	s.RegisterService(&AdRetriever_ServiceDesc, srv)
}

func _AdRetriever_GetAds_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TargetingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdRetrieverServer).GetAds(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AdRetriever_GetAds_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdRetrieverServer).GetAds(ctx, req.(*TargetingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AdRetriever_ServiceDesc is the grpc.ServiceDesc for AdRetriever service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AdRetriever_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Yektanet.AdRetriever",
	HandlerType: (*AdRetrieverServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "get_ads",
			Handler:    _AdRetriever_GetAds_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protos/ad.proto",
}
