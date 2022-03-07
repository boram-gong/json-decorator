package operation

import (
	"errors"
	"fmt"
	"github.com/boram-gong/json-decorator/common"
	"github.com/boram-gong/json-decorator/rule"
)

func DecoratorJson(rules []*rule.Rule, jsonMap interface{}) error {
	var (
		split      = false
		rightValue interface{}
	)
	for _, r := range rules {
		r.MakeRule()
		if r.ERR != "" {
			return errors.New(fmt.Sprintf("%v %v %v %v", r.Key, r.Operation, r.Content, r.ERR))
		}
		if r.AT {
			atl := r.ATList[""]
			if len(atl) != 0 {
				realData := GetJsonValue(atl, jsonMap, r.Del)
				if r.Split && common.Interface2Slice(realData) != nil {
					split = true
				}
				rightValue = realData
			} else {
				realData := r.Content
				for k, at := range r.ATList {
					realData.(map[string]interface{})[k] = GetJsonValue(at, jsonMap, r.Del)
				}
				rightValue = realData
			}
			if err := SaveJsonMap(r.KeyList, jsonMap, r.RealOperation, split, rightValue); err != nil {
				return err
			}
		} else {
			if r.RealOperation == "rename" {
				rightValue = GetJsonValue(r.KeyList, jsonMap, r.Del)
				atl := r.ATList[""]
				if err := SaveJsonMap(atl, jsonMap, r.RealOperation, split, rightValue); err != nil {
					return err
				}
			} else if r.RealOperation == "delete" {
				GetJsonValue(r.KeyList, jsonMap, true)
			} else {
				rightValue = r.Content
				if err := SaveJsonMap(r.KeyList, jsonMap, r.RealOperation, split, rightValue); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func DealJson(name string, jsonMap interface{}) error {
	var (
		rules      = rule.AllRule.Load(name)
		split      = false
		rightValue interface{}
	)
	for _, r := range rules {
		if r.ERR != "" {
			return errors.New(fmt.Sprintf("%v %v %v %v", r.Key, r.Operation, r.Content, r.ERR))
		}
		if r.AT {
			// 如果存在 @
			atl := r.ATList[""] // r.ATList[""]存放着只含1个@的key
			if len(atl) != 0 {
				// 只含1个@
				realData := GetJsonValue(atl, jsonMap, r.Del)
				if r.Split && common.Interface2Slice(realData) != nil {
					split = true
				}
				rightValue = realData
			} else {
				// 键值对类型含@
				realData := r.Content
				for k, at := range r.ATList {
					realData.(map[string]interface{})[k] = GetJsonValue(at, jsonMap, r.Del)
				}
				rightValue = realData
			}
			if err := SaveJsonMap(r.KeyList, jsonMap, r.RealOperation, split, rightValue); err != nil {
				return err
			}
		} else {
			if r.RealOperation == "rename" {
				rightValue = GetJsonValue(r.KeyList, jsonMap, r.Del)
				atl := r.ATList[""]
				if err := SaveJsonMap(atl, jsonMap, r.RealOperation, split, rightValue); err != nil {
					return err
				}
			} else if r.RealOperation == "delete" {
				GetJsonValue(r.KeyList, jsonMap, true)
			} else {
				rightValue = r.Content
				if err := SaveJsonMap(r.KeyList, jsonMap, r.RealOperation, split, rightValue); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
