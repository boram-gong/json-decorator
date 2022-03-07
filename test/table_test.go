package test

import (
	"fmt"
	"github.com/boram-gong/json-decorator/common"
	"github.com/boram-gong/json-decorator/operation"
	"github.com/boram-gong/json-decorator/rule"
	json "github.com/json-iterator/go"
	"strings"
	"testing"
)

const (
	SOURCE1 = `
{
	"key": {
		"k":"k"
	}
}
`
	SOURCE2 = `
{
	"head": "head",
	"key": {
		"k":"k"
	}
}
`
	SOURCE3 = `
{
	"list": [
		{"l1": "l1"}
	]
}
`
	SOURCE4 = `
{
	"head": "l",
	"list": [
		{"l1": "l1"}
	]
}
`
	SOURCE5 = `
{
	"list": [
		{"l0": "l0"},
		{"l1": "l1"}
	],
	"list2": [
		{"l2": "l2"}
	]
}
`
	SOURCE6 = `
{
	"list": [
		[
			{"l0": "l0"},
			{"l1": "l1"}
		]
	],
	"key": {
		"k":"k"
	}
}
`
	SOURCE7 = `
{
	"list": [
		[
			{"l0": "l0"},
			{"l1": "l1"}
		]
	],
	"list2": [
		[
			{"l2": "l2"}
		]
	]
}
`

	SOURCE8 = `
{
	"head": 1,
}
`
	SOURCE9 = `
{
	"list": [
		[
			{"l1": "l1"}
		]
	]
}
`
)

type Son0 struct {
	K    string `json:"k,omitempty"`
	NewK string `json:"newK,omitempty"`
	L1   string `json:"l1,omitempty"`
}

type Son1 struct {
	Head string `json:"head,omitempty"`
	L0   string `json:"l0,omitempty"`
	L1   string `json:"l1,omitempty"`
	L2   string `json:"l2,omitempty"`
}

type Result0 struct {
	Head  interface{}   `json:"head,omitempty"`
	Key   *Son0         `json:"key,omitempty"`
	New   *Son0         `json:"new,omitempty"`
	K     string        `json:"k,omitempty"`
	List  []interface{} `json:"list,omitempty"`
	List2 []interface{} `json:"list2,omitempty"`
}

var (
	result0 = Result0{
		Key: &Son0{
			NewK: "k",
		},
	}
	result1 = Result0{
		Key: &Son0{},
	}
	result2 = Result0{
		Key: &Son0{},
		K:   "k",
	}
	result3 = Result0{
		Key: &Son0{
			K: "newValue",
		},
	}
	result4 = Result0{
		Head: "k",
		Key: &Son0{
			K: "k",
		},
	}
	result5 = Result0{
		Head: "insert",
		Key: &Son0{
			K: "k",
		},
	}
	result6 = Result0{
		Head: "head",
		Key: &Son0{
			K:    "k",
			NewK: "head",
		},
	}
	result7 = Result0{
		List: []interface{}{
			Son1{L0: "l0"},
			Son1{L1: "l1"},
		},
	}
	result8 = Result0{
		List: []interface{}{
			Son1{L1: "l1"},
			Son1{L2: "l2"},
		},
	}
	result9 = Result0{
		List: []interface{}{
			Son1{Head: "l"},
			Son1{L1: "l1"},
		},
	}
	result10 = Result0{
		List: []interface{}{
			Son1{L1: "l1"},
			Son1{Head: "l"},
		},
	}
	result11 = Result0{
		List: []interface{}{
			"data",
		},
	}
	result12 = Result0{
		Head: "l",
		List: []interface{}{
			Son1{Head: "l"},
		},
	}
	result13 = Result0{
		List: []interface{}{
			Son1{L1: "l1"},
			Son1{L1: "l1"},
		},
	}
	result14 = Result0{
		List2: []interface{}{
			Son1{L2: "l2"},
			Son1{L0: "l0"},
			Son1{L1: "l1"},
		},
	}
	result16 = Result0{
		Key: &Son0{
			K:  "k",
			L1: "l1",
		},
		List: []interface{}{
			[]interface{}{Son1{L0: "l0"}},
		},
	}
	result17 = Result0{
		Key: &Son0{
			K:  "k",
			L1: "l1",
		},
		List: []interface{}{
			[]interface{}{Son1{L0: "l0"}, Son1{L1: "l1"}},
		},
	}
	result18 = Result0{
		List: []interface{}{
			[]interface{}{Son1{L0: "l0"}},
		},
		List2: []interface{}{
			[]interface{}{Son1{L1: "l1"}, Son1{L2: "l2"}},
		},
	}
	result19 = Result0{
		List: []interface{}{
			[]interface{}{Son1{L0: "l0"}, Son1{L1: "l1"}},
		},
		List2: []interface{}{
			[]interface{}{nil, Son1{L2: "l2"}},
		},
	}
	result22 = Result0{
		Key: &Son0{NewK: "new"},
	}
	result23 = Result0{
		New: &Son0{K: "k"},
	}
	result24 = Result0{
		Key: &Son0{K: "head+k"},
	}
	result25 = Result0{
		Key: &Son0{K: "k+tail"},
	}
	result26 = Result0{
		List2: []interface{}{
			Son1{L0: "l0"},
			Son1{L1: "l1"},
			Son1{L2: "l2"},
		},
	}
	result27 = Result0{
		List: []interface{}{
			Son1{L0: "l0"},
			Son1{L1: "l1"},
		},
		List2: []interface{}{
			Son1{L0: "l0"},
			Son1{L1: "l1"},
		},
	}
	result28 = Result0{
		Head: 2,
	}
	result29 = Result0{
		List: []interface{}{
			Son1{L1: "l1"},
		},
		List2: []interface{}{
			Son1{L2: "l2"},
		},
	}
	result30 = Result0{
		List: []interface{}{
			[]interface{}{Son1{L1: "l1"}},
		},
		List2: []interface{}{
			[]interface{}{Son1{L2: "l2"}},
		},
	}
	result31 = Result0{
		Head: 1,
		List: []interface{}{
			"0",
		},
	}
	result32 = Result0{
		List: []interface{}{
			Son1{L2: "l1"},
		},
	}
	result33 = Result0{
		List: []interface{}{
			[]interface{}{Son1{L2: "l1"}},
		},
	}
)

func TestTable(t *testing.T) {
	tests := []struct {
		TestName   string
		Source     string
		Json       interface{}
		TestRule   *rule.Rule
		RealResult string
		ResultType string
	}{
		{"0.键更名",
			Struct2String(Json2Map(SOURCE1), ""),
			Json2Map(SOURCE1),
			&rule.Rule{Key: "key.k", Operation: "rename", Content: "key.newK"},
			Struct2String(result0, ""),
			"",
		},
		{"1.键删除",
			Struct2String(Json2Map(SOURCE1), ""),
			Json2Map(SOURCE1),
			&rule.Rule{Key: "key.k", Operation: "delete", Content: ""},
			Struct2String(result1, ""),
			"",
		},
		{"2.键移动",
			Struct2String(Json2Map(SOURCE1), ""),
			Json2Map(SOURCE1),
			&rule.Rule{Key: "k", Operation: "move", Content: "@key.k"},
			Struct2String(result2, ""),
			"Result0",
		},
		{"3.值替换-自定义值",
			Struct2String(Json2Map(SOURCE1), ""),
			Json2Map(SOURCE1),
			&rule.Rule{Key: "key.k", Operation: "replace", Content: "newValue"},
			Struct2String(result3, ""),
			"",
		},
		{"4.值替换-@其他键值",
			Struct2String(Json2Map(SOURCE2), ""),
			Json2Map(SOURCE2),
			&rule.Rule{Key: "head", Operation: "replace", Content: "@key.k"},
			Struct2String(result4, ""),
			"Result0",
		},
		{"5.新插入键值对-自定义",
			Struct2String(Json2Map(SOURCE1), ""),
			Json2Map(SOURCE1),
			&rule.Rule{Key: "", Operation: "insert", Content: "{\"head\":\"insert\"}"},
			Struct2String(result5, ""),
			"Result0",
		},
		{"6.新插入键值对-@其他键值",
			Struct2String(Json2Map(SOURCE2), ""),
			Json2Map(SOURCE2),
			&rule.Rule{Key: "key", Operation: "insert", Content: "{\"newK\":\"@head\"}"},
			Struct2String(result6, ""),
			"Result0",
		},
		{"7.在数组头部新增数据-自定义",
			Struct2String(Json2Map(SOURCE3), ""),
			Json2Map(SOURCE3),
			&rule.Rule{Key: "list", Operation: "append-head", Content: "{\"l0\":\"l0\"}"},
			Struct2String(result7, ""),
			"",
		},
		{"8.在数组尾部新增数据-自定义",
			Struct2String(Json2Map(SOURCE3), ""),
			Json2Map(SOURCE3),
			&rule.Rule{Key: "list", Operation: "append-tail", Content: "{\"l2\":\"l2\"}"},
			Struct2String(result8, ""),
			"",
		},
		{"9.移动数据到数组头部",
			Struct2String(Json2Map(SOURCE4), ""),
			Json2Map(SOURCE4),
			&rule.Rule{Key: "list", Operation: "move-head", Content: "{\"head\":\"@head\"}"},
			Struct2String(result9, ""),
			"",
		},
		{"10.移动数据到数组尾部",
			Struct2String(Json2Map(SOURCE4), ""),
			Json2Map(SOURCE4),
			&rule.Rule{Key: "list", Operation: "move-tail", Content: "{\"head\":\"@head\"}"},
			Struct2String(result10, ""),
			"",
		},
		{"11.替换数组某索引的值-自定义",
			Struct2String(Json2Map(SOURCE3), ""),
			Json2Map(SOURCE3),
			&rule.Rule{Key: "list[0]", Operation: "replace", Content: "data"},
			Struct2String(result11, ""),
			"",
		},
		{"12.替换数组某索引的值-@其他键值",
			Struct2String(Json2Map(SOURCE4), ""),
			Json2Map(SOURCE4),
			&rule.Rule{Key: "list[0]", Operation: "replace", Content: "{\"head\":\"@head\"}"},
			Struct2String(result12, ""),
			"Result0",
		},
		{"13.取数组最后一个值的用法示例",
			Struct2String(Json2Map(SOURCE3), ""),
			Json2Map(SOURCE3),
			&rule.Rule{Key: "list", Operation: "append-tail", Content: "@list[-1]"},
			Struct2String(result13, ""),
			"",
		},
		{"15.取数组每一个值的用法示例(尾增加)",
			Struct2String(Json2Map(SOURCE5), ""),
			Json2Map(SOURCE5),
			&rule.Rule{Key: "list2", Operation: "move-tail", Content: "@list[...]"},
			Struct2String(result14, ""),
			"",
		},
		{"16.深层取数组用法示例",
			Struct2String(Json2Map(SOURCE6), ""),
			Json2Map(SOURCE6),
			&rule.Rule{Key: "key", Operation: "move", Content: "@list[0][1]"},
			Struct2String(result16, ""),
			"Result0",
		},
		{"17.深层取数组值的key用法示例",
			Struct2String(Json2Map(SOURCE6), ""),
			Json2Map(SOURCE6),
			&rule.Rule{Key: "key.l1", Operation: "insert", Content: "@list[0][1].l1"},
			Struct2String(result17, ""),
			"Result0",
		},
		{"18.深层取数组值的迁移",
			Struct2String(Json2Map(SOURCE7), ""),
			Json2Map(SOURCE7),
			&rule.Rule{Key: "list2[0]", Operation: "move-head", Content: "@list[0][1]"},
			Struct2String(result18, ""),
			"Result0",
		},
		{"19.错误的规则：越界取值", // 越界取值会取出null
			Struct2String(Json2Map(SOURCE7), ""),
			Json2Map(SOURCE7),
			&rule.Rule{Key: "list2[0]", Operation: "move-head", Content: "@list[0][2]"},
			Struct2String(result19, ""),
			"Result0",
		},
		{"20.错误的规则：越界存值", // 不发生变化
			Struct2String(Json2Map(SOURCE3), ""),
			Json2Map(SOURCE3),
			&rule.Rule{Key: "list[8]", Operation: "replace", Content: "123"},
			Struct2String(Json2Map(SOURCE3), ""),
			"Result0",
		},
		{"21.错误的规则：操作表达式错误", // 不发生变化
			Struct2String(Json2Map(SOURCE3), ""),
			Json2Map(SOURCE3),
			&rule.Rule{Key: "list[0]", Operation: "insert", Content: "123"},
			Struct2String(Json2Map(SOURCE3), ""),
			"Result0",
		},
		{"22.值替换-json",
			Struct2String(Json2Map(SOURCE1), ""),
			Json2Map(SOURCE1),
			&rule.Rule{Key: "key", Operation: "replace", Content: "{\"newK\":\"new\"}"},
			Struct2String(result22, ""),
			"",
		},
		{"23.顶级更名",
			Struct2String(Json2Map(SOURCE1), ""),
			Json2Map(SOURCE1),
			&rule.Rule{Key: "key", Operation: "rename", Content: "new"},
			Struct2String(result23, ""),
			"",
		},
		{"24.字符串操作：头增加",
			Struct2String(Json2Map(SOURCE1), ""),
			Json2Map(SOURCE1),
			&rule.Rule{Key: "key.k", Operation: "append-head", Content: "head+"},
			Struct2String(result24, ""),
			"",
		},
		{"25.字符串操作：尾增加",
			Struct2String(Json2Map(SOURCE1), ""),
			Json2Map(SOURCE1),
			&rule.Rule{Key: "key.k", Operation: "append-tail", Content: "+tail"},
			Struct2String(result25, ""),
			"",
		},
		{"26.取数组每一个值的用法示例(头增加)",
			Struct2String(Json2Map(SOURCE5), ""),
			Json2Map(SOURCE5),
			&rule.Rule{Key: "list2", Operation: "move-head", Content: "@list[...]"},
			Struct2String(result26, ""),
			"",
		},
		{"27.值类型为数组的替换",
			Struct2String(Json2Map(SOURCE5), ""),
			Json2Map(SOURCE5),
			&rule.Rule{Key: "list2", Operation: "replace", Content: "@list"},
			Struct2String(result27, ""),
			"Result0",
		},
		{"28.其他类型直接替换",
			Struct2String(Json2Map(SOURCE8), ""),
			Json2Map(SOURCE8),
			&rule.Rule{Key: "head", Operation: "replace", Content: 2},
			Struct2String(result28, ""),
			"",
		},
		{"29.删除数组某一个元素-1层",
			Struct2String(Json2Map(SOURCE5), ""),
			Json2Map(SOURCE5),
			&rule.Rule{Key: "list[0]", Operation: "delete", Content: ""},
			Struct2String(result29, ""),
			"Result0",
		},
		{"30.删除数组某一个元素-多层",
			Struct2String(Json2Map(SOURCE7), ""),
			Json2Map(SOURCE7),
			&rule.Rule{Key: "list[0][0]", Operation: "delete", Content: ""},
			Struct2String(result30, ""),
			"Result0",
		},
		{"31.往一个不存在的数组增值",
			Struct2String(Json2Map(SOURCE8), ""),
			Json2Map(SOURCE8),
			&rule.Rule{Key: "list[0]", Operation: "append-head", Content: "0"},
			Struct2String(result31, ""),
			"Result0",
		},
		{"32.修改数组内部的值",
			Struct2String(Json2Map(SOURCE3), ""),
			Json2Map(SOURCE3),
			&rule.Rule{Key: "list[0].l1", Operation: "rename", Content: "list[0].l2"},
			Struct2String(result32, ""),
			"",
		},
		{"32.修改数组(多层)内部的值",
			Struct2String(Json2Map(SOURCE9), ""),
			Json2Map(SOURCE9),
			&rule.Rule{Key: "list[0][0].l1", Operation: "rename", Content: "list[0][0].l2"},
			Struct2String(result33, ""),
			"",
		},
		// 删除数组某一个元素、新增数组、深度存值
	}

	for _, tt := range tests {
		tt.TestRule.MakeRule()
		work(tt.TestRule, tt.Json)

		getResult := Struct2String(tt.Json, tt.ResultType)
		if strings.Compare(getResult, tt.RealResult) != 0 {
			fmt.Printf("[%v] %v变为%v，结果:%v fail!\n", tt.TestName, tt.Source, tt.RealResult, getResult)
			t.Fail()
			return
		} else {
			// fmt.Printf("[%v] %v变为%v，结果:%v success!\n", tt.TestName, tt.Source, tt.RealResult, getResult)
			fmt.Printf("[%v] success!\n", tt.TestName)
		}
	}
}

func Json2Map(s string) interface{} {
	var m map[string]interface{}
	json.UnmarshalFromString(s, &m)
	return m
}

func Struct2String(data interface{}, Type string) string {
	body, _ := json.Marshal(data)
	switch Type {
	case "Result0":
		var temp Result0
		json.Unmarshal(body, &temp)
		body, _ = json.Marshal(temp)

	}
	return string(body)
}

func work(r *rule.Rule, jsonMap interface{}) {
	var (
		split      = false
		rightValue interface{}
	)
	if r.AT {
		atl := r.ATList[""]
		if len(atl) != 0 {
			realData := operation.GetJsonValue(atl, jsonMap, r.Del)
			if r.Split && common.Interface2Slice(realData) != nil {
				split = true
			}
			rightValue = realData
		} else {
			realData := r.Content
			for k, at := range r.ATList {
				realData.(map[string]interface{})[k] = operation.GetJsonValue(at, jsonMap, r.Del)
			}
			rightValue = realData
		}
		operation.SaveJsonMap(r.KeyList, jsonMap, r.RealOperation, split, rightValue)
	} else {
		if r.RealOperation == "rename" {
			rightValue = operation.GetJsonValue(r.KeyList, jsonMap, r.Del)
			if r.Split && common.Interface2Slice(rightValue) != nil {
				for _, value := range common.Interface2Slice(rightValue) {
					operation.SaveJsonMap(nil, value, r.RealOperation, split, rightValue)
				}
			} else {
				atl := r.ATList[""]
				operation.SaveJsonMap(atl, jsonMap, r.RealOperation, split, rightValue)
			}
		} else if r.RealOperation == "delete" {
			operation.GetJsonValue(r.KeyList, jsonMap, true)
		} else {
			rightValue = r.Content
			operation.SaveJsonMap(r.KeyList, jsonMap, r.RealOperation, split, rightValue)
		}
	}
}
