package svc

import (
	"context"
	"encoding/json"
	"net/http"

	svc_http "github.com/boram-gong/json-decorator/service/svc/http"

	"github.com/gin-gonic/gin"
)

type errorWrapper struct {
	Error string `json:"error"`
}

func errorEncoder(_ context.Context, err error, w http.ResponseWriter) {
	body, _ := json.Marshal(errorWrapper{Error: err.Error()})
	if marshal, ok := err.(json.Marshaler); ok {
		if jsonBody, marshalErr := marshal.MarshalJSON(); marshalErr == nil {
			body = jsonBody
		}
	}
	w.Header().Set("Content-Type", contentType)
	if head, ok := err.(svc_http.Headerer); ok {
		for k := range head.Headers() {
			w.Header().Set(k, head.Headers().Get(k))
		}
	}
	code := http.StatusInternalServerError
	if sc, ok := err.(svc_http.StatusCoder); ok {
		code = sc.StatusCode()
	}
	w.WriteHeader(code)
	_, _ = w.Write(body)
}

func MakeHTTPHandler(engine *gin.Engine, endpoints Endpoints) {
	serverOptions := []svc_http.ServerOption{
		svc_http.ServerBefore(headersToContext),
		svc_http.ServerErrorEncoder(errorEncoder),
		svc_http.ServerErrorHandler(svc_http.NewNopErrorHandler()),
		svc_http.ServerAfter(svc_http.SetContentType(contentType)),
	}

	engine.Handle("GET", "/lgi/responseAdapter/json", func(c *gin.Context) {
		svc_http.NewServer(
			endpoints.JsonDecoratorEndpoint,
			svc_http.WrapS(c, DecodeTagJsonReq),
			EncodeHTTPGenericResponse,
			serverOptions...,
		).ServeHTTP(c.Writer, c.Request)
	})
	engine.Handle("GET", "/lgi/responseAdapter/rule", func(c *gin.Context) {
		svc_http.NewServer(
			endpoints.ReadRuleEndpoint,
			svc_http.WrapS(c, DecodeRuleId),
			EncodeHTTPGenericResponse,
			serverOptions...,
		).ServeHTTP(c.Writer, c.Request)
	})
	engine.Handle("POST", "/lgi/responseAdapter/rule", func(c *gin.Context) {
		svc_http.NewServer(
			endpoints.SaveRuleEndpoint,
			svc_http.WrapS(c, DecodePostRule),
			EncodeHTTPGenericResponse,
			serverOptions...,
		).ServeHTTP(c.Writer, c.Request)
	})
	engine.Handle("PUT", "/lgi/responseAdapter/rule", func(c *gin.Context) {
		svc_http.NewServer(
			endpoints.SaveRuleEndpoint,
			svc_http.WrapS(c, DecodePutRule),
			EncodeHTTPGenericResponse,
			serverOptions...,
		).ServeHTTP(c.Writer, c.Request)
	})
	engine.Handle("DELETE", "/lgi/responseAdapter/rule", func(c *gin.Context) {
		svc_http.NewServer(
			endpoints.DeleteRuleEndpoint,
			svc_http.WrapS(c, DecodeRuleId),
			EncodeHTTPGenericResponse,
			serverOptions...,
		).ServeHTTP(c.Writer, c.Request)
	})
	engine.Handle("GET", "/lgi/responseAdapter/re", func(c *gin.Context) {
		svc_http.NewServer(
			endpoints.ReRuleEndpoint,
			svc_http.WrapS(c, DecodeNull),
			EncodeHTTPGenericResponse,
			serverOptions...,
		).ServeHTTP(c.Writer, c.Request)
	})

}
