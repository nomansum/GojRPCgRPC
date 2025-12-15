package gateway

import (
	"context"
	"encoding/json"
	"strconv"

	"jrpc/internal/grpc"
	pb "jrpc/internal/grpc/pb"
)

func TranscodeCreateOrder(
	ctx context.Context,
	params json.RawMessage,
	client *grpc.Client,
) (interface{}, error) {

	var p struct {
		ID string `json:"id"`
	}

	if err := json.Unmarshal(params, &p); err != nil {
		return nil, err
	}

	resp, err := client.OrderService.CreateOrder(ctx, &pb.CreateOrderRequest{
		Id: p.ID,
	})
	if err != nil {
		return nil, err
	}

	return map[string]string{
		"status":     resp.Status,
		"statusCode": strconv.Itoa(int(resp.StatusCode)),
	}, nil
}

func TranscodeCancelOrder(ctx context.Context,
	params json.RawMessage,
	client *grpc.Client,
) (interface{}, error) {

	var p struct {
		status     string `json:"status"`
		statusCode int    `json:"statusCode"`
	}

	if err := json.Unmarshal(params, &p); err != nil {
		return nil, err
	}

	resp, err := client.CancelService.CancelOrder(ctx, &pb.CancelOrderRequest{
		Status:     p.status,
		StatusCode: int32(p.statusCode),
	})
	if err != nil {
		return nil, err
	}

	return map[string]string{
		"status":     resp.Status,
		"statusCode": strconv.Itoa(int(resp.StatusCode)),
	}, nil

}
