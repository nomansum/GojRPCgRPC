package gateway

import (
	"encoding/json"
	"errors"
	"net/http"

	"jrpc/internal/grpc"
)

func HandleJSONRPC(
	w http.ResponseWriter,
	r *http.Request,
	client *grpc.Client,
) {
	var req Request
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var (
		result interface{}
		err    error
	)

	switch req.Method {
	case "CreateOrder":
		result, err = TranscodeCreateOrder(r.Context(), req.Params, client)
	case "CancelOrder":
		result, err = TranscodeCancelOrder(r.Context(), req.Params, client)
	default:
		err = errors.New("method not found")
	}

	resp := Response{
		JSONRPC: "2.0",
		ID:      req.ID,
		Result:  result,
	}

	if err != nil {
		resp.Error = err.Error()
		resp.Result = nil
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
