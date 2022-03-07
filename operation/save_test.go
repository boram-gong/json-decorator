package operation

import (
	"fmt"
	"github.com/boram-gong/json-decorator/rule"
	json "github.com/json-iterator/go"
	"testing"
)

const (
	JSON = `
{
    "head": "head",
    "demo": {
        "key": "value"
    },
    "list1": [
        {
            "d1": "d1"
        },
        {
            "d2": "d2"
        }
    ],
    "list2": [
        [
            {
                "k1": "v1",
                "s": [
                    {
                        "xxx": 1,
                        "qqq": 1
                    }
                ]
            },
            {
                "k11": "v11"
            }
        ],
        [
            [
				{"k2": "v2"}
			]
        ]
    ]
}
`
)

func TestName(t *testing.T) {
	var data map[string]interface{}
	json.UnmarshalFromString(JSON, &data)
	old, _ := json.Marshal(data)
	fmt.Println(string(old))
	var kr = []rule.KeyStruct{
		{
			Key:   "list2",
			Index: []int{1, 0},
		},
	}
	kv := make(map[string]interface{})
	kv["new"] = "new"
	SaveJsonMap(kr, data, "append", true, []interface{}{1, 2, 3})
	newJson, _ := json.Marshal(data)
	fmt.Println(string(newJson))

}

func save1(source []interface{}, value interface{}, allIndex []int) {
	var p interface{}
	p = &source
	for _, n := range allIndex {
		switch p.(type) {
		case *[]interface{}:
			p = (*p.(*[]interface{}))[n]
		case []interface{}:
			p = &(p.([]interface{})[n])
		case *interface{}:
			p = &((*p.(*interface{})).([]interface{})[n])
		}

	}
	(*p.(*interface{})) = append((*p.(*interface{})).([]interface{}), value)
	fmt.Printf("%T \n", *p.(*interface{}))

	// source[1].([]interface{})[0] = append(source[1].([]interface{})[0].([]interface{}), value)

	fmt.Println(source)
}
