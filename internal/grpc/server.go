package grpc

import (
	"context"
	"log"
	"strconv"

	pb "jrpc/internal/grpc/pb"
	"jrpc/internal/observability"
)

type OrderServer struct {
	pb.UnimplementedOrderServiceServer
	pb.UnimplementedCancelServiceServer
	orders map[int32]bool // added 16 th Dec.
}

func NewOrderServer() *OrderServer {
	return &OrderServer{
		orders: make(map[int32]bool),
	}
}

func (s *OrderServer) CreateOrder(
	ctx context.Context,
	req *pb.CreateOrderRequest,
) (*pb.CreateOrderResponse, error) {

	if _, ok := s.orders[req.Id]; !ok {
		s.orders[req.Id] = true
		msg := "Order " + strconv.Itoa(int(req.Id)) + " Created Successfully"
		log.Println("CreateOrder success", "order_id= ", req.Id)
		observability.OrdersCreated.Inc()
		return &pb.CreateOrderResponse{
			StatusCode: int32(201),
			Msg:        msg,
		}, nil
	}

	return &pb.CreateOrderResponse{
		StatusCode: 404,
		Msg:        "Order Already Exists!!!",
	}, nil
}

func (s *OrderServer) CancelOrder(
	ctx context.Context,
	req *pb.CancelOrderRequest,
) (*pb.CancelOrderResponse, error) {

	if _, ok := s.orders[req.Id]; !ok {
		return &pb.CancelOrderResponse{
			StatusCode: 404,
			Msg:        "No such order was created",
		}, nil
	}
	delete(s.orders, req.Id)

	return &pb.CancelOrderResponse{
		StatusCode: 200,
		Msg:        "Order " + strconv.Itoa(int(req.Id)) + " is cacelled !. ",
	}, nil
}
