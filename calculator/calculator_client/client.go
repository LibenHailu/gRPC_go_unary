package main

import (
	"context"
	"fmt"
	"log"

	"github.com/LibenHailu/grpc-go/calculator/calculatorpb"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Calcuator client  ")
	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("could not found connect %v ", err)
	}

	defer cc.Close()

	c := calculatorpb.NewCalculatorServiceClient(cc)
	// fmt.Printf("created client : %f", c)
	doUnary(c)

}

func doUnary(c calculatorpb.CalculatorServiceClient) {
	fmt.Println("starting to do a sum unary RPC...")
	req := &calculatorpb.SumRequest{
		FirstNumber:  5,
		SecondNumber: 40,
	}
	res, err := c.Sum(context.Background(), req)

	if err != nil {
		log.Fatalf("error while calling Sum RPC: %v", err)
	}

	log.Printf("Response form Sum: %v", res.SumResult)
}
