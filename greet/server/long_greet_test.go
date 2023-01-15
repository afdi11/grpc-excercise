package main

import (
	"context"
	"testing"

	pb "github.com/afdi11/grpc-excercise/greet/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func TestLongGreet(t *testing.T) {
	ctx := context.Background()
	creds := grpc.WithTransportCredentials(insecure.NewCredentials())
	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), creds)

	if err != nil {
		t.Fatalf("Failed to dial bufnet: %v", err)
	}

	defer conn.Close()
	c := pb.NewGreetServiceClient(conn)

	requests := []*pb.GreetRequest{
		{
			FirstName: "Afdi",
		},
		{
			FirstName: "Marie",
		},
		{
			FirstName: "Test",
		},
	}

	stream, err := c.LongGreet(context.Background())

	if err != nil {
		t.Errorf("GreetManyTimes(%v) got unexpected error", err)
	}

	for _, req := range requests {
		stream.Send(req)
	}

	res, err := stream.CloseAndRecv()

	if err != nil {
		t.Errorf("Error while receiving response from LongGreet: %v", err)
	}

	expected := "Hello Afdi!\nHello Marie!\nHello Test!\n"

	if res.Result != expected {
		t.Errorf("Expected \"%s\" elements, got: \"%v\"", expected, res.Result)
	}
}
