package http

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
)

type DecodeRequestFunc func(context.Context, *http.Request) (response interface{}, err error)
type EncodeRequestFunc func(context.Context, *http.Request, interface{}) error
type EncodeResponseFunc func(context.Context, http.ResponseWriter, interface{}) error
type DecodeResponseFunc func(context.Context, *http.Response) (response interface{}, err error)

type DecodeRequestFuncFromGin func(c *gin.Context) (interface{}, error)

func WrapS(c *gin.Context, f DecodeRequestFuncFromGin) DecodeRequestFunc {
	return func(ctx context.Context, request *http.Request) (response interface{}, err error) {
		return f(c)
	}
}
