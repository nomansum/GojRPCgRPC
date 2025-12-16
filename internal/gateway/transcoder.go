package gateway

import (
	"context"
	"encoding/json"

	"jrpc/internal/grpc"
	pb "jrpc/internal/grpc/pb"
)

func TranscodeCreateOrder(
	ctx context.Context,
	params json.RawMessage,
	client *grpc.Client,
) (interface{}, error) {

	var p struct {
		ID int32 `json:"id"`
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

	return map[string]interface{}{
		"statusCode": resp.StatusCode,
		"msg":        resp.Msg,
	}, nil
}

func TranscodeCancelOrder(ctx context.Context,
	params json.RawMessage,
	client *grpc.Client,
) (interface{}, error) {

	var p struct {
		ID int32 `json:"id"`
	}

	if err := json.Unmarshal(params, &p); err != nil {
		return nil, err
	}

	resp, err := client.CancelService.CancelOrder(ctx, &pb.CancelOrderRequest{
		Id: p.ID,
	})
	if err != nil {
		return nil, err
	}

	return map[string]interface{}{
		"statusCode": resp.StatusCode,
		"msg":        resp.Msg,
	}, nil

}
