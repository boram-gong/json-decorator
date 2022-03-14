package svc

import (
	"github.com/boram-gong/json-decorator/service/svc/endpoint"
)

type Endpoints struct {
	JsonDecoratorEndpoint endpoint.Endpoint
	SaveRuleEndpoint      endpoint.Endpoint
	ReRuleEndpoint        endpoint.Endpoint
	DeleteRuleEndpoint    endpoint.Endpoint
	ReadRuleEndpoint      endpoint.Endpoint
}
