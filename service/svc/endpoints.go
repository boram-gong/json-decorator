package svc

import (
	"github.com/boram-gong/json-decorator/service/svc/endpoint"
)

type Endpoints struct {
	JsonDecoratorEndpoint endpoint.Endpoint
	ReCfgEndpoint         endpoint.Endpoint
}
