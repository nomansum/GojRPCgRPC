package main

import (
	"context"
	pb "jrpc/internal/grpc/pb"
	"log"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {

	ctx := context.Background()
	mux := runtime.NewServeMux()

	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}

	err := pb.RegisterOrderServiceHandlerFromEndpoint(
		ctx,
		mux,
		"grpc:50051",
		opts,
	)
	err = pb.RegisterCancelServiceHandlerFromEndpoint(
		ctx,
		mux,
		"grpc:50051",
		opts,
	)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("REST gRPC-Gateway running on : 8081")
	serverError := http.ListenAndServe(":8081", mux)
	if serverError != nil {
		log.Println("Server Error>> Can't start REST gRPC-Gateway")
	}

}
