package main

import pb "github.com/afdi11/grpc-excercise/calculator/proto"

type Server struct {
	pb.CalculatorServiceServer
}
