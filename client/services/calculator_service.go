package services

import (
	context "context"
	"fmt"
	"io"
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"
)

type CalculatorService interface {
	Hello(name string) error
	Fibonacci(n uint32) error
	Average(numbers ...float64) error
	Sum(numbers ...int32) error
}

type calculatorService struct {
	calculatorClient CalculatorClient
}

// Sum implements CalculatorService.
func (c calculatorService) Sum(numbers ...int32) error {
	stream, err := c.calculatorClient.Sum(context.Background())
	if err != nil {
		return err
	}

	// send req
	fmt.Printf("Service : Sum\n")
	go func() {
		for _, number := range numbers {
			req := SumRequest{
				Number: number,
			}
			stream.Send(&req)
			fmt.Printf("Request : %v\n", req.Number)
			time.Sleep(time.Second * 2)
		}
		stream.CloseSend()
	}()

	// recv resp
	done := make(chan bool)
	errs := make(chan error)
	go func() {
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				break
			}

			if err != nil {
				errs <- err
			}
			fmt.Printf("response: %v\n", res.Result)
		}
		done <- true
	}()

	select {
	case <-done:
		return nil
	case err := <-errs:
		return err
	}
}

// Average implements CalculatorService.
// Variadic Functions
func (c calculatorService) Average(numbers ...float64) error {

	stream, err := c.calculatorClient.Average(context.Background())
	if err != nil {
		return err
	}

	fmt.Printf("Service : Average\n")
	for _, number := range numbers {
		req := AverageRequest{
			Number: number,
		}
		stream.Send(&req)
		fmt.Printf("Request : %v\n", req.Number)
		time.Sleep(time.Second)
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		return err
	}
	fmt.Printf("Response: %v\n", res.Result)
	return nil

}

// Fibonacci implements CalculatorService.
func (c calculatorService) Fibonacci(n uint32) error {
	req := FibonacciRequest{
		N: n,
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	stream, err := c.calculatorClient.Fibonacci(ctx, &req)
	if err != nil {
		fmt.Printf("err.Error(): %v\n", err.Error())
		return err
	}

	fmt.Printf("Service : Fibonacci\n")
	fmt.Printf("Request : %v\n", req.N)

	for {
		res, err := stream.Recv()
		// if server stream is finished => err will be EOF (End Of File)
		if err == io.EOF {
			break
		}
		// other error
		if err != nil {
			return err
		}
		fmt.Printf("Response: %v\n", res.Result)
	}
	return nil
}

// Hello implements CalculatorService.
func (c calculatorService) Hello(name string) error {
	req := HelloRequest{
		Name:        name,
		CreatedDate: timestamppb.Now(),
	}

	res, err := c.calculatorClient.Hello(context.Background(), &req)
	if err != nil {
		return err
	}

	fmt.Printf("Service : Hello\n")
	fmt.Printf("Request : %v\n", req.Name)
	fmt.Printf("Response: %v\n", res.Result)
	return nil
}

func NewCalculatorService(calculatorClient CalculatorClient) CalculatorService {
	return calculatorService{calculatorClient}
}
