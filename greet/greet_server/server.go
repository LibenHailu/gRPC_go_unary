package main

import (
	"fmt"
	"log"
	"net"

	"github.com/LibenHailu/grpc-go/greet/greetpb"
	"google.golang.org/grpc"
)

type server struct {
	greetpb.UnimplementedGreetServiceServer
}

func (*server)Greet(ctx context.Context, req *greetpb.GreetRequest) (*greetpb.GreetResponse, error){
	fmt.Printf("Greet funciton was invoked with %v \n",req)
	firstName:=req.GetGreeting().GetFirstName()
	result :="Hello" + firstName
	res := &greetpb.GreetResponse{
		Result: result,
	}
	return res, nil
}

func main() {

	fmt.Println("Calculator Server ")

	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("failed to listen: %v ", err)
	}

	s := grpc.NewServer()
	greetpb.RegisterGreetServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve:", err)
	}
}
