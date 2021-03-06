package test

import (
	"fmt"
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
		{"0.?????????",
			Struct2String(Json2Map(SOURCE1), ""),
			Json2Map(SOURCE1),
			&rule.Rule{Key: "key.k", Operation: "rename", Content: "key.newK", Stat: 1},
			Struct2String(result0, ""),
			"",
		},
		{"1.?????????",
			Struct2String(Json2Map(SOURCE1), ""),
			Json2Map(SOURCE1),
			&rule.Rule{Key: "key.k", Operation: "delete", Content: "", Stat: 1},
			Struct2String(result1, ""),
			"",
		},
		{"2.?????????",
			Struct2String(Json2Map(SOURCE1), ""),
			Json2Map(SOURCE1),
			&rule.Rule{Key: "k", Operation: "move", Content: "@key.k", Stat: 1},
			Struct2String(result2, ""),
			"Result0",
		},
		{"3.?????????-????????????",
			Struct2String(Json2Map(SOURCE1), ""),
			Json2Map(SOURCE1),
			&rule.Rule{Key: "key.k", Operation: "replace", Content: "newValue", Stat: 1},
			Struct2String(result3, ""),
			"",
		},
		{"4.?????????-@????????????",
			Struct2String(Json2Map(SOURCE2), ""),
			Json2Map(SOURCE2),
			&rule.Rule{Key: "head", Operation: "replace", Content: "@key.k", Stat: 1},
			Struct2String(result4, ""),
			"Result0",
		},
		{"5.??????????????????-?????????",
			Struct2String(Json2Map(SOURCE1), ""),
			Json2Map(SOURCE1),
			&rule.Rule{Key: "", Operation: "insert", Content: "{\"head\":\"insert\"}", Stat: 1},
			Struct2String(result5, ""),
			"Result0",
		},
		{"6.??????????????????-@????????????",
			Struct2String(Json2Map(SOURCE2), ""),
			Json2Map(SOURCE2),
			&rule.Rule{Key: "key", Operation: "insert", Content: "{\"newK\":\"@head\"}", Stat: 1},
			Struct2String(result6, ""),
			"Result0",
		},
		{"7.???????????????????????????-?????????",
			Struct2String(Json2Map(SOURCE3), ""),
			Json2Map(SOURCE3),
			&rule.Rule{Key: "list", Operation: "append-head", Content: "{\"l0\":\"l0\"}", Stat: 1},
			Struct2String(result7, ""),
			"",
		},
		{"8.???????????????????????????-?????????",
			Struct2String(Json2Map(SOURCE3), ""),
			Json2Map(SOURCE3),
			&rule.Rule{Key: "list", Operation: "append-tail", Content: "{\"l2\":\"l2\"}", Stat: 1},
			Struct2String(result8, ""),
			"",
		},
		{"9.???????????????????????????",
			Struct2String(Json2Map(SOURCE4), ""),
			Json2Map(SOURCE4),
			&rule.Rule{Key: "list", Operation: "move-head", Content: "{\"head\":\"@head\"}", Stat: 1},
			Struct2String(result9, ""),
			"",
		},
		{"10.???????????????????????????",
			Struct2String(Json2Map(SOURCE4), ""),
			Json2Map(SOURCE4),
			&rule.Rule{Key: "list", Operation: "move-tail", Content: "{\"head\":\"@head\"}", Stat: 1},
			Struct2String(result10, ""),
			"",
		},
		{"11.???????????????????????????-?????????",
			Struct2String(Json2Map(SOURCE3), ""),
			Json2Map(SOURCE3),
			&rule.Rule{Key: "list[0]", Operation: "replace", Content: "data", Stat: 1},
			Struct2String(result11, ""),
			"",
		},
		{"12.???????????????????????????-@????????????",
			Struct2String(Json2Map(SOURCE4), ""),
			Json2Map(SOURCE4),
			&rule.Rule{Key: "list[0]", Operation: "replace", Content: "{\"head\":\"@head\"}", Stat: 1},
			Struct2String(result12, ""),
			"Result0",
		},
		{"13.???????????????????????????????????????",
			Struct2String(Json2Map(SOURCE3), ""),
			Json2Map(SOURCE3),
			&rule.Rule{Key: "list", Operation: "append-tail", Content: "@list[-1]", Stat: 1},
			Struct2String(result13, ""),
			"",
		},
		{"15.????????????????????????????????????(?????????)",
			Struct2String(Json2Map(SOURCE5), ""),
			Json2Map(SOURCE5),
			&rule.Rule{Key: "list2", Operation: "move-tail", Content: "@list[...]", Stat: 1},
			Struct2String(result14, ""),
			"",
		},
		{"16.???????????????????????????",
			Struct2String(Json2Map(SOURCE6), ""),
			Json2Map(SOURCE6),
			&rule.Rule{Key: "key", Operation: "move", Content: "@list[0][1]", Stat: 1},
			Struct2String(result16, ""),
			"Result0",
		},
		{"17.?????????????????????key????????????",
			Struct2String(Json2Map(SOURCE6), ""),
			Json2Map(SOURCE6),
			&rule.Rule{Key: "key.l1", Operation: "insert", Content: "@list[0][1].l1", Stat: 1},
			Struct2String(result17, ""),
			"Result0",
		},
		{"18.???????????????????????????",
			Struct2String(Json2Map(SOURCE7), ""),
			Json2Map(SOURCE7),
			&rule.Rule{Key: "list2[0]", Operation: "move-head", Content: "@list[0][1]", Stat: 1},
			Struct2String(result18, ""),
			"Result0",
		},
		{"19.??????????????????????????????", // ?????????????????????null
			Struct2String(Json2Map(SOURCE7), ""),
			Json2Map(SOURCE7),
			&rule.Rule{Key: "list2[0]", Operation: "move-head", Content: "@list[0][2]", Stat: 1},
			Struct2String(result19, ""),
			"Result0",
		},
		{"20.??????????????????????????????", // ???????????????
			Struct2String(Json2Map(SOURCE3), ""),
			Json2Map(SOURCE3),
			&rule.Rule{Key: "list[8]", Operation: "replace", Content: "123", Stat: 1},
			Struct2String(Json2Map(SOURCE3), ""),
			"Result0",
		},
		{"21.???????????????????????????????????????", // ???????????????
			Struct2String(Json2Map(SOURCE3), ""),
			Json2Map(SOURCE3),
			&rule.Rule{Key: "list[0]", Operation: "insert", Content: "123", Stat: 1},
			Struct2String(Json2Map(SOURCE3), ""),
			"Result0",
		},
		{"22.?????????-json",
			Struct2String(Json2Map(SOURCE1), ""),
			Json2Map(SOURCE1),
			&rule.Rule{Key: "key", Operation: "replace", Content: "{\"newK\":\"new\"}", Stat: 1},
			Struct2String(result22, ""),
			"",
		},
		{"23.????????????",
			Struct2String(Json2Map(SOURCE1), ""),
			Json2Map(SOURCE1),
			&rule.Rule{Key: "key", Operation: "rename", Content: "new", Stat: 1},
			Struct2String(result23, ""),
			"",
		},
		{"24.???????????????????????????",
			Struct2String(Json2Map(SOURCE1), ""),
			Json2Map(SOURCE1),
			&rule.Rule{Key: "key.k", Operation: "append-head", Content: "head+", Stat: 1},
			Struct2String(result24, ""),
			"",
		},
		{"25.???????????????????????????",
			Struct2String(Json2Map(SOURCE1), ""),
			Json2Map(SOURCE1),
			&rule.Rule{Key: "key.k", Operation: "append-tail", Content: "+tail", Stat: 1},
			Struct2String(result25, ""),
			"",
		},
		{"26.????????????????????????????????????(?????????)",
			Struct2String(Json2Map(SOURCE5), ""),
			Json2Map(SOURCE5),
			&rule.Rule{Key: "list2", Operation: "move-head", Content: "@list[...]", Stat: 1},
			Struct2String(result26, ""),
			"",
		},
		{"27.???????????????????????????",
			Struct2String(Json2Map(SOURCE5), ""),
			Json2Map(SOURCE5),
			&rule.Rule{Key: "list2", Operation: "replace", Content: "@list", Stat: 1},
			Struct2String(result27, ""),
			"Result0",
		},
		{"28.????????????????????????",
			Struct2String(Json2Map(SOURCE8), ""),
			Json2Map(SOURCE8),
			&rule.Rule{Key: "head", Operation: "replace", Content: 2, Stat: 1},
			Struct2String(result28, ""),
			"",
		},
		{"29.???????????????????????????-1???",
			Struct2String(Json2Map(SOURCE5), ""),
			Json2Map(SOURCE5),
			&rule.Rule{Key: "list[0]", Operation: "delete", Content: "", Stat: 1},
			Struct2String(result29, ""),
			"Result0",
		},
		{"30.???????????????????????????-??????",
			Struct2String(Json2Map(SOURCE7), ""),
			Json2Map(SOURCE7),
			&rule.Rule{Key: "list[0][0]", Operation: "delete", Content: "", Stat: 1},
			Struct2String(result30, ""),
			"Result0",
		},
		{"31.?????????????????????????????????",
			Struct2String(Json2Map(SOURCE8), ""),
			Json2Map(SOURCE8),
			&rule.Rule{Key: "list[0]", Operation: "append-head", Content: "0", Stat: 1},
			Struct2String(result31, ""),
			"Result0",
		},
		{"32.????????????????????????",
			Struct2String(Json2Map(SOURCE3), ""),
			Json2Map(SOURCE3),
			&rule.Rule{Key: "list[0].l1", Operation: "rename", Content: "list[0].l2", Stat: 1},
			Struct2String(result32, ""),
			"",
		},
		{"32.????????????(??????)????????????",
			Struct2String(Json2Map(SOURCE9), ""),
			Json2Map(SOURCE9),
			&rule.Rule{Key: "list[0][0].l1", Operation: "rename", Content: "list[0][0].l2", Stat: 1},
			Struct2String(result33, ""),
			"",
		},
		// ?????????????????????????????????????????????????????????
	}

	for _, tt := range tests {
		if err := operation.DecoratorJson([]*rule.Rule{tt.TestRule}, tt.Json); err != nil {
			fmt.Printf("[%v] ????????????: %v \n", tt.TestName, err)
		} else {
			getResult := Struct2String(tt.Json, tt.ResultType)
			if strings.Compare(getResult, tt.RealResult) != 0 {
				fmt.Printf("[%v] %v??????%v?????????:%v fail!\n", tt.TestName, tt.Source, tt.RealResult, getResult)
				t.Fail()
				return
			} else {
				// fmt.Printf("[%v] %v??????%v?????????:%v success!\n", tt.TestName, tt.Source, tt.RealResult, getResult)
				fmt.Printf("[%v] success!\n", tt.TestName)
			}
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
