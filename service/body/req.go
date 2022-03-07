package body

type JsonReq struct {
	Name      string                 `json:"name"`
	Data      interface{}            `json:"data"`
	JsonMap   map[string]interface{} `json:"-"`
	JsonSlice []interface{}          `json:"-"`
}
