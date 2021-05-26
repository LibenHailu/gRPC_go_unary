package main

import (
	"context"
	"fmt"
	"log"

	"github.com/LibenHailu/grpc-go/greet/greetpb"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Hello i am a client")
	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("could not found connect %v ", err)
	}

	defer cc.Close()

	c := greetpb.NewGreetServiceClient(cc)
	// fmt.Printf("created client : %f", c)
	doUnary(c)

}

func doUnary(c greetpb.GreetServiceClient) {

	req := &greetpb.GreetRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Liben",
			LastName:  "Hailu",
		},
	}
	res, err := c.Greet(context.Background(), req)

	if err != nil {
		log.Fatalf("error while calling Greet RPC: %v", err)
	}

	log.Printf("Response form Great: %v", res.Result)
}
