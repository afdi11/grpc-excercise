package main

import pb "github.com/afdi11/grpc-excercise/greet/proto"

type Server struct {
	pb.GreetServiceServer
}
