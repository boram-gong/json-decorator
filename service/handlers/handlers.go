package handlers

import (
	"context"
	body "github.com/boram-gong/json-decorator/common/body"
	"github.com/boram-gong/json-decorator/operation"
	"github.com/boram-gong/json-decorator/rule"
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

func ReadRule(ctx context.Context, request interface{}) (interface{}, error) {
	if request.(int) == 0 {
		respBody := body.RespPool.Get().(*body.CommonResp)
		respBody.Init()
		respBody.Data = rule.GetAllRule()
		body.RespPool.Put(respBody)
		return respBody, nil
	} else {
		oneRule := rule.GetOneRule(request.(int))
		respBody := body.RespPool.Get().(*body.CommonResp)
		respBody.Init()
		respBody.Data = oneRule
		body.RespPool.Put(respBody)
		return respBody, nil
	}
}

func SaveRule(ctx context.Context, request interface{}) (interface{}, error) {
	respBody := body.RespPool.Get().(*body.CommonResp)
	respBody.Init()
	saveData := request.(*body.SaveRuleReq)
	if err := rule.SaveRule(saveData); err != nil {
		respBody.FailResp(500, err.Error())
	}
	respBody.Data = rule.ReAllRule()
	body.RespPool.Put(respBody)
	return respBody, nil
}

func DeleteRule(ctx context.Context, request interface{}) (interface{}, error) {
	respBody := body.RespPool.Get().(*body.CommonResp)
	respBody.Init()
	if err := rule.DeleteRule(request.(int)); err != nil {
		respBody.FailResp(400, err.Error())
	}
	respBody.Data = rule.ReAllRule()
	body.RespPool.Put(respBody)
	return respBody, nil
}

func ReRule(ctx context.Context, request interface{}) (interface{}, error) {
	rule.ReAllRule()
	respBody := body.RespPool.Get().(*body.CommonResp)
	respBody.Init()
	body.RespPool.Put(respBody)
	return respBody, nil
}
