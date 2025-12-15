package jrpc

import (
	"encoding/json"
	"net/http"
)

func Handler(router *Router) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req Request

		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		resp := Response{
			JSONRPC: "2.0",
			ID:      req.ID,
		}

		// Basic JSON-RPC validation
		if req.JSONRPC != "2.0" || req.Method == "" {
			resp.Error = &Error{
				Code:    -32600,
				Message: "Invalid Request",
			}
		} else {
			result, err := router.Call(r.Context(), req.Method, req.Params)
			if err != nil {
				resp.Error = &Error{
					Code:    -32601,
					Message: err.Error(),
				}
			} else {
				resp.Result = result
			}
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)
	}
}
