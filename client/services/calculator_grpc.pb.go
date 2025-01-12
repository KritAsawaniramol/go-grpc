// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.28.1
// source: calculator.proto

package services

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
	Calculator_Hello_FullMethodName     = "/services.Calculator/Hello"
	Calculator_Fibonacci_FullMethodName = "/services.Calculator/Fibonacci"
	Calculator_Average_FullMethodName   = "/services.Calculator/Average"
	Calculator_Sum_FullMethodName       = "/services.Calculator/Sum"
)

// CalculatorClient is the client API for Calculator service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CalculatorClient interface {
	// Unary
	Hello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloResponse, error)
	// Server streaming
	Fibonacci(ctx context.Context, in *FibonacciRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[FibonacciResponse], error)
	// Client streaming
	Average(ctx context.Context, opts ...grpc.CallOption) (grpc.ClientStreamingClient[AverageRequest, AverageResponse], error)
	// Bi Directional streaming
	Sum(ctx context.Context, opts ...grpc.CallOption) (grpc.BidiStreamingClient[SumRequest, SumResponse], error)
}

type calculatorClient struct {
	cc grpc.ClientConnInterface
}

func NewCalculatorClient(cc grpc.ClientConnInterface) CalculatorClient {
	return &calculatorClient{cc}
}

func (c *calculatorClient) Hello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(HelloResponse)
	err := c.cc.Invoke(ctx, Calculator_Hello_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *calculatorClient) Fibonacci(ctx context.Context, in *FibonacciRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[FibonacciResponse], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &Calculator_ServiceDesc.Streams[0], Calculator_Fibonacci_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[FibonacciRequest, FibonacciResponse]{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type Calculator_FibonacciClient = grpc.ServerStreamingClient[FibonacciResponse]

func (c *calculatorClient) Average(ctx context.Context, opts ...grpc.CallOption) (grpc.ClientStreamingClient[AverageRequest, AverageResponse], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &Calculator_ServiceDesc.Streams[1], Calculator_Average_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[AverageRequest, AverageResponse]{ClientStream: stream}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type Calculator_AverageClient = grpc.ClientStreamingClient[AverageRequest, AverageResponse]

func (c *calculatorClient) Sum(ctx context.Context, opts ...grpc.CallOption) (grpc.BidiStreamingClient[SumRequest, SumResponse], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &Calculator_ServiceDesc.Streams[2], Calculator_Sum_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[SumRequest, SumResponse]{ClientStream: stream}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type Calculator_SumClient = grpc.BidiStreamingClient[SumRequest, SumResponse]

// CalculatorServer is the server API for Calculator service.
// All implementations must embed UnimplementedCalculatorServer
// for forward compatibility.
type CalculatorServer interface {
	// Unary
	Hello(context.Context, *HelloRequest) (*HelloResponse, error)
	// Server streaming
	Fibonacci(*FibonacciRequest, grpc.ServerStreamingServer[FibonacciResponse]) error
	// Client streaming
	Average(grpc.ClientStreamingServer[AverageRequest, AverageResponse]) error
	// Bi Directional streaming
	Sum(grpc.BidiStreamingServer[SumRequest, SumResponse]) error
	mustEmbedUnimplementedCalculatorServer()
}

// UnimplementedCalculatorServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedCalculatorServer struct{}

func (UnimplementedCalculatorServer) Hello(context.Context, *HelloRequest) (*HelloResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Hello not implemented")
}
func (UnimplementedCalculatorServer) Fibonacci(*FibonacciRequest, grpc.ServerStreamingServer[FibonacciResponse]) error {
	return status.Errorf(codes.Unimplemented, "method Fibonacci not implemented")
}
func (UnimplementedCalculatorServer) Average(grpc.ClientStreamingServer[AverageRequest, AverageResponse]) error {
	return status.Errorf(codes.Unimplemented, "method Average not implemented")
}
func (UnimplementedCalculatorServer) Sum(grpc.BidiStreamingServer[SumRequest, SumResponse]) error {
	return status.Errorf(codes.Unimplemented, "method Sum not implemented")
}
func (UnimplementedCalculatorServer) mustEmbedUnimplementedCalculatorServer() {}
func (UnimplementedCalculatorServer) testEmbeddedByValue()                    {}

// UnsafeCalculatorServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CalculatorServer will
// result in compilation errors.
type UnsafeCalculatorServer interface {
	mustEmbedUnimplementedCalculatorServer()
}

func RegisterCalculatorServer(s grpc.ServiceRegistrar, srv CalculatorServer) {
	// If the following call pancis, it indicates UnimplementedCalculatorServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Calculator_ServiceDesc, srv)
}

func _Calculator_Hello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalculatorServer).Hello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Calculator_Hello_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalculatorServer).Hello(ctx, req.(*HelloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Calculator_Fibonacci_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(FibonacciRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(CalculatorServer).Fibonacci(m, &grpc.GenericServerStream[FibonacciRequest, FibonacciResponse]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type Calculator_FibonacciServer = grpc.ServerStreamingServer[FibonacciResponse]

func _Calculator_Average_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(CalculatorServer).Average(&grpc.GenericServerStream[AverageRequest, AverageResponse]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type Calculator_AverageServer = grpc.ClientStreamingServer[AverageRequest, AverageResponse]

func _Calculator_Sum_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(CalculatorServer).Sum(&grpc.GenericServerStream[SumRequest, SumResponse]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type Calculator_SumServer = grpc.BidiStreamingServer[SumRequest, SumResponse]

// Calculator_ServiceDesc is the grpc.ServiceDesc for Calculator service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Calculator_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "services.Calculator",
	HandlerType: (*CalculatorServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Hello",
			Handler:    _Calculator_Hello_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Fibonacci",
			Handler:       _Calculator_Fibonacci_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "Average",
			Handler:       _Calculator_Average_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "Sum",
			Handler:       _Calculator_Sum_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "calculator.proto",
}
