package grpc

import (
	"log"
	"sync"

	pb "jrpc/internal/grpc/pb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Client struct {
	conn          *grpc.ClientConn
	OrderService  pb.OrderServiceClient
	CancelService pb.CancelServiceClient
}

var (
	client *Client
	once   sync.Once
)

// NewClient creates (or returns) a singleton gRPC client
func NewClient(addr string) *Client {
	once.Do(func() {
		conn, err := grpc.Dial(
			addr,
			grpc.WithTransportCredentials(insecure.NewCredentials()),
		)
		if err != nil {
			log.Fatalf("failed to connect to gRPC server: %v", err)
		}

		client = &Client{
			conn:          conn,
			OrderService:  pb.NewOrderServiceClient(conn),
			CancelService: pb.NewCancelServiceClient(conn),
		}
	})

	return client
}

// Close closes the underlying gRPC connection
func (c *Client) Close() {
	if c.conn != nil {
		_ = c.conn.Close()
	}
}
