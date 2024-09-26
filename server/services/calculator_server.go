package services

import (
	"context"
	"fmt"
	"io"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type calculatorServer struct {
}

// Sum implements CalculatorServer.
func (c calculatorServer) Sum(stream grpc.BidiStreamingServer[SumRequest, SumResponse]) error {
	sum := int32(0)
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			break
		}

		if err != nil {
			return err
		}

		sum += req.Number
		res := SumResponse{
			Result: sum,
		}

		err = stream.Send(&res)
		if err != nil {
			return err
		}
	}
	return nil
}

// Average implements CalculatorServer.
func (c calculatorServer) Average(stream grpc.ClientStreamingServer[AverageRequest, AverageResponse]) error {
	sum := 0.0
	count := 0.0
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			break
		}

		if err != nil {
			return err
		}

		sum += req.Number
		count++
	}
	res := AverageResponse{
		Result: sum / count,
	}
	return stream.SendAndClose(&res)
}

// Fibonacci implements CalculatorServer.
func (c calculatorServer) Fibonacci(req *FibonacciRequest, stream grpc.ServerStreamingServer[FibonacciResponse]) error {
	for n := uint32(0); n <= req.N; n++ {
		result := fib(n)
		res := FibonacciResponse{
			Result: result,
		}
		fmt.Printf("send %v\n", res)
		stream.Send(&res)
		time.Sleep(time.Second)
	}
	return nil
}

func fib(n uint32) uint32 {
	switch n {
	case 0:
		return 0
	case 1:
		return 1
	default:
		return fib(n-1) + fib(n-2)
	}
}

// Hello implements CalculatorServer.
func (c calculatorServer) Hello(ctx context.Context, req *HelloRequest) (*HelloResponse, error) {

if req.Name == "" {
	return nil, status.Errorf(
		codes.InvalidArgument,
		"name is required",
	)
}

	result := fmt.Sprintf("Hello %v at %v ", req.Name, req.CreatedDate.AsTime().Local())
	res := HelloResponse{
		Result: result,
	}
	return &res, nil
}

// mustEmbedUnimplementedCalculatorServer implements CalculatorServer.
func (c calculatorServer) mustEmbedUnimplementedCalculatorServer() {
	panic("unimplemented")
}

func NewCalculatorServer() CalculatorServer {
	return calculatorServer{}
}
