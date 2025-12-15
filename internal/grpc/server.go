package grpc

import (
	"context"

	pb "jrpc/internal/grpc/pb"
)

type OrderServer struct {
	pb.UnimplementedOrderServiceServer
	pb.UnimplementedCancelServiceServer
}

func (s *OrderServer) CreateOrder(
	ctx context.Context,
	req *pb.CreateOrderRequest,
) (*pb.CreateOrderResponse, error) {

	return &pb.CreateOrderResponse{
		Status: "order " + req.Id + " created",
	}, nil
}

func (s *OrderServer) CancelOrder(
	ctx context.Context,
	req *pb.CancelOrderRequest,
) (*pb.CancelOrderResponse, error) {

	return &pb.CancelOrderResponse{
		Status:     "order cancelled",
		StatusCode: 200,
	}, nil
}
