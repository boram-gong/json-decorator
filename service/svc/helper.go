package svc

import (
	"context"
	"net/http"
	"strings"

	json "github.com/json-iterator/go"
)

const contentType = "application/json; charset=utf-8"

// Helper functions

func headersToContext(ctx context.Context, r *http.Request) context.Context {
	for k, _ := range r.Header {
		// The key is added both in http format (k) which has had
		// http.CanonicalHeaderKey called on it in transport as well as the
		// strings.ToLower which is the grpc metadata format of the key so
		// that it can be accessed in either format
		ctx = context.WithValue(ctx, k, r.Header.Get(k))
		ctx = context.WithValue(ctx, strings.ToLower(k), r.Header.Get(k))
	}

	// add the access key to context
	accessKey := r.URL.Query().Get("access_key")
	if len(accessKey) > 0 {
		ctx = context.WithValue(ctx, "access_key", accessKey)
	}

	// Tune specific change.
	// also add the request url
	ctx = context.WithValue(ctx, "request-url", r.URL.Path)
	ctx = context.WithValue(ctx, "transport", "HTTPJSON")

	return ctx
}

func EncodeHTTPGenericResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	encoder := json.ConfigFastest.NewEncoder(w)
	encoder.SetEscapeHTML(false)
	return encoder.Encode(response)
}
