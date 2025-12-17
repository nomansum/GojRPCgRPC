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

	var p struct {
		ID int32 `json:"id"`
	}

	if err := json.Unmarshal(req.Params, &p); err != nil {
		p.ID = -1
	}

	resp := Response{
		JSONRPC: "2.0",
		ID:      int(p.ID),
		Result:  result,
	}

	if err != nil {
		resp.Error = err.Error()
		resp.Result = nil
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
