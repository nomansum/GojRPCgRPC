package jrpc

import (
	"context"
	"encoding/json"
	"errors"
)

type HandlerFunc func(ctx context.Context, params json.RawMessage) (interface{}, error)

type Router struct {
	methods map[string]HandlerFunc
}

func (r *Router) Register(method string, handler HandlerFunc) {
	r.methods[method] = handler
}

func (r *Router) Call(
	ctx context.Context,
	method string,
	params json.RawMessage,
) (interface{}, error) {
	correspondingFunc, ok := r.methods[method]
	if !ok {
		return nil, errors.New("Method not found")
	}
	return correspondingFunc(ctx, params)

}
