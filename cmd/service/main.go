package main

import (
	"log"
	"net"

	grpcsvc "jrpc/internal/grpc"
	pb "jrpc/internal/grpc/pb"

	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal(err)
	}

	server := grpc.NewServer()
	orderSvc := grpcsvc.NewOrderServer()

	pb.RegisterOrderServiceServer(server, orderSvc)
	pb.RegisterCancelServiceServer(server, orderSvc)
	log.Println("gRPC service running on :50051")
	server.Serve(lis)
}
