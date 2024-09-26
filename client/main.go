package main

import (
	"client/services"
	"flag"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
)

func main() {

	var cc *grpc.ClientConn
	var err error
	var creds credentials.TransportCredentials

	host := flag.String("host", "localhost:50051", "grpc server host")
	tls := flag.Bool("tls", false, "use a secure TLS connection")
	flag.Parse()

	if *tls {
		certFile := "../tls/ca.crt"
		creds, err = credentials.NewClientTLSFromFile(certFile, "")
		if err != nil {
			log.Fatal(err)
		}
	} else {
		creds = insecure.NewCredentials()
	}
	cc, err = grpc.NewClient(*host, grpc.WithTransportCredentials(creds))
	if err != nil {
		log.Fatal(err)
	}

	defer cc.Close()

	calculatorClient := services.NewCalculatorClient(cc)
	calculatorService := services.NewCalculatorService(calculatorClient)

	err = calculatorService.Hello("James")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// err = calculatorService.Fibonacci(5)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// err = calculatorService.Average(1,2,3,4,5)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// err = calculatorService.Sum(1,2,3,4,5)
	if err != nil {
		// if err come from grpc
		if grpcErr, ok := status.FromError(err); ok {
			log.Printf("[%v] %v\n", grpcErr.Code(), grpcErr.Message())
		} else {
			log.Fatal(err)
		}

	}
}
