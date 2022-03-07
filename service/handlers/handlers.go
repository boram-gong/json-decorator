package handlers

import (
	"context"
	"github.com/boram-gong/json-decorator/operation"
	"github.com/boram-gong/json-decorator/rule"
	"github.com/boram-gong/json-decorator/service/body"
)

func JsonDecorator(ctx context.Context, request interface{}) (interface{}, error) {
	respBody := body.NewCommonResp()
	reqBody := request.(*body.JsonReq)
	if reqBody.JsonMap != nil {
		respJson := reqBody.JsonMap
		err := operation.DecoratorJsonByRule(reqBody.Name, respJson)
		if err != nil {
			respBody.FailResp(400, err.Error())
		} else {
			respBody.Data = respJson
		}
		return respBody, nil
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
		return respBody, nil
	} else {
		respBody.FailResp(404, "no data")
		return respBody, nil
	}
}

func ReCfg(ctx context.Context, request interface{}) (interface{}, error) {
	rule.GetRuleCfg()
	return body.NewCommonResp(), nil
}
