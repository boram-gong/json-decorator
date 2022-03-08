package handlers

import (
	"context"
	"github.com/boram-gong/json-decorator/operation"
	"github.com/boram-gong/json-decorator/rule"
	"github.com/boram-gong/json-decorator/service/body"
)

func JsonDecorator(ctx context.Context, request interface{}) (interface{}, error) {
	respBody := body.RespPool.Get().(*body.CommonResp)
	respBody.Init()
	reqBody := request.(*body.JsonReq)
	if reqBody.JsonMap != nil {
		respJson := reqBody.JsonMap
		err := operation.DecoratorJsonByRule(reqBody.Name, respJson)
		if err != nil {
			respBody.FailResp(400, err.Error())
		} else {
			respBody.Data = respJson
		}
	} else if reqBody.JsonSlice != nil {
		var respList []interface{}
		for _, j := range reqBody.JsonSlice {
			if err := operation.DecoratorJsonByRule(reqBody.Name, j); err != nil {
				respBody.FailResp(400, err.Error())
				return respBody, nil
			}
			respList = append(respList, j)
		}
		respBody.Data = respList
	} else {
		respBody.FailResp(404, "no data")
	}
	body.RespPool.Put(respBody)
	return respBody, nil
}

func ReCfg(ctx context.Context, request interface{}) (interface{}, error) {
	rule.GetRuleCfg()
	respBody := body.RespPool.Get().(*body.CommonResp)
	respBody.Init()
	body.RespPool.Put(respBody)
	return respBody, nil
}
