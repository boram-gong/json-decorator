package http

import "context"

type ErrorHandler interface {
	Handle(ctx context.Context, err error)
}

type NopErrorHandler struct {
}

func NewNopErrorHandler() *NopErrorHandler {
	return &NopErrorHandler{}
}

func (n *NopErrorHandler) Handle(ctx context.Context, err error) {
}
