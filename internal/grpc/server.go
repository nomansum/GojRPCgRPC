package grpc

import (
	"context"
	"fmt"
	"strconv"

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
		Status:     "order " + req.Id + " created",
		StatusCode: 201,
	}, nil
}

func (s *OrderServer) CancelOrder(
	ctx context.Context,
	req *pb.CancelOrderRequest,
) (*pb.CancelOrderResponse, error) {

	fmt.Println("printing req status : " + req.Status)
	fmt.Println("Printing Req satus code : " + strconv.Itoa(int(req.StatusCode)))
	return &pb.CancelOrderResponse{
		Status:     "order" + req.Status + "cancelled",
		StatusCode: req.StatusCode,
	}, nil
}
